package git

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

// getCachePath returns the cache path for a Git repository
func getCachePath(repoURL string) string {
	// Hash the repo URL to create a unique cache directory
	hash := sha256.Sum256([]byte(repoURL))
	hashStr := fmt.Sprintf("%x", hash[:8])

	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return filepath.Join(homeDir, ".kite", "cache", "registry", hashStr)
}

// isLocalPath checks if a path is a local filesystem path
func isLocalPath(path string) bool {
	return filepath.IsAbs(path) ||
		len(path) > 0 && path[0] == '.' ||
		(!hasPrefix(path, "http://") &&
			!hasPrefix(path, "https://") &&
			!hasPrefix(path, "git@"))
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
