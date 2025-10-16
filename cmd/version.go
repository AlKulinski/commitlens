package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newVersionCmd(version string) *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show version",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("version: %s", version)
			return nil
		},
	}
}
