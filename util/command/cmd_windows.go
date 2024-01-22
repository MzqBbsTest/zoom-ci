// +build windows

package command

import (
	"golang.org/x/sys/windows"
	"os/exec"
	"time"
)

func (c *Command) terminate() error {
	if c.Setpgid {
		handle := windows.Handle(c.command.Process.Pid)
		return windows.TerminateProcess(handle, 1)
	} else {
		handle := windows.Handle(c.command.Process.Pid)
		return windows.TerminateProcess(handle, 1)
	}
}

func NewCmd(c *Command) (*Command, error) {
	if c.Timeout == 0*time.Second {
		c.Timeout = DEFAULT_RUN_TIMEOUT * time.Second
	}
	if c.TerminateChan == nil {
		c.TerminateChan = make(chan int)
	}
	cmd := exec.Command("cmd", "/C", c.Cmd)
	cmd.Stderr = &c.stderr
	cmd.Stdout = &c.stdout
	c.command = cmd

	return c, nil
}
