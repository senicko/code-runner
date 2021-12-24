package types

// Request represents incoming request config.
type Request struct {
	Files []struct {
		Name string
		Body string
	} `json:"files"`
}
