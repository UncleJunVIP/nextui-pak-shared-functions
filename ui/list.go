package ui

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/UncleJunVIP/nextui-pak-shared-functions/models"
	"os"
	"os/exec"
	"strings"
)

func DisplayMinUiList(list string, format string, title string, options ...string) models.Selection {
	return DisplayMinUiListWithAction(list, format, title, "", options...)
}

func DisplayMinUiListWithAction(list string, format string, title string, actionText string, options ...string) models.Selection {
	args := []string{"--format", format, "--title", title, "--file", "-"}

	if actionText != "" {
		args = append(args, "--action-button", "X", "--action-text", actionText)
	}

	if options != nil {
		args = append(args, options...)
	}

	cmd := exec.Command("minui-list", args...)
	cmd.Env = os.Environ()
	cmd.Env = os.Environ()

	var stdoutbuf, stderrbuf bytes.Buffer
	cmd.Stdout = &stdoutbuf
	cmd.Stderr = &stderrbuf

	cmd.Stdin = strings.NewReader(list)

	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	if err := cmd.Run(); err != nil {
		return models.Selection{Code: cmd.ProcessState.ExitCode(), Error: fmt.Errorf("failed to run minui-list: %w", err)}
	}

	outValue := stdoutbuf.String()
	_ = stderrbuf.String()

	return models.Selection{Value: outValue, Code: cmd.ProcessState.ExitCode()}
}
