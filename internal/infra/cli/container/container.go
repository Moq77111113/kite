package container

import (
	"github.com/moq77111113/kite/internal/domain/install"
	"github.com/moq77111113/kite/internal/domain/port"
	"github.com/moq77111113/kite/internal/domain/repo"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
	"github.com/moq77111113/kite/internal/infra/storage/git"
)

type Key string

const ContainerKey Key = "container"

type Container struct {
	Storage              port.Storage
	Repository           *repo.Repository
	Config               *config.Config
	InstallationRegistry *install.LocalKits
	FsInstaller          install.FsInstaller
	KitLifecycle         *install.KitLifecycle
	ConflictChecker      *repo.ConflictChecker
	VersionComparator    *repo.VersionComparator
}

func NewContainer(cfgPath string) (*Container, error) {
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return nil, err
	}

	gitClient := git.NewClient()
	storage, err := git.NewStorage(cfg.Registry, gitClient)
	if err != nil {
		return nil, err
	}

	repository := repo.NewRepository(storage)

	installer := install.NewFsInstaller()
	locals := install.NewLocalKits(config.NewKitRegistry(cfg))
	kitLifecycle := install.NewKitLifecycle(installer, locals)

	conflictChecker := repo.NewConflictChecker()
	versionComparator := repo.NewVersionComparator()

	return &Container{
		FsInstaller:          installer,
		Repository:           repository,
		Storage:              storage,
		Config:               cfg,
		InstallationRegistry: locals,
		ConflictChecker:      conflictChecker,
		KitLifecycle:         kitLifecycle,
		VersionComparator:    versionComparator,
	}, nil
}
