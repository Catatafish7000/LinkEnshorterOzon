package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

const Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
const AlphLen = 63

type Service struct {
	Repo      repo
	generator generator
}

func NewService(Repo repo, generator generator) *Service {
	return &Service{
		Repo,
		generator,
	}
}

func (s *Service) ShowLink(ctx context.Context, hash string) (string, error) {
	url, err := s.Repo.GetURL(ctx, hash)
	return url, err
}

func (s *Service) SaveShortURL(ctx context.Context, url string) (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	var hash string
	var errGen error
	for {
		hash = s.generator.GenerateHash()
		if errGen != nil {
			return "", errGen
		}
		err := s.Repo.SaveHashByURL(ctx, url, hash)
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
