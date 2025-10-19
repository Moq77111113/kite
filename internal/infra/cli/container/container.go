package container

import (
	"github.com/moq77111113/kite/internal/domain/local"
	"github.com/moq77111113/kite/internal/domain/remote"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
	"github.com/moq77111113/kite/internal/infra/storage/git"
)

type Key string

const ContainerKey Key = "container"

type Container struct {
	Config            *config.Config
	Repository        *remote.Repository
	Tracker           *local.Tracker
	Manager           *local.Manager
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

	installer := local.NewFsInstaller()
	locals := local.NewTracker(config.NewKitRegistry(cfg))
	kitLifecycle := local.NewManager(installer, locals)

	conflictChecker := local.NewConflictChecker()
	versionComparator := local.NewVersionComparator()

	return &Container{
		Config:            cfg,
		Repository:        repository,
		Tracker:           locals,
		Manager:           kitLifecycle,
		ConflictChecker:   conflictChecker,
		VersionComparator: versionComparator,
	}, nil
}
