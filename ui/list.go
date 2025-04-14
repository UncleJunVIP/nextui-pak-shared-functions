package ui

import (
	"bytes"
	"errors"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/models"
	"os"
	"os/exec"
	"strings"
)

type Entries interface {
	Values() []string
}

func DisplayList(entries Entries, title string, actionText string, options ...string) (models.ListSelection, error) {
	args := []string{"--format", "text", "--title", title, "--file", "-"}

	if actionText != "" {
		args = append(args, "--action-button", "X", "--action-text", actionText)
	}

	if options != nil {
		args = append(args, options...)
	}

	cmd := exec.Command("minui-list", args...)
	cmd.Env = os.Environ()

	var stdoutbuf, stderrbuf bytes.Buffer
	cmd.Stdout = &stdoutbuf
	cmd.Stderr = &stderrbuf

	newlineList := strings.Join(entries.Values(), "\n")

	cmd.Stdin = strings.NewReader(newlineList)

	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	if err := cmd.Run(); err != nil {
		return models.ListSelection{
			Value:    "",
			ExitCode: cmd.ProcessState.ExitCode(),
		}, err
	}

	out := strings.TrimSpace(stdoutbuf.String())

	return models.ListSelection{
		Value:    out,
		ExitCode: cmd.ProcessState.ExitCode(),
	}, nil
}
