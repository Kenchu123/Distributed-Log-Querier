package config

import (
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Config is the configuration for the servers
type Config struct {
	Machines []Machine `yaml:"machines"`
}

// Machine is the configuration for a single server
type Machine struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
}

// New reads the configuration file and returns the configuration
func New(path string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// FilterMachines filters the machines based on the regex
func (c *Config) FilterMachines(regex string) ([]Machine, error) {
	var machines []Machine
	reg, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	for _, machine := range c.Machines {
		if reg.MatchString(machine.Hostname) {
			machines = append(machines, machine)
		}
	}
	return machines, nil
}
