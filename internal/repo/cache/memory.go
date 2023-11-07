package cache

import (
	"context"
	"errors"
	_ "github.com/lib/pq"
	"sync"
	"time"
)

const ttl = time.Hour * 24

type link struct {
	url       string
	createdAt time.Time
}
type repo struct {
	data map[string]link
	mx   sync.Mutex
}

func NewRepo() *repo {
	data := make(map[string]link)
	return &repo{
		data: data,
		mx:   sync.Mutex{},
	}
}

func (r *repo) SaveHashByURL(ctx context.Context, url, hash string) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	_, ok := r.data[hash]
	if ok {
		err := errors.New("Error: duplicate key value violates unique constraint")
		return err
	}
	r.data[hash] = link{
		url:       url,
		createdAt: time.Now(),
	}
	return nil
}

func (r *repo) GetURL(ctx context.Context, hash string) (string, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	ans, ok := r.data[hash]
	if !ok {
		err := errors.New("no such hash in cache")
		return ans.url, err
	}
	return ans.url, nil
}

func (r *repo) Clear(ctx context.Context) {
	r.mx.Lock()
	defer r.mx.Unlock()
	current := time.Now()
	for i := range r.data {
		if current.Sub(r.data[i].createdAt) >= ttl {
			delete(r.data, i)
		}
	}
}
