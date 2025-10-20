package factories

import (
	"github.com/alkowskey/commitlens/internal/diff/domain"
	"github.com/alkowskey/commitlens/internal/diff/infra"
)

func CreateDiffer(algorithm domain.DiffAlgorithm) domain.Differ {
	if algorithm == domain.AlgorithmPatience {
		return infra.NewPatienceDiffer()
	}

	return infra.NewBaseDiffer()
}
