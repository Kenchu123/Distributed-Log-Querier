package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/constant"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/grep"
)

func main() {
	var opts struct {
		ConfigPath   string `long:"config" description:"path to config file"`
		MachineRegex string `long:"machine" description:"regex to match machine names" default:".*"`
		Help         bool   `short:"h" long:"help" description:"show this help message"`
	}
	args, err := flags.NewParser(&opts, flags.IgnoreUnknown).ParseArgs(os.Args[1:])
	if err != nil {
		logrus.Fatal(err)
	}

	if opts.Help {
		fmt.Println(grep.Help())
		return
	}

	if len(opts.ConfigPath) == 0 {
		opts.ConfigPath = constant.CONFIG_PATH
	}

	// Read config file
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// Filter machines based on regex
	machines, err := conf.FilterMachines(opts.MachineRegex)
	if err != nil {
		logrus.Fatal(err)
	}

	// Send request to each machine
	var wg = &sync.WaitGroup{}
	for _, machine := range machines {
		wg.Add(1)
		go func(machine config.Machine) {
			defer wg.Done()
			response, err := sendRequest(machine.Hostname, machine.Port, strings.Join(args, " "))
			if err != nil {
				logrus.Errorf("failed to send request to %s:%s: %v\n", machine.Hostname, machine.Port, err)
				return
			}
			logrus.Printf("Response from %s:%s:\n%s\n", machine.Hostname, machine.Port, response)
		}(machine)
	}
	wg.Wait()
	// TODO: handle response
}

func sendRequest(Hostname string, port string, msg string) (string, error) {
	if len(msg) == 0 {
		return "", fmt.Errorf("empty message")
	}
	client, err := client.New(Hostname, port)
	if err != nil {
		return "", err
	}
	defer client.Close()
	client.Send(msg)
	_, response, err := client.Receive()
	if err != nil {
		return "", err
	}
	return string(response), nil
}
