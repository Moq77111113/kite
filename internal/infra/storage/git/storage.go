package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/remote"
	"github.com/moq77111113/kite/pkg/console"
)

type Storage struct {
	cachePath string
	repoURL   string
	git       Client
}

func NewStorage(repoURL string, gitClient Client) (remote.Storage, error) {
	var cachePath string
	var isLocalRepo bool

	if IsLocalPath(repoURL) && gitClient.IsCloned(repoURL) {
		cachePath = repoURL
		isLocalRepo = true
	} else {
		cachePath = GetCachePath(repoURL)
		isLocalRepo = false
	}

	store := &Storage{
		cachePath: cachePath,
		repoURL:   repoURL,
		git:       gitClient,
	}

	if !isLocalRepo && !gitClient.IsCloned(cachePath) {
		console.Info("Cloning registry for the first time...")
		if err := store.ensureClone(); err != nil {
			return nil, fmt.Errorf("failed to clone: %w", err)
		}
		console.Info("Registry cloned successfully")
	}

	return store, nil
}

func (s *Storage) ListDirectories() ([]string, error) {
	var dirs []string

	err := filepath.Walk(s.cachePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != s.cachePath {
			relPath, _ := filepath.Rel(s.cachePath, path)
			dirs = append(dirs, relPath)
		}

		return nil
	})

	return dirs, err
}

func (s *Storage) ListFiles(dir string) ([]string, error) {
	var files []string

	root := filepath.Join(s.cachePath, dir)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(s.cachePath, path)
		if err != nil {
			return nil
		}
		files = append(files, rel)
		return nil
	})

	return files, err
}

func (s *Storage) ReadFile(path string) ([]byte, error) {
	fullPath := filepath.Join(s.cachePath, path)
	return os.ReadFile(fullPath)
}

func (s *Storage) FileExists(path string) bool {
	fullPath := filepath.Join(s.cachePath, path)
	_, err := os.Stat(fullPath)
	return err == nil
}

func (s *Storage) Sync() error {
	if IsLocalPath(s.repoURL) {
		return nil
	}

	console.Info("Syncing registry...")
	if err := s.syncRepo(); err != nil {
		return err
	}
	console.Info("Registry synced successfully")
	return nil
}

func (s *Storage) ensureClone() error {
	if s.git.IsCloned(s.cachePath) {
		return nil
	}
	return s.git.Clone(s.repoURL, s.cachePath)
}

func (s *Storage) syncRepo() error {
	if !s.git.IsCloned(s.cachePath) {
		return s.ensureClone()
	}
	return s.git.Pull(s.cachePath)
}
