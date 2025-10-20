package domain

type DiffAlgorithm string

const (
	AlgorithmBase     DiffAlgorithm = "base"
	AlgorithmPatience DiffAlgorithm = "patience"
)

func (d DiffAlgorithm) IsValid() bool {
	switch d {
	case AlgorithmBase, AlgorithmPatience:
		return true
	}
	return false
}

func ParseDiffAlgorithm(s string) (DiffAlgorithm, error) {
	switch s {
	case "base":
		return AlgorithmBase, nil
	case "patience":
		return AlgorithmPatience, nil
	default:
		return AlgorithmPatience, nil
	}
}
