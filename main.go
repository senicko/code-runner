package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/senicko/bee/pkg/types"
)

func main() {
	config, err := waitForConfig()
	if err != nil {
		panic(err)
	}

	if err := createFiles(config); err != nil {
		panic(err)
	}

	result, err := run()
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(output))
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

// execute runs the program.
// Returns the result and any error encountered.
func run() (*types.Result, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("go", "run", "./files/main.go")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	var e *exec.ExitError
	if err := cmd.Run(); err != nil && !errors.As(err, &e) {
		return nil, err
	}

	return &types.Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: cmd.ProcessState.ExitCode(),
	}, nil
}
