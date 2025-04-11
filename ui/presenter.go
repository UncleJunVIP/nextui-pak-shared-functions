package ui

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ShowMessage(message string, timeout string) (int, error) {
	return ShowMessageWithOptions(message, timeout)
}

func ShowMessageWithOptions(message string, timeout string, options ...string) (int, error) {
	args := []string{"--message", message, "--timeout", timeout}

	if options != nil {
		args = append(args, options...)
	}

	cmd := exec.Command("minui-presenter", args...)

	var stdoutbuf, stderrbuf bytes.Buffer
	cmd.Stdout = &stdoutbuf
	cmd.Stderr = &stderrbuf

	err := cmd.Run()

	if err != nil && cmd.ProcessState.ExitCode() == 1 {
		return cmd.ProcessState.ExitCode(),
			fmt.Errorf("issue running minui-presenter: %w\n%w", err, stderrbuf.String())
	}

	return cmd.ProcessState.ExitCode(), nil
}
