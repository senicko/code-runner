package types

type Response struct {
	Stdout    string `json:"stdout"`
	Stderr    string `json:"stderr"`
	ExecError string `json:"execError"`
	ExitCode  int    `json:"exitCode"`
}
