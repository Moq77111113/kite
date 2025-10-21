package git

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Client interface {
	Clone(url, dest string) error
	Pull(repoPath string) error
	GetLatestCommit(repoPath string) (string, error)
	GetLastModifiedDate(repoPath, path string) (time.Time, error)
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

func (c *client) GetLastModifiedDate(repoPath, path string) (time.Time, error) {
	cmd := exec.Command("git", "-C", repoPath, "log", "-1", "--format=%ct", "--", path)
	output, err := cmd.Output()
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get last modified date: %w", err)
	}

	timestampStr := strings.TrimSpace(string(output))
	if timestampStr == "" {
		return time.Time{}, fmt.Errorf("no commits found for path: %s", path)
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse timestamp: %w", err)
	}

	return time.Unix(timestamp, 0), nil
}

func (c *client) IsCloned(repoPath string) bool {
	cmd := exec.Command("git", "-C", repoPath, "rev-parse", "--git-dir")
	return cmd.Run() == nil
}
