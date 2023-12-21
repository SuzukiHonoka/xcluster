package executor

import (
	"io"
	"os/exec"
	"runtime"
	"strings"
)

type Executor struct {
	Cmd        string
	Privileged bool
}

func NewExecutor(cmd string, privileged bool) *Executor {
	return &Executor{
		Cmd:        cmd,
		Privileged: privileged,
	}
}

func (e *Executor) Execute() (string, error) {
	return e.ExecuteString()
}

func (e *Executor) ExecuteString() (string, error) {
	var out strings.Builder
	if err := e.ExecuteStream(&out); err != nil {
		return "", err
	}
	return out.String(), nil
}

func (e *Executor) ExecuteStream(writer io.Writer) error {
	// default run as root user
	var s string
	// sudo only works on non-windows OS
	if runtime.GOOS == "windows" || e.Privileged {
		s = e.Cmd
	} else {
		s = "sudo -u nobody " + e.Cmd
	}
	// enable capability of using variables
	s = "\"" + s + "\""
	cmd := exec.Command(DefaultShell, DefaultShellArg, s)
	cmd.Stdout = writer
	cmd.Stderr = writer
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
