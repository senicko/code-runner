package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/senicko/code-runner/pkg/runner"
	"github.com/senicko/code-runner/pkg/types"
)

func main() {
	result, err := start()

	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}

		jsonResponse, _ := json.Marshal(response)
		fmt.Fprintln(os.Stderr, string(jsonResponse))
	} else {
		// TODO: handle the error.
		output, _ := json.Marshal(result)
		fmt.Print(string(output))
	}
}

// Starts the run process.
func start() (*types.Response, error) {
	request, err := waitForRequest()
	if err != nil {
		return nil, err
	}

	if err := createFiles(request); err != nil {
		return nil, err
	}

	runnerConfig := runner.NewConfig(request.Config.Language)

	var stdout, stderr bytes.Buffer

	buildError, err := build(runnerConfig.BuildChain, &stderr)

	if err != nil {
		return nil, err
	}

	if buildError != nil {
		return buildError, nil
	}

	return run(runnerConfig, &stdout, &stderr, request.Stdin), nil
}

// Waits for a request to be passed on stdin.
func waitForRequest() (*types.Request, error) {
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

// Creates all input files specified in request.
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

// build the app.
func build(chain []*exec.Cmd, stderr *bytes.Buffer) (*types.Response, error) {
	for _, cmd := range chain {
		cmd.Stderr = stderr

		var e *exec.ExitError
		if err := cmd.Run(); err != nil && !errors.As(err, &e) {
			return nil, err
		}

		if exitCode := cmd.ProcessState.ExitCode(); stderr.Len() != 0 && exitCode > 0 {
			return &types.Response{
				Stderr:   stderr.String(),
				ExitCode: exitCode,
			}, nil
		}
	}

	return nil, nil
}

// run the submitted code.
func run(config *runner.Config, stdout, stderr *bytes.Buffer, stdin string) *types.Response {
	config.Exec.Stdout = stdout
	config.Exec.Stderr = stderr
	config.Exec.Stdin = strings.NewReader(stdin)

	if err := config.Exec.Run(); err != nil {
    return &types.Response{Stderr: stderr.String(), ExecError: err.Error(), ExitCode: config.Exec.ProcessState.ExitCode()}
	}

	return &types.Response{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: config.Exec.ProcessState.ExitCode(),
	}
}
