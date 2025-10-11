package registry

import (
	"fmt"

	registryv1 "github.com/moq77111113/kite/api/registry/v1"
	"github.com/moq77111113/kite/internal/infra/git"
	"github.com/moq77111113/kite/internal/infra/http"
	"github.com/moq77111113/kite/pkg/console"
)

// NewClient creates the appropriate registry client based on URL
func NewClient(registryURL string) registryv1.Client {
	registryType := DetectRegistryType(registryURL)

	switch registryType {
	case RegistryTypeGit:
		return newGitClient(registryURL)

	case RegistryTypeHTTP:
		return http.NewHTTPClient(registryURL)

	case RegistryTypeLocal:
		return newLocalClient(registryURL)

	default:
		console.Warning("Unknown registry type, using mock")
		return http.NewMockClient()
	}
}

// newGitClient creates a Git-based client with error handling
func newGitClient(url string) registryv1.Client {
	gitClient := git.NewClient()
	client, err := git.NewGitClient(url, gitClient)
	if err != nil {
		console.Warning(fmt.Sprintf("Failed to initialize Git registry, using mock: %v", err))
		return http.NewMockClient()
	}
	return client
}

// newLocalClient creates a local file or Git client
func newLocalClient(url string) registryv1.Client {
	gitClient := git.NewClient()
	if gitClient.IsCloned(url) {
		// It's a local Git repo
		client, err := git.NewGitClient(url, gitClient)
		if err != nil {
			console.Warning(fmt.Sprintf("Failed to initialize local Git registry, using mock: %v", err))
			return http.NewMockClient()
		}
		return client
	}

	// TODO: Implement file-based LocalClient
	console.Warning("File-based local registry not yet implemented, using mock")
	return http.NewMockClient()
}
