package remote

import (
	"path/filepath"
	"strings"

	"github.com/moq77111113/kite/internal/domain/models"
	"github.com/moq77111113/kite/internal/domain/template"
)

const ReadmeFile = "README.md"

func FindKit(store Storage, name string) (*models.Kit, error) {
	kitDir := name
	kitePath := filepath.Join(kitDir, KiteManifestFile)

	content, err := store.ReadFile(kitePath)
	if err != nil {
		return nil, err
	}

	metadata, err := ParseMetadata(content)
	if err != nil {
		return nil, err
	}

	files, err := readKitFiles(store, kitDir)
	if err != nil {
		return nil, err
	}

	readme := readReadme(store, kitDir)
	kit := metadata.ToKitDetail(name, files, readme)

	engine := template.NewEngine()
	detected := engine.ExtractFromFiles(files)
	kit.Variables = template.MergeWithMetadata(detected, metadata.Variables)

	return kit, nil
}

func readKitFiles(store Storage, kitDir string) ([]models.File, error) {
	paths, err := store.ListFiles(kitDir)
	if err != nil {

		return nil, err
	}

	var files []models.File
	for _, p := range paths {
		if filepath.Base(p) == KiteManifestFile || strings.EqualFold(filepath.Base(p), ReadmeFile) {
			continue
		}
		b, err := store.ReadFile(p)
		if err != nil {
			continue
		}
		relPath, _ := filepath.Rel(kitDir, p)
		files = append(files, models.File{Path: relPath, Content: string(b)})
	}

	return files, nil
}

func readReadme(store Storage, kitDir string) string {
	readmePath := filepath.Join(kitDir, ReadmeFile)
	if !store.FileExists(readmePath) {
		// Try other case variations
		files, err := store.ListFiles(kitDir)
		if err != nil {
			return ""
		}
		found := false
		for _, f := range files {
			if strings.EqualFold(filepath.Base(f), ReadmeFile) {
				readmePath = f
				found = true
				break
			}
		}
		if !found {
			return ""
		}
	}

	content, err := store.ReadFile(readmePath)
	if err != nil {
		return ""
	}
	return string(content)
}
