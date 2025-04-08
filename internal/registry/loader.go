package registry

import (
	"github.com/moq77111113/kite/internal/config"
)
type LoaderOptions struct {
	Path 			string
	Flavor 			string
}
type Loader interface {
    LoadIndex(config config.Config) (*Registry, error)
    LoadModules(config config.Config, names []string) ([]*Module, error)
}