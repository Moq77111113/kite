package registry

import (
	"fmt"

	registryv1 "github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/infra/registry/git"
	registryhttp "github.com/moq77111113/kite/internal/infra/registry/http"
	registrymock "github.com/moq77111113/kite/internal/infra/registry/mock"
	"github.com/moq77111113/kite/pkg/console"
)

// NewClient creates the appropriate registry client based on URL with optional sync callback
func NewClient(registryURL string, syncCallback git.SyncCallback) registryv1.Client {
	registryType := DetectRegistryType(registryURL)

	switch registryType {
	case RegistryTypeGit:
		return newGitClient(registryURL, syncCallback)

	case RegistryTypeHTTP:
		return registryhttp.NewHTTPClient(registryURL)

	case RegistryTypeLocal:
		return newLocalClient(registryURL, syncCallback)

	default:
		console.Warning("Unknown registry type, using mock")
		return registrymock.NewMockClient()
	}
}

func newGitClient(url string, syncCallback git.SyncCallback) registryv1.Client {
	gitClient := git.NewClient()
	client, err := git.NewGitClient(url, gitClient, syncCallback)
	if err != nil {
		console.Warning(fmt.Sprintf("Failed to initialize Git registry, using mock: %v", err))
		return registrymock.NewMockClient()
	}
	return client
}

func newLocalClient(url string, syncCallback git.SyncCallback) registryv1.Client {
	gitClient := git.NewClient()
	if gitClient.IsCloned(url) {
		client, err := git.NewGitClient(url, gitClient, syncCallback)
		if err != nil {
			console.Warning(fmt.Sprintf("Failed to initialize local Git registry, using mock: %v", err))
			return registrymock.NewMockClient()
		}
		return client
	}

	console.Warning("File-based local registry not yet implemented, using mock")
	return registrymock.NewMockClient()
}
