package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/senicko/bee/pkg/runner"
	"github.com/senicko/bee/pkg/types"
)

func main() {
	output, err := start()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	}

	fmt.Print(output)
}

func start() (string, error) {
	config, err := waitForConfig()
	if err != nil {
		return "", err
	}

	if err := createFiles(config); err != nil {
		return "", err
	}

	runnerCofig := runner.NewConfig(config.Language, "./files/main.go")
	result, err := run(runnerCofig)
	if err != nil {
		return "", err
	}

	output, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// waitForConfig waits for a config to be passed on stdin.
// Returns unmarshalled config and any error encountered.
func waitForConfig() (*types.Request, error) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	b := scanner.Bytes()

	var request types.Request
	if err := json.Unmarshal(b, &request); err != nil {
		return nil, err
	}

	return &request, nil
}

// createFiles creates all input files specified in request.
// Returns any error encountered.
func createFiles(request *types.Request) error {
	for _, f := range request.Files {
		file, err := os.Create("./files/" + f.Name)
		if err != nil {
			return err
		}

		if _, err := file.Write([]byte(f.Body)); err != nil {
			return err
		}
	}

	return nil
}

// run runs the submitted code.
// Returns the result and any error encountered.
func run(config *runner.Config) (*types.Result, error) {
	var stdout, stderr bytes.Buffer

	config.Cmd.Stdout = &stdout
	config.Cmd.Stderr = &stderr

	var e *exec.ExitError
	if err := config.Cmd.Run(); err != nil && !errors.As(err, &e) {
		return nil, err
	}

	return &types.Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: config.Cmd.ProcessState.ExitCode(),
	}, nil
}
