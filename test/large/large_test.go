package large_test

import (
	"reflect"
	"testing"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
)

func isEqual(a map[string]client.Result, b map[string]client.Result) bool {
	return reflect.DeepEqual(a, b)
}

func TestLargeFrequent(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "0[1-2]",
		MachineILog:       true,
		MachineILogFolder: "test/large/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "283553\n",
		},
		"fa23-cs425-8702.cs.illinois.edu": {
			Hostname: "fa23-cs425-8702.cs.illinois.edu",
			Message:  "267938\n",
		},
	}
	expectedTotalLine := 551491
	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "/"}
	output, totalLine := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
	if totalLine != expectedTotalLine {
		t.Errorf("Output %c is not equal to Expected %c", totalLine, expectedTotalLine)
	}
}

func TestLargeSomewhatFrequent(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "0[1-2]",
		MachineILog:       true,
		MachineILogFolder: "test/large/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "56879\n",
		},
		"fa23-cs425-8702.cs.illinois.edu": {
			Hostname: "fa23-cs425-8702.cs.illinois.edu",
			Message:  "53604\n",
		},
	}
	expectedTotalLine := 110483
	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "PUT"}
	output, totalLine := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
	if totalLine != expectedTotalLine {
		t.Errorf("Output %c is not equal to Expected %c", totalLine, expectedTotalLine)
	}
}

func TestLargeInfrequent(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "0[1-2]",
		MachineILog:       true,
		MachineILogFolder: "test/large/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "9348\n",
		},
		"fa23-cs425-8702.cs.illinois.edu": {
			Hostname: "fa23-cs425-8702.cs.illinois.edu",
			Message:  "9018\n",
		},
	}
	expectedTotalLine := 18366
	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "iPod"}
	output, totalLine := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
	if totalLine != expectedTotalLine {
		t.Errorf("Output %c is not equal to Expected %c", totalLine, expectedTotalLine)
	}
}

func TestLargeRegex(t *testing.T) {
	opts := &client.Options{
		ConfigPath:        "./config.yml",
		MachineRegex:      "0[1-2]",
		MachineILog:       true,
		MachineILogFolder: "test/large/logs",
	}
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "163817\n",
		},
		"fa23-cs425-8702.cs.illinois.edu": {
			Hostname: "fa23-cs425-8702.cs.illinois.edu",
			Message:  "151487\n",
		},
	}
	expectedTotalLine := 315304
	testClient, err := client.New(conf, opts)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "[I-J]"}
	output, totalLine := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
	if totalLine != expectedTotalLine {
		t.Errorf("Output %c is not equal to Expected %c", totalLine, expectedTotalLine)
	}
}
