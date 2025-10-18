package repo

import (
	"fmt"
	"os"
)

type ConflictChecker struct{}

func NewConflictChecker() *ConflictChecker {
	return &ConflictChecker{}
}

type ConflictResult struct {
	HasConflict  bool
	Reason       string
	ExistingPath string
}

func (d *ConflictChecker) Check(targetPath string) (*ConflictResult, error) {
	if targetPath == "" {
		return nil, fmt.Errorf("target path cannot be empty")
	}

	info, err := os.Stat(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &ConflictResult{
				HasConflict:  false,
				Reason:       "",
				ExistingPath: "",
			}, nil
		}
		return nil, fmt.Errorf("failed to check path: %w", err)
	}

	if info.IsDir() {
		return &ConflictResult{
			HasConflict:  true,
			Reason:       "directory already exists at target path",
			ExistingPath: targetPath,
		}, nil
	}

	return &ConflictResult{
		HasConflict:  true,
		Reason:       "file exists at target path (expected directory)",
		ExistingPath: targetPath,
	}, nil
}
