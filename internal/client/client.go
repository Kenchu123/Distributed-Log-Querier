package client

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
)

type Client struct {
	machines []config.Machine
}

type Result struct {
	Hostname string
	Message  string
	Err      error
}

// New creates a new client
func New(conf *config.Config, machineRegex string) (*Client, error) {
	machines, err := conf.FilterMachines(machineRegex)
	if err != nil {
		return nil, err
	}
	return &Client{
		machines: machines,
	}, nil
}

func (c *Client) Run(args []string) map[string]Result {
	var wg = &sync.WaitGroup{}
	result := make(chan Result)
	for _, machine := range c.machines {
		wg.Add(1)
		go func(machine config.Machine) {
			defer wg.Done()
			response, err := sendRequest(machine.Hostname, machine.Port, strings.Join(args, " "))
			if err != nil {
				result <- Result{
					Hostname: machine.Hostname,
					Message:  "",
					Err:      err,
				}
				return
			}
			result <- Result{
				Hostname: machine.Hostname,
				Message:  response,
				Err:      nil,
			}
		}(machine)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	// combine all the results
	// TODO: handle error
	var results map[string]Result = map[string]Result{}
	for r := range result {
		if r.Err != nil {
			logrus.Error(r.Err)
		} else {
			results[r.Hostname] = r
		}
	}
	return results
}

func sendRequest(Hostname string, port string, msg string) (string, error) {
	if len(msg) == 0 {
		return "", fmt.Errorf("empty message")
	}
	socketClient, err := NewSocketClient(Hostname, port)
	if err != nil {
		return "", err
	}
	defer socketClient.Close()
	socketClient.Send(msg)
	_, response, err := socketClient.Receive()
	if err != nil {
		return "", err
	}
	return string(response), nil
}
