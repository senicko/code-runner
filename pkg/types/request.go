package types

type File struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type Config struct {
  Language string `json:"language"`
}

type Request struct {
  Config Config `json:"config"`
	Language string `json:"language"`
	Stdin    string `json:"stdin"`
	Files    []File `json:"files"`
}
