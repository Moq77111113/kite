package remote

import "time"

type Storage interface {
	ListDirectories() ([]string, error)
	ListFiles(dir string) ([]string, error)
	FileExists(path string) bool
	ReadFile(path string) ([]byte, error)
	LastUpdate(path string) (*time.Time, error)
	Sync() error
}
