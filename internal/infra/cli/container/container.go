package container

import (
	"github.com/moq77111113/kite/internal/domain/local"
	"github.com/moq77111113/kite/internal/domain/remote"
	"github.com/moq77111113/kite/internal/infra/filesystem"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
	"github.com/moq77111113/kite/internal/infra/storage/git"
)

type Key string

const ContainerKey Key = "container"

type Container struct {
	Config            *config.Config
	Repository        *remote.Repository
	Tracker           *local.Tracker
	Installer         *local.Installer
	ConflictChecker   *local.ConflictChecker
	VersionComparator *local.VersionComparator
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

	repository := remote.NewRepository(storage)

	writer := filesystem.NewWriter()
	tracker := local.NewTracker(config.NewKitRegistry(cfg))
	installer := local.NewInstaller(writer, tracker)

	conflictChecker := local.NewConflictChecker()
	versionComparator := local.NewVersionComparator()

	return &Container{
		Config:            cfg,
		Repository:        repository,
		Tracker:           tracker,
		Installer:         installer,
		ConflictChecker:   conflictChecker,
		VersionComparator: versionComparator,
	}, nil
}
