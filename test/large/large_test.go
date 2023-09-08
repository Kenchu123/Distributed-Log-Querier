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
	conf, err := config.New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}
	machineRegex := "0[1-2]"
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

	testClient, err := client.New(conf, machineRegex)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "PUT", "test/large/logs/machine.i.log"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}

func TestLargeInfrequent(t *testing.T) {
	conf, err := config.New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}
	machineRegex := "0[1-2]"
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

	testClient, err := client.New(conf, machineRegex)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "iPod", "test/large/logs/machine.i.log"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}

func TestLargeRegex(t *testing.T) {
	conf, err := config.New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}
	machineRegex := "0[1-2]"
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

	testClient, err := client.New(conf, machineRegex)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "[I-J]", "test/large/logs/machine.i.log"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}

