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
			newTrackStopCmd(),
		},
	}
}

func newTrackStartCmd(db *sql.DB) *cli.Command {
	return &cli.Command{
		Name: "start",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			snapshotRepository := repository.NewSnapshotRepository(db)
			usecase := usecases.NewTrackStartUsecase(snapshotRepository)
			err := usecase.Execute(
				"test_files/b.txt",
				"test_files/a.txt",
			)
			if err != nil {
				return err
			}
			fmt.Printf("tacking started")
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
