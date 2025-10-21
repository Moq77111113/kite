package git

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const (
	cacheMetaFileName = ".kite-meta"
	DefaultCacheTTL   = 15 * time.Minute
)

type CacheMeta struct {
	LastSync time.Time `json:"last_sync"`
}

func loadCacheMeta(cachePath string) (*CacheMeta, error) {
	metaPath := filepath.Join(cachePath, cacheMetaFileName)

	data, err := os.ReadFile(metaPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &CacheMeta{LastSync: time.Time{}}, nil
		}
		return nil, err
	}

	var meta CacheMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

func saveCacheMeta(cachePath string, meta *CacheMeta) error {
	metaPath := filepath.Join(cachePath, cacheMetaFileName)

	data, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	return os.WriteFile(metaPath, data, 0644)
}

func shouldSync(cachePath string, ttl time.Duration) bool {
	meta, err := loadCacheMeta(cachePath)
	if err != nil {
		return true
	}

	if meta.LastSync.IsZero() {
		return true
	}

	elapsed := time.Since(meta.LastSync)
	return elapsed >= ttl
}

func updateLastSync(cachePath string) error {
	meta := &CacheMeta{
		LastSync: time.Now(),
	}
	return saveCacheMeta(cachePath, meta)
}
