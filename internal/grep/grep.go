package grep

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

// Handler handles `grep` executation
type Handler struct{}

// New creates a new handler for grep
func New() *Handler {
	return &Handler{}
}

// Handle execute 'grep' command locally and returns the response
func (h *Handler) Handle(args []string) (string, error) {
	logrus.Printf("grep %v\n", args)

	cmd := exec.Command("grep", args...)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				// ignore exit status 1
				logrus.Info("grep returned exit status 1, no match found")
			} else {
				return "", fmt.Errorf("failed to execute grep: %w\n%s", err, stdErr.String())
			}
		}
	}
	return stdOut.String(), nil
}
