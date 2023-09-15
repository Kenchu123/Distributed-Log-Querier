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
	options  *Options
}

type Options struct {
	ConfigPath        string `long:"config" description:"path to config file"`
	MachineRegex      string `long:"machine" description:"regex to match machine names" default:".*"`
	MachineILog       bool   `long:"machine-ilog" description:"append machine.$i.log to grep file path"`
	MachineILogFolder string `long:"machine-ilog-folder" description:"folder to store machine.$i.log" default:"logs"`
	Help              bool   `short:"h" long:"help" description:"show this help message"`
}

type Result struct {
	Hostname string
	Message  string
	Err      error
}

// New creates a new client
func New(conf *config.Config, opts *Options) (*Client, error) {
	machines, err := conf.FilterMachines(opts.MachineRegex)
	if err != nil {
		return nil, err
	}
	return &Client{
		machines: machines,
		options:  opts,
	}, nil
}

// Run runs the client
func (c *Client) Run(args []string) (map[string]Result, int) {
	var wg = &sync.WaitGroup{}
	result := make(chan Result)
	for _, machine := range c.machines {
		wg.Add(1)
		go func(machine config.Machine) {
			defer wg.Done()
			response, err := sendRequest(machine.Hostname, machine.Port, buildArgs(args, c.options, machine))
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

	totalLine := 0
	// Count total number of log lines
	for _, r := range results {
		c := false
		// Iterate through args to check if -c or -cE is present
		for _, arg := range args {
			if arg == "-c" || arg == "-cE" {
				c = true
				break
			}
		}
		if (c) {
			var num int
			_, err := fmt.Sscan(r.Message, &num)
			if err != nil {
				logrus.Error(err)
			}
			totalLine += num
			totalLine -= strings.Count(r.Message, "\n\n")
		} else {
			totalLine += strings.Count(r.Message, "\n")
			totalLine -= strings.Count(r.Message, "\n\n")
		}
	}
	return results, totalLine
}

func buildArgs(args []string, opts *Options, machine config.Machine) string {
	if opts.MachineILog {
		args = append(args, fmt.Sprintf("%s/machine.%s.log", opts.MachineILogFolder, machine.ID))
	}
	return strings.Join(args, " ")
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
