package types

// file represents a single file.
type file struct {
	Name string
	Body string
}

// Request represents incoming request config.
type Request struct {
	Language string `json:"language"`
	Files    []file `json:"files"`
}
