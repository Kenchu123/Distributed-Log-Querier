package basic_test

import (
	"sync"
	"testing"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
)

func TestParallelClient(t *testing.T) {
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
	// test CLIENT_NUM client
	CLIENT_NUM := 10
	args := []string{"-c", "PUT"}

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
	expectedTotalLine := 191
	// Start testing
	testClients := []*client.Client{}
	for i := 0; i < CLIENT_NUM; i++ {
		testClient, err := client.New(conf, opts)
		if err != nil {
			t.Fatal(err)
		}
		testClients = append(testClients, testClient)
	}

	var wg = &sync.WaitGroup{}
	outputChan := make(chan map[string]client.Result)
	var totalLine int
	for _, testClient := range testClients {
		wg.Add(1)
		go func(testClient *client.Client) {
			defer wg.Done()
			response, n := testClient.Run(args)
			totalLine = n
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
		if totalLine != expectedTotalLine {
			t.Errorf("Output %c is not equal to Expected %c", totalLine, expectedTotalLine)
		}
	}
}
