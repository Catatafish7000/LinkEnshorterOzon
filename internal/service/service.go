package service

import (
	"context"
	"fmt"
	"log"
	"strings"
)

type Service struct {
	repo      repo
	generator generator
}

func NewService(Repo repo, generator generator) *Service {
	return &Service{
		Repo,
		generator,
	}
}

func (s *Service) ShowLink(ctx context.Context, hash string) (string, error) {
	url, err := s.repo.GetURL(ctx, hash)
	return url, err
}

func (s *Service) SaveShortURL(ctx context.Context, url string) (string, error) {
	var hash string
	for {
		hash = s.generator.GenerateHash()
		err := s.repo.SaveHashByURL(ctx, url, hash)
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			continue
		} else {
			log.Println(fmt.Sprintf("Failed to create hash. Error: %v", err))
			return "", err
		}
	}
	return hash, nil
}
