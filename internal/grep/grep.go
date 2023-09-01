package grep

import (
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
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute grep: %w", err)
	}
	return string(out), nil
}
