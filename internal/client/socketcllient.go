package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/constant"
)

// SocketClient handles socket client
type SocketClient struct {
	conn net.Conn
}

// New creates a new socket client
func NewSocketClient(hostname string, port string) (*SocketClient, error) {
	d := net.Dialer{
		Timeout: constant.CONNECTION_TIMEOUT,
	}
	client, err := d.Dial("tcp", hostname+":"+port)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s:%s: %w", hostname, port, err)
	}
	return &SocketClient{
		conn: client,
	}, nil
}

// Close closes the socket client
func (c *SocketClient) Close() {
	c.conn.Close()
}

// Send sends a message to the server
func (c *SocketClient) Send(msg string) {
	// set write deadline to now + constant.WRITE_TIMEOUT
	c.conn.SetWriteDeadline(time.Now().Add(constant.WRITE_TIMEOUT))
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		logrus.Errorf("failed to send message: %v\n", err)
	}
}

// Receive receives a message from the server
func (c *SocketClient) Receive() (int, []byte, error) {
	var received int
	buffer := bytes.NewBuffer(nil)
	for {
		chunk := make([]byte, constant.CHUNK_SIZE)
		read, err := c.conn.Read(chunk)
		if err != nil && !errors.Is(err, io.EOF) {
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
