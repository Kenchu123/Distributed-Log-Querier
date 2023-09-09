package basic_test

import (
	"sync"
	"testing"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
)

func TestParallelClient(t *testing.T) {
	conf, err := config.New("./config.yml")
	if err != nil {
		t.Fatal(err)
	}
	machineRegex := "01"

	// test CLIENT_NUM client
	CLIENT_NUM := 10
	args := []string{"-c", "PUT", "test/basic/logs/machine.i.log"}

	// Build expected
	expected := []map[string]client.Result{}
	for i := 0; i < CLIENT_NUM; i++ {
		expected = append(expected, map[string]client.Result{
			"fa23-cs425-8701.cs.illinois.edu": {
				Hostname: "fa23-cs425-8701.cs.illinois.edu",
				Message:  "191\n",
			},
		})
	}

	// Start testing
	testClients := []*client.Client{}
	for i := 0; i < CLIENT_NUM; i++ {
		testClient, err := client.New(conf, machineRegex)
		if err != nil {
			t.Fatal(err)
		}
		testClients = append(testClients, testClient)
	}

	var wg = &sync.WaitGroup{}
	outputChan := make(chan map[string]client.Result)
	for _, testClient := range testClients {
		wg.Add(1)
		go func(testClient *client.Client) {
			defer wg.Done()
			response := testClient.Run(args)
			outputChan <- response
		}(testClient)
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	output := []map[string]client.Result{}
	for out := range outputChan {
		output = append(output, out)
	}

	if len(output) != len(expected) {
		t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
	}
	for i := 0; i < CLIENT_NUM; i++ {
		if isEqual(output[i], expected[i]) == false {
			t.Errorf("Output %+v is not equal to Expected %+v", output, expected)
		}
	}
}