package filesystem

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/local"
	"github.com/moq77111113/kite/internal/domain/models"
	"github.com/moq77111113/kite/internal/domain/template"
)

type Writer struct {
	engine *template.Engine
}

func NewWriter() *Writer {
	return &Writer{engine: template.NewEngine()}
}

func (w *Writer) Install(kit *models.Kit, destPath string) error {
	return w.InstallWithOptions(kit, destPath, local.InstallOptions{})
}

func (w *Writer) InstallWithOptions(kit *models.Kit, destPath string, opts local.InstallOptions) error {
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	for _, file := range kit.Files {
		if err := w.writeFile(destPath, file, kit.Variables, opts.Variables); err != nil {
			return err
		}
	}
	return nil
}

func (w *Writer) writeFile(basePath string, file models.File, defs []models.Variable, vals map[string]string) error {
	filePath := filepath.Join(basePath, file.Path)
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	content := file.Content
	if !template.IsBinaryContent(content) && len(vals) > 0 {
		interpolated, err := w.engine.Interpolate(content, vals, defs)
		if err != nil {
			return fmt.Errorf("template interpolation failed for %s: %w", file.Path, err)
		}
		content = interpolated
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}
