package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	BuildChain []*exec.Cmd
	Exec       *exec.Cmd
}

type rawConfig struct {
	BuildChain []string `json:"buildChain"`
	Exec       string   `json:"exec"`
}

// LoadConfig loads runner configuration from runner.config.json file.
func LoadConfig() (*Config, error) {
	configFile, err := os.Open("./runner.config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	bytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var raw rawConfig
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return nil, err
	}

	var config Config

	buildChain, err := parseBuildChain(raw.BuildChain)
	if err != nil {
		return nil, err
	}
	config.BuildChain = buildChain

	execCmd, err := stringToCommand(raw.Exec)
	if err != nil {
		return nil, err
	}
	config.Exec = execCmd

	return &config, nil
}

func stringToCommand(cmd string) (*exec.Cmd, error) {
	splitted := strings.Split(cmd, " ")

	if len(splitted) == 0 {
		return nil, fmt.Errorf("invalid build command")
	}

	return exec.Command(splitted[0], splitted[1:]...), nil
}

func parseBuildChain(rawCmds []string) ([]*exec.Cmd, error) {
	var buildChain []*exec.Cmd

	for _, rawCmd := range rawCmds {
		buildCmd, err := stringToCommand(rawCmd)
		if err != nil {
			return nil, err
		}
		buildChain = append(buildChain, buildCmd)
	}

	return buildChain, nil
}
