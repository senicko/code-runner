package runner

import "os/exec"

type Config struct {
	Cmd *exec.Cmd
}

func NewConfig(lang, entry string) *Config {
	config := &Config{}

	switch lang {
	case "golang", "go":
		config.Cmd = exec.Command("go", "run", entry)
	}

	return config
}
