package fault_test

import (
	"reflect"
	"testing"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
)

func isEqual(a map[string]client.Result, b map[string]client.Result) bool {
	return reflect.DeepEqual(a, b)
}

func TestNoneExistingWithValidMachine(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "01|failure",
		MachineILog:       true,
		MachineILogFolder: "test/fault/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "191\n",
		},
	}

	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "PUT"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}

func TestOnlyNoneExistingMachine(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "failure",
		MachineILog:       true,
		MachineILogFolder: "test/fault/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{}

	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "PUT"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}
