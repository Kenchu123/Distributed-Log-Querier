package client

import (
	"testing"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/constant"
)

// TestBasic tests the basic functionality of the client
func TestBasic(t *testing.T) {
	// build config
	conf, err := config.New("../../" + constant.CONFIG_PATH)
	if err != nil {
		t.Fatal(err)
	}
	// build machineRegex
	machineRegex := "0[1-5]"
	// build client
	client, err := New(conf, machineRegex)
	if err != nil {
		t.Fatal(err)
	}
	// build args
	args := []string{"PUT", "logs/machine.i.log"}
	// TODO: test response
	result := client.Run(args)
	t.Log(result)
}
