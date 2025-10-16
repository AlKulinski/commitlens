package cmd

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func newRootCmd(version string) *cli.Command {
	return &cli.Command{
		Commands: []*cli.Command{
			newVersionCmd(version),
			newDiffCmd(),
		},
	}
}

func Execute(version string) error {
	cmd := newRootCmd(version)

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	return nil
}
