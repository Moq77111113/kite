package scan

import (
	"path/filepath"
	"strings"

	"github.com/moq77111113/kite/internal/domain/parse"
	"github.com/moq77111113/kite/internal/domain/port"
	"github.com/moq77111113/kite/internal/domain/types"
)

const KiteManifestFile = "kite.yaml"
const ReadmeFile = "README.md"

func ScanForKits(store port.Storage) ([]types.KitSummary, error) {
	dirs, err := store.ListDirectories()
	if err != nil {
		return nil, err
	}

	var kits []types.KitSummary

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

		metadata, err := parse.ParseMetadata(content)
		if err != nil {
			continue
		}

		kits = append(kits, metadata.ToKitSummary())
	}

	return kits, nil
}

func FindKit(store port.Storage, name string) (*types.KitDetailResponse, error) {
	kitDir := name
	kitePath := filepath.Join(kitDir, KiteManifestFile)

	content, err := store.ReadFile(kitePath)
	if err != nil {
		return nil, err
	}

	metadata, err := parse.ParseMetadata(content)
	if err != nil {
		return nil, err
	}

	files, err := readKitFiles(store, kitDir)
	if err != nil {
		return nil, err
	}

	readme := readReadme(store, kitDir)

	return metadata.ToKitDetail(files, readme), nil
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

func readKitFiles(store port.Storage, kitDir string) ([]types.KitFile, error) {
	paths, err := store.ListFiles(kitDir)
	if err != nil {

		return nil, err
	}

	var files []types.KitFile
	for _, p := range paths {
		if filepath.Base(p) == KiteManifestFile || strings.EqualFold(filepath.Base(p), ReadmeFile) {
			continue
		}
		b, err := store.ReadFile(p)
		if err != nil {

			continue
		}
		files = append(files, types.KitFile{Path: p, Content: string(b)})
	}

	return files, nil
}

func readReadme(store port.Storage, kitDir string) string {
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
