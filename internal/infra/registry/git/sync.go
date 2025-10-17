package git

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	msgCloningRegistry    = "Cloning registry for the first time..."
	msgRegistryCloned     = "Registry cloned successfully"
	msgTemplateNotInCache = "Template not found in cache, syncing registry..."
	msgRegistrySynced     = "Registry synced, retrying..."
)

type SyncCallback func(message string)

func (c *GitClient) ensureRepositoryClone() error {
	if c.git.IsCloned(c.cachePath) {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(c.cachePath), 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	return c.git.Clone(c.repoURL, c.cachePath)
}

func (c *GitClient) syncRepository() error {
	if !c.git.IsCloned(c.cachePath) {
		return c.ensureRepositoryClone()
	}
	return c.git.Pull(c.cachePath)
}

func (c *GitClient) syncIfNeeded(templatePath string) error {
	if isLocalPath(c.repoURL) || !c.git.IsCloned(c.cachePath) {
		return nil
	}

	if c.syncCallback != nil {
		c.syncCallback(msgTemplateNotInCache)
	}

	if err := c.syncRepository(); err != nil {
		return err
	}

	if c.syncCallback != nil {
		c.syncCallback(msgRegistrySynced)
	}

	return nil
}
