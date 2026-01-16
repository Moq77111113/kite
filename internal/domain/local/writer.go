package local

import "github.com/moq77111113/kite/internal/domain/models"

type InstallOptions struct {
	Variables map[string]string
}

type FileWriter interface {
	Install(kit *models.Kit, destPath string) error
	InstallWithOptions(kit *models.Kit, destPath string, opts InstallOptions) error
}
