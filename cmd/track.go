package cmd

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alkowskey/commitlens/internal/diff/infra"
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
			&cli.StringFlag{
				Name:     "directory",
				Aliases:  []string{"d"},
				Usage:    "Directory to track",
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			subdirectory := cmd.String("directory")

			snapshotRepository := repository.NewSnapshotRepository(db)
			differ := infra.NewBaseDiffer()
			diffService := diffServices.NewDiffService(differ)
			snapshotService := services.NewSnapshotService(snapshotRepository, diffService)
			usecase := usecases.NewTrackStartUsecase(snapshotService)

			err := usecase.Execute(subdirectory)
			if err != nil {
				return err
			}
			fmt.Println("tacking started")
			return nil
		},
	}
}

func newFlushCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:    "flush",
		Aliases: []string{"f"},
		Usage:   "flush all data snapshots",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			snapshotRepository := repository.NewSnapshotRepository(db)
			usecase := usecases.NewFlushSnapshotsUsecase(snapshotRepository)
			err := usecase.Execute()
			if err != nil {
				return err
			}
			fmt.Printf("Snapshots flushed")
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
			&cli.StringFlag{
				Name:     "directory",
				Aliases:  []string{"d"},
				Usage:    "Directory to track",
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			subdirectory := cmd.String("directory")

			snapshotRepository := repository.NewSnapshotRepository(db)
			differ := infra.NewBaseDiffer()
			diffService := diffServices.NewDiffService(differ)
			snapshotService := services.NewSnapshotService(snapshotRepository, diffService)
			usecase := usecases.NewCompareUsecase(snapshotService)

			diff, err := usecase.Execute(subdirectory)
			if err != nil {
				return err
			}
			fmt.Println(diff)
			fmt.Printf("Snapshots compared")
			return nil
		},
	}
}
