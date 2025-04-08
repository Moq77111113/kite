package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/moq77111113/kite/internal/config"
	"gopkg.in/yaml.v3"
)

type HttpLoader struct{}


func (h *HttpLoader) LoadIndex(config config.Config) (*Registry, error) {

	u := fmt.Sprintf("%s/flavors/%s", config.Registry, config.Flavor) 
	resp, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}


	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r Registry

	switch fileType := http.DetectContentType(body); fileType {
	case "application/json":
		err = json.Unmarshal(body, &r)
	case "application/x-yaml", "text/yaml":
		err = yaml.Unmarshal(body, &r)
	default:
		return nil, fmt.Errorf("unsupported content type: %s", fileType)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid registry format: %w", err)
	}
	return &r, nil	
}

func (h *HttpLoader) LoadModules(config config.Config, names []string) ([]*Module, error) {
	
	modules := make([]*Module, 0, len(names))
	flavorUrl := fmt.Sprintf("%s/flavors/%s", config.Registry, config.Flavor)
	
	for _, name := range names {
		u := fmt.Sprintf("%s/modules/%s.json", flavorUrl, name)
		resp, err := http.Get(u)
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var m Module

		switch fileType := http.DetectContentType(body); fileType {
		case "application/json":
			err = json.Unmarshal(body, &m)
		case "application/x-yaml", "text/yaml":
			err = yaml.Unmarshal(body, &m)
		default:
			return nil, fmt.Errorf("unsupported content type: %s", fileType)
		}

		if err != nil {
			return nil, fmt.Errorf("invalid module format: %w", err)
		}

		modules = append(modules, &m)
	}

	return modules, nil

}