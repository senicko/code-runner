package types

type File struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type Request struct {
	Stdin string `json:"stdin"`
	Files []File `json:"files"`
}
