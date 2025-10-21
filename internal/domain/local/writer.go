package local

import (
	"github.com/moq77111113/kite/internal/domain/models"
)

type FileWriter interface {
	Install(kit *models.Kit, destPath string) error
}
