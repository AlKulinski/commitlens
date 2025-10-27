package domain

type DiffRequest struct {
	Options DiffOptions
}

type DiffOptions struct {
	Ignore []string
}

type DiffResult struct {
	HasDifferences bool
	Added          []string
	Removed        []string
	Path           string
}

func (d DiffResult) String() string {
	if !d.HasDifferences {
		return "No differences found"
	}

	result := ""

	result += "\033[1m [FILE] " + d.Path + "\033[0m\n"

	for _, line := range d.Removed {
		result += "\033[31m- " + line + "\033[0m\n"
	}

	for _, line := range d.Added {
		result += "\033[32m+ " + line + "\033[0m\n"
	}

	return result
}

func (d DiffResult) StringPlain() string {
	if !d.HasDifferences {
		return "No differences found"
	}

	result := ""

	for _, line := range d.Removed {
		result += "- " + line + "\n"
	}

	for _, line := range d.Added {
		result += "+ " + line + "\n"
	}

	return result
}

type Differ interface {
	Compare(sourcePath string, targetPath string) (DiffResult, error)
}
