package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/constant"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/grep"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/server"
)

func main() {
	configPath := flag.String("config", constant.CONFIG_PATH, "path to config file")
	flag.Parse()
	config, err := config.New(*configPath)
	if err != nil {
		logrus.Fatal(err)
	}
	server, err := server.New(config)
	if err != nil {
		logrus.Fatal(err)
	}
	defer server.Close()
	handler := grep.New()
	server.Run(handler)
}
