package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newDiffCmd() *cli.Command {
	return &cli.Command{
		Name:  "diff",
		Usage: "Show changes between two commits",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "from",
				Aliases: []string{"f"},
				Usage:   "File to compare from",
			},
			&cli.StringFlag{
				Name:    "to",
				Aliases: []string{"t"},
				Usage:   "File to compare to",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path_a := cmd.String("from")
			path_b := cmd.String("to")
			fmt.Printf("diff of %s ... %s", path_a, path_b)
			return nil
		},
	}

}
