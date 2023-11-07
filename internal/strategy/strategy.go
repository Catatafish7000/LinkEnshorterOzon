package strategy

import (
	"LinkEnshorter/internal/repo/cache"
	"LinkEnshorter/internal/repo/database"
	"errors"
)

var errWrongFormat = errors.New("wrong format")

func RepoStrategy(args []string) (repo, error) {
	if len(args) != 2 {
		return nil, errWrongFormat
	}
	switch args[1] {
	case "-db":
		repo := database.NewRepo()
		return repo, nil

	case "-im":
		repo := cache.NewRepo()
		return repo, nil
	default:
		return nil, errWrongFormat
	}
}
