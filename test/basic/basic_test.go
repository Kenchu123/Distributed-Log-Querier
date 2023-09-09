package basic_test

import (
	"reflect"
	"testing"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
)

func isEqual(a map[string]client.Result, b map[string]client.Result) bool {
	return reflect.DeepEqual(a, b)
}

func TestOneMachine(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "01",
		MachineILog:       true,
		MachineILogFolder: "test/basic/logs",
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

func TestMachinesFrequent(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "0[1-5]",
		MachineILog:       true,
		MachineILogFolder: "test/basic/logs",
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
		"fa23-cs425-8702.cs.illinois.edu": {
			Hostname: "fa23-cs425-8702.cs.illinois.edu",
			Message:  "213\n",
		},
		"fa23-cs425-8703.cs.illinois.edu": {
			Hostname: "fa23-cs425-8703.cs.illinois.edu",
			Message:  "203\n",
		},
		"fa23-cs425-8704.cs.illinois.edu": {
			Hostname: "fa23-cs425-8704.cs.illinois.edu",
			Message:  "194\n",
		},
		"fa23-cs425-8705.cs.illinois.edu": {
			Hostname: "fa23-cs425-8705.cs.illinois.edu",
			Message:  "200\n",
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

func TestMachinesInfrequent(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "0[1-5]",
		MachineILog:       true,
		MachineILogFolder: "test/basic/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "30\n",
		},
		"fa23-cs425-8702.cs.illinois.edu": {
			Hostname: "fa23-cs425-8702.cs.illinois.edu",
			Message:  "33\n",
		},
		"fa23-cs425-8703.cs.illinois.edu": {
			Hostname: "fa23-cs425-8703.cs.illinois.edu",
			Message:  "31\n",
		},
		"fa23-cs425-8704.cs.illinois.edu": {
			Hostname: "fa23-cs425-8704.cs.illinois.edu",
			Message:  "33\n",
		},
		"fa23-cs425-8705.cs.illinois.edu": {
			Hostname: "fa23-cs425-8705.cs.illinois.edu",
			Message:  "28\n",
		},
	}

	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "iPod"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}

func TestMachinesRegex(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "0[1-5]",
		MachineILog:       true,
		MachineILogFolder: "test/basic/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "463\n",
		},
		"fa23-cs425-8702.cs.illinois.edu": {
			Hostname: "fa23-cs425-8702.cs.illinois.edu",
			Message:  "476\n",
		},
		"fa23-cs425-8703.cs.illinois.edu": {
			Hostname: "fa23-cs425-8703.cs.illinois.edu",
			Message:  "462\n",
		},
		"fa23-cs425-8704.cs.illinois.edu": {
			Hostname: "fa23-cs425-8704.cs.illinois.edu",
			Message:  "466\n",
		},
		"fa23-cs425-8705.cs.illinois.edu": {
			Hostname: "fa23-cs425-8705.cs.illinois.edu",
			Message:  "446\n",
		},
	}

	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "[I-J]"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}
