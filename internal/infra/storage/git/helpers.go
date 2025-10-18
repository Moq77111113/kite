package git

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

func GetCachePath(repoURL string) string {
	hash := sha256.Sum256([]byte(repoURL))
	hashStr := fmt.Sprintf("%x", hash[:8])

	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return filepath.Join(homeDir, ".kite", "cache", "registry", hashStr)
}

func IsLocalPath(path string) bool {
	return filepath.IsAbs(path) ||
		len(path) > 0 && path[0] == '.' ||
		(!hasPrefix(path, "http://") &&
			!hasPrefix(path, "https://") &&
			!hasPrefix(path, "git@"))
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
