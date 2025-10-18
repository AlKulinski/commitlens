package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func newTrackCmd() *cli.Command {
	return &cli.Command{
		Name: "track",
		Commands: []*cli.Command{
			newTrackStartCmd(),
			newTrackStopCmd(),
		},
	}
}

func newTrackStartCmd() *cli.Command {
	return &cli.Command{
		Name: "start",
		Action: func(ctx context.Context, cmd *cli.Command) error {
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
