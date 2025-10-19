package cmd

import (
	"context"
	"fmt"

	"github.com/alkowskey/commitlens/internal/diff/domain"
	"github.com/alkowskey/commitlens/internal/diff/infra"
	"github.com/alkowskey/commitlens/internal/diff/usecases"
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
			differ := infra.NewBaseDiffer()
			usecase := usecases.NewDiffUsecase(path_a, path_b, differ)
			diffRequest := domain.DiffRequest{
				Options: domain.DiffOptions{
					Ignore: []string{"*.md", "*.txt"},
				},
			}
			result, err := usecase.Execute(diffRequest)

			if err != nil {
				return err
			}

			fmt.Println(result.Added)
			fmt.Println(result.Removed)
			return nil
		},
	}

}
