package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/yaml.v3"
)

type HttpLoader struct{}

func (h *HttpLoader) Load(url string) (*Registry, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
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