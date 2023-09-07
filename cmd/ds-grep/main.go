package main

import (
	"fmt"
	"os"

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

	// Create client
	client, err := client.New(conf, opts.MachineRegex)
	if err != nil {
		logrus.Fatal(err)
	}
	// TODO: handle response
	result := client.Run(args)
	fmt.Print(result)
}
