package cleaner

import (
	"context"
	"github.com/robfig/cron"
)

type Сleaner struct {
	repo repo
}

func NewCleaner(repo repo) *Сleaner {
	return &Сleaner{repo: repo}
}

func (c *Сleaner) Clean(ctx context.Context) {
	cron := cron.New()
	cron.AddFunc("@daily", func() {
		c.repo.Clear(ctx)
	})
	cron.Start()
}
