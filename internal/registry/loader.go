package registry

type Loader interface {
	Load(path string) (*Registry, error)
}