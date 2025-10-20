package flags

import (
	"github.com/urfave/cli/v3"
)

var (
	DirectoryFlag = &cli.StringFlag{
		Name:     "directory",
		Aliases:  []string{"d"},
		Usage:    "Directory to track",
		Required: true,
	}

	AlgorithmFlag = &cli.StringFlag{
		Name:    "algorithm",
		Aliases: []string{"a"},
		Usage:   "Diff algorithm (base, patience)",
		Value:   "patience",
	}
)
