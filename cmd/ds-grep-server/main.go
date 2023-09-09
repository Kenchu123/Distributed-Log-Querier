package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/grep"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/server"
)

func main() {
	var opts struct {
		Port string `short:"p" long:"port" description:"port to listen on" default:"7122"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		logrus.Fatal(err)
	}
	server, err := server.New(opts.Port)
	if err != nil {
		logrus.Fatal(err)
	}
	defer server.Close()
	handler := grep.New()
	server.Run(handler)
}
