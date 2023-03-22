package activity

import (
	"fmt"
	"os/exec"
)

type Shell struct {
	name string
	path string
	args []string
}

func NewShell(name string, path string, args ...string) *Shell {
	return &Shell{
		name: name,
		path: path,
		args: args,
	}
}

func (s *Shell) GetName() string {
	return s.name
}

func (s *Shell) Execute() error {
	fmt.Println("Executing Activity Shell ", s.name)
	
	if output, err := exec.Command(s.path, s.args...).Output(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}

	return nil
}
