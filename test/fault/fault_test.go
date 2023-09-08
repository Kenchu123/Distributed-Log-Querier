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
	conf, err := config.New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}
	machineRegex := "01|failure"
	expected := map[string]client.Result{
		"fa23-cs425-8701.cs.illinois.edu": {
			Hostname: "fa23-cs425-8701.cs.illinois.edu",
			Message:  "191\n",
		},
	}

	testClient, err := client.New(conf, machineRegex)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "PUT", "test/fault/logs/machine.i.log"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}

func TestOnlyNoneExistingMachine(t *testing.T) {
	conf, err := config.New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}
	machineRegex := "failure"
	expected := map[string]client.Result{}

	testClient, err := client.New(conf, machineRegex)
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"-c", "PUT", "test/fault/logs/machine.i.log"}
	output := testClient.Run(args)

	if isEqual(output, expected) == false {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
}

