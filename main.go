package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Config represents an object received on stdin containing configuration.
type Config struct {
	Files []struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
}

func main() {
	config, err := waitForConfig()
	if err != nil {
		panic(err)
	}

	if err := createFiles(config); err != nil {
		panic(err)
	}

	out, err := run()
	if err != nil {
		panic(err)
	}

	fmt.Print(string(out))
}

// waitForConfig waits for a config to be passed on stdin.
// Returns unmarshalled config and any error encountered.
func waitForConfig() (*Config, error) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	b := scanner.Bytes()

	var config Config
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// createFiles creates all input files specified in config.
// Returns any error encountered.
func createFiles(config *Config) error {
	for _, f := range config.Files {
		file, err := os.Create(f.Name)
		if err != nil {
			return err
		}

		if _, err := file.Write([]byte(f.Content)); err != nil {
			return err
		}
	}

	return nil
}

// execute runs the program.
// Returns the result and any error encountered.
func run() ([]byte, error) {
	out, err := exec.Command("go", "run", "./out/main.go").Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
