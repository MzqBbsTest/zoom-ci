// Copyright 2018 zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package command

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const (
	DEFAULT_RUN_TIMEOUT = 3600
)

type Command struct {
	Cmd           string
	Timeout       time.Duration
	TerminateChan chan int
	Setpgid       bool
	command       *exec.Cmd
	stdout        bytes.Buffer
	stderr        bytes.Buffer
}


func (c *Command) Run() error {
	if err := c.command.Start(); err != nil {
		return err
	}

	errChan := make(chan error)
	go func() {
		errChan <- c.command.Wait()
		defer close(errChan)
	}()

	var err error
	select {
	case err = <-errChan:
	case <-time.After(c.Timeout):
		err = c.terminate()
		if err == nil {
			err = errors.New(fmt.Sprintf("cmd run timeout, cmd [%s], time[%v]", c.Cmd, c.Timeout))
		}
	case <-c.TerminateChan:
		err = c.terminate()
		if err == nil {
			err = errors.New(fmt.Sprintf("cmd is terminated, cmd [%s]", c.Cmd))
		}
	}
	return err
}

func (c *Command) Stderr() string {
	return strings.TrimSpace(string(c.stderr.Bytes()))
}

func (c *Command) Stdout() string {
	return strings.TrimSpace(string(c.stdout.Bytes()))
}

