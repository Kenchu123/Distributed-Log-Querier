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
	var opts client.Options
	args, err := flags.NewParser(&opts, flags.IgnoreUnknown).ParseArgs(os.Args[1:])
	if err != nil {
		logrus.Fatal(err)
	}

	if len(args) == 0 || opts.Help {
		fmt.Println(grep.Help())
		os.Exit(0)
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
	client, err := client.New(conf, &opts)
	if err != nil {
		logrus.Fatal(err)
	}
	results := client.Run(args)

	// Print result to terminal
	for hostname, result := range results {
		fmt.Printf("%s:%s\n", hostname, result.Message)
	}
}
