package types

// Request represents incoming request config.
type Request struct {
	Files []struct {
		Name string `json:"name"`
		Body string `json:"body"`
	} `json:"files"`
}
