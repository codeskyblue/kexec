package kexec

import (
	"os"
	"os/exec"
	"strconv"
)

func Command(name string, arg ...string) *Process {
	return &KCommand{
		Cmd: exec.Command(name, arg...),
	}
}

func CommandString(command string) *Process {
	cmd := exec.Command("cmd", "/c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return &Process{
		Cmd: cmd,
	}
}

func (p *KCommand) Terminate(sig os.Signal) (err error) {
	if p.Process == nil {
		return nil
	}
	pid := p.Process.Pid
	c := exec.Command("taskkill", "/t", "/f", "/pid", strconv.Itoa(pid))
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
