package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type Client interface {
	Clone(url, dest string) error
	Pull(repoPath string) error
	GetLatestCommit(repoPath string) (string, error)
	IsCloned(repoPath string) bool
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) Clone(url, dest string) error {
	cmd := exec.Command("git", "clone", url, dest)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git clone failed: %w\n%s", err, output)
	}
	return nil
}

func (c *client) Pull(repoPath string) error {
	cmd := exec.Command("git", "-C", repoPath, "pull")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git pull failed: %w\n%s", err, output)
	}
	return nil
}

func (c *client) GetLatestCommit(repoPath string) (string, error) {
	cmd := exec.Command("git", "-C", repoPath, "rev-parse", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get commit: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func (c *client) IsCloned(repoPath string) bool {
	cmd := exec.Command("git", "-C", repoPath, "rev-parse", "--git-dir")
	return cmd.Run() == nil
}
