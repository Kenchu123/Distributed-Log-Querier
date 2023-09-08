package unit_tests

import (
	"fmt"
	"testing"
	"os/exec"

	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/client"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/config"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/constant"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp1/internal/grep"
)

// TestBasic tests the basic functionality of the grep function
func TestGrepFrequent(t *testing.T) {
	t.Log("Testing Frequent Pattarns")
	var opts struct {
		ConfigPath   string `long:"config" description:"path to config file"`
		MachineRegex string `long:"machine" description:"regex to match machine names" default:".*"`
		Help         bool   `short:"h" long:"help" description:"show this help message"`
	}

	if opts.Help {
		fmt.Println(grep.Help())
		return
	}

	if len(opts.ConfigPath) == 0 {
		opts.ConfigPath = constant.CONFIG_TEST
	}

	// Read config file
	conf, err := config.New(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	// Create client
	client, err := client.New(conf, opts.MachineRegex)
	if err != nil {
		t.Fatal(err)
	}

	// build args
	args := []string{"-c", "PUT", "logs/machine.i.log"}

	// run ds-grep with args
	ds_result := client.Run(args)
	fmt.Print("ds_grep results: ", ds_result)

	// run system grep with args 
	sys_result, err := exec.Command("grep", "-c", "PUT", "../logs/machine.i.log").Output()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print("system grep results: " + "dev:" + string(sys_result))

	if (("dev:" + string(sys_result)) != ds_result) {
		t.Fatalf("The results of ds_grep is not equal to the system grep")
	}

}


