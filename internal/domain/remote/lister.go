package remote

import (
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/models"
)

const KiteManifestFile = "kite.yaml"

func ScanForKits(store Storage) ([]models.KitSummary, error) {
	dirs, err := store.ListDirectories()
	if err != nil {
		return nil, err
	}

	var kits []models.KitSummary

	for _, dir := range dirs {
		if shouldSkipDirectory(dir) {
			continue
		}

		kitePath := filepath.Join(dir, KiteManifestFile)
		if !store.FileExists(kitePath) {
			continue
		}

		content, err := store.ReadFile(kitePath)
		if err != nil {
			continue
		}

		metadata, err := ParseMetadata(content)
		if err != nil {
			continue
		}

		kits = append(kits, metadata.ToKitSummary())
	}

	return kits, nil
}

func shouldSkipDirectory(path string) bool {
	base := filepath.Base(path)
	skipDirs := map[string]bool{
		".git":         true,
		".github":      true,
		"node_modules": true,
	}
	return skipDirs[base]
}
