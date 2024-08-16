package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"time"
)

func New() (*bigcache.BigCache, error) {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		return nil, err
	}

	return cache, nil
}
