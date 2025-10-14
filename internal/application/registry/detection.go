package registry

import (
	"strings"
)

// RegistryType represents the type of registry
type RegistryType int

const (
	RegistryTypeUnknown RegistryType = iota
	RegistryTypeHTTP
	RegistryTypeGit
	RegistryTypeLocal
)

// DetectRegistryType determines the registry type from a URL/path
func DetectRegistryType(registry string) RegistryType {
	// HTTP/HTTPS URLs
	if strings.HasPrefix(registry, "http://") || strings.HasPrefix(registry, "https://") {
		// Check if it's a Git hosting service URL
		if isGitHostingURL(registry) {
			return RegistryTypeGit
		}
		return RegistryTypeHTTP
	}

	// Git URLs (SSH)
	if strings.HasPrefix(registry, "git@") {
		return RegistryTypeGit
	}

	// GitHub/GitLab shorthand
	if strings.Contains(registry, "github.com/") || strings.Contains(registry, "gitlab.com/") {
		return RegistryTypeGit
	}

	// Local file path
	if strings.HasPrefix(registry, "/") || strings.HasPrefix(registry, "./") || strings.HasPrefix(registry, "../") {
		return RegistryTypeLocal
	}

	return RegistryTypeUnknown
}

// isGitHostingURL checks if a URL is from a Git hosting service
func isGitHostingURL(url string) bool {
	gitHosts := []string{
		"github.com",
		"gitlab.com",
		"bitbucket.org",
		"git.",
	}

	for _, host := range gitHosts {
		if strings.Contains(url, host) {
			return true
		}
	}

	return false
}
