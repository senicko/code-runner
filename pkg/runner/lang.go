package runner

import "os/exec"

type Config struct {
	BuildChain []*exec.Cmd
	Exec       *exec.Cmd
}

func NewConfig(lang string) *Config {
	config := &Config{}

	switch lang {
	case "golang", "go":
		config.BuildChain = []*exec.Cmd{
      exec.Command("go", "build", "-o", "./bin/main", "./files/main.go"),
		}
		config.Exec = exec.Command("./bin/main")
	case "cpp", "c++":
		config.BuildChain = []*exec.Cmd{
			exec.Command("g++", "./files/main.cpp"),
		}
		config.Exec = exec.Command("./a.out")
	case "c":
		config.BuildChain = []*exec.Cmd{
			exec.Command("gcc", "./files/main.c"),
		}
		config.Exec = exec.Command("./a.out")
	}

	return config
}
