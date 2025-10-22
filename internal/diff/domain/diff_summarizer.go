package domain

import "context"

type DiffSummary struct {
	Total   int
	Added   int
	Removed int
	Changed int
	Message string
}

type DiffSumarizer interface {
	Summarize(context.Context, []DiffResult) (string, error)
}
