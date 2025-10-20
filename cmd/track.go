package cmd

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alkowskey/commitlens/internal/common/flags"
	"github.com/alkowskey/commitlens/internal/diff/domain"
	"github.com/alkowskey/commitlens/internal/diff/factories"
	diffServices "github.com/alkowskey/commitlens/internal/diff/services"
	"github.com/alkowskey/commitlens/internal/snapshot/repository"
	"github.com/alkowskey/commitlens/internal/snapshot/services"
	"github.com/alkowskey/commitlens/internal/snapshot/usecases"
	"github.com/urfave/cli/v3"
)

func newTrackCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:  "track",
		Usage: "track changes in a directory",
		Commands: []*cli.Command{
			newTrackStartCmd(db),
			newFlushCmd(db),
			newTrackCompareCmd(db),
		},
	}
}

func newTrackStartCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start tracking changes in a directory",
		Flags: []cli.Flag{
			flags.DirectoryFlag,
			flags.AlgorithmFlag,
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			subdirectory, snapshotService, err := prepareTrackCommand(db, cmd)
			if err != nil {
				return err
			}

			usecase := usecases.NewTrackStartUsecase(snapshotService)
			if err := usecase.Execute(subdirectory); err != nil {
				return err
			}
			return nil
		},
	}
}

func newTrackCompareCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:    "compare",
		Aliases: []string{"c"},
		Usage:   "Compare changes in a directory",
		Flags: []cli.Flag{
			flags.DirectoryFlag,
			flags.AlgorithmFlag,
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			subdirectory, snapshotService, err := prepareTrackCommand(db, cmd)
			if err != nil {
				return err
			}

			usecase := usecases.NewCompareUsecase(snapshotService)
			diff, err := usecase.Execute(subdirectory)
			if err != nil {
				return err
			}

			fmt.Println(diff)
			return nil
		},
	}
}

func newFlushCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:    "flush",
		Aliases: []string{"f"},
		Usage:   "flush all data snapshots",
		Flags: []cli.Flag{
			flags.AlgorithmFlag,
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			alg, err := domain.ParseDiffAlgorithm(cmd.String("algorithm"))
			if err != nil {
				return err
			}

			snapshotService := prepareSnapshotService(db, alg)
			usecase := usecases.NewFlushSnapshotsUsecase(snapshotService)
			if err := usecase.Execute(); err != nil {
				return err
			}
			return nil
		},
	}
}

func prepareTrackCommand(db *sql.DB, cmd *cli.Command) (string, services.SnapshotService, error) {
	subdirectory := cmd.String("directory")
	alg, err := domain.ParseDiffAlgorithm(cmd.String("algorithm"))
	if err != nil {
		return "", nil, err
	}

	snapshotService := prepareSnapshotService(db, alg)
	return subdirectory, snapshotService, nil
}

func prepareSnapshotService(db *sql.DB, algorithm domain.DiffAlgorithm) services.SnapshotService {
	snapshotRepository := repository.NewSnapshotRepository(db)
	differ := factories.CreateDiffer(algorithm)
	diffService := diffServices.NewDiffService(differ)
	return services.NewSnapshotService(snapshotRepository, diffService)
}
