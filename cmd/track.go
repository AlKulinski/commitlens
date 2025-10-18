package cmd

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alkowskey/commit-suggester/internal/snapshot/repository"
	"github.com/alkowskey/commit-suggester/internal/track/usecases"
	"github.com/urfave/cli/v3"
)

func newTrackCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name: "track",
		Commands: []*cli.Command{
			newTrackStartCmd(db),
			newFlushCmd(db),
			newTrackStopCmd(),
		},
	}
}

func newTrackStartCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name: "start",
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
			usecase := usecases.NewTrackStartUsecase(snapshotRepository)
			err := usecase.Execute(subdirectory)
			if err != nil {
				return err
			}
			fmt.Printf("tacking started")
			return nil
		},
	}
}

func newFlushCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name:    "flush",
		Aliases: []string{"f"},
		Usage:   "flush all data",
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

func newTrackStopCmd() *cli.Command {
	return &cli.Command{
		Name: "stop",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("tacking stopped")
			return nil
		},
	}
}
