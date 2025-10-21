package cmd

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func newRootCmd(version string, db *sql.DB) *cli.Command {
	return &cli.Command{
		Commands: []*cli.Command{
			newVersionCmd(version),
			newDiffCmd(),
			newTrackCmd(db),
			newGroqCmd(),
		},
	}
}

func Execute(version string, db *sql.DB) error {
	cmd := newRootCmd(version, db)

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	return nil
}
