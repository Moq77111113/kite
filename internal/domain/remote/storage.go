package remote

type Storage interface {
	ListDirectories() ([]string, error)
	ListFiles(dir string) ([]string, error)
	FileExists(path string) bool
	ReadFile(path string) ([]byte, error)
	Sync() error
}
