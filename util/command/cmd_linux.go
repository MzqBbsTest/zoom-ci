// +build linux darwin

package command

import "golang.org/x/sys/unix"

func (c *Command) terminate() error {
	if c.Setpgid {
		unix.Syscall()
		return syscall.Kill(-c.command.Process.Pid, syscall.SIGKILL)
	} else {
		return syscall.Kill(c.command.Process.Pid, syscall.SIGKILL)
	}
}


func NewCmd(c *Command) (*Command, error) {
	if c.Timeout == 0*time.Second {
		c.Timeout = DEFAULT_RUN_TIMEOUT * time.Second
	}
	if c.TerminateChan == nil {
		c.TerminateChan = make(chan int)
	}
	cmd := exec.Command("/bin/bash", "-c", c.Cmd)
	if c.Setpgid {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}
	cmd.Stderr = &c.stderr
	cmd.Stdout = &c.stdout
	c.command = cmd

	return c, nil
}
