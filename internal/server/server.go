package server

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/constant"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/grep"
)

// Server handles server
type Server struct {
	server net.Listener
}

// New creates a new server
func New(port string) (*Server, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to get hostname: %w", err)
	}
	server, err := net.Listen("tcp", hostName+":"+port)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s:%s: %w", hostName, port, err)
	}
	logrus.Printf("Listening on %s:%s\n", hostName, port)
	return &Server{
		server: server,
	}, nil
}

// Close closes the server
func (s *Server) Close() {
	s.server.Close()
}

// Accept accepts a new connection
func (s *Server) Accept() (net.Conn, error) {
	conn, err := s.server.Accept()
	if err != nil {
		return nil, fmt.Errorf("failed to accept new connection: %w", err)
	}
	return conn, nil
}

// Run runs the server
func (s *Server) Run(handler *grep.Handler) {
	for {
		conn, err := s.Accept()
		if err != nil {
			logrus.Printf("failed to accept new connection: %v\n", err)
			continue
		}
		go handleConnection(conn, handler)
	}
}

// handleConnection handles a new connection
func handleConnection(conn net.Conn, handler *grep.Handler) {
	defer conn.Close()
	_, content, err := receive(conn)
	if err != nil {
		logrus.Errorf("failed to receive message: %v\n", err)
		return
	}
	result, err := handler.Handle(strings.Split(string(content), " "))
	if err != nil {
		logrus.Error(err)
		conn.Write([]byte(err.Error()))
		return
	}
	// set write deadline to now + 10 seconds
	if _, err = send(conn, result); err != nil {
		logrus.Errorf("failed to send message: %v\n", err)
		return
	}
}

// send sends a message to the client connection
func send(conn net.Conn, msg string) (int, error) {
	var sent int
	for {
		if len(msg) == 0 {
			break
		}
		chunk := []byte(msg)
		if len(chunk) > constant.CHUNK_SIZE {
			chunk = chunk[:constant.CHUNK_SIZE]
		}
		n, err := conn.Write(chunk)
		if err != nil {
			return sent, fmt.Errorf("failed to write to connection: %w", err)
		}
		sent += n
		msg = msg[n:]
	}
	return sent, nil
}

// receive receives a message from the client connection
func receive(conn net.Conn) (int, []byte, error) {
	var received int
	buffer := bytes.NewBuffer(nil)
	for {
		chunk := make([]byte, constant.CHUNK_SIZE)
		read, err := conn.Read(chunk)
		if err != nil {
			return received, buffer.Bytes(), fmt.Errorf("failed to read from connection: %w", err)
		}
		received += read
		buffer.Write(chunk[:read])

		if read == 0 || read < constant.CHUNK_SIZE {
			break
		}
	}
	return received, buffer.Bytes(), nil
}
