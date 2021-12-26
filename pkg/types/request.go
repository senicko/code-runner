package types

type Request struct {
	Language string `json:"language"`
	Files    []struct {
		Name string `json:"name"`
		Body string `json:"body"`
	} `json:"files"`
}
