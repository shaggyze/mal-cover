package cache

import (
	"context"
	"net/http"

	"github.com/rl404/fairy/cache"
	"github.com/rl404/fairy/errors/stack"
	"github.com/shaggyze/mal-cover/internal/domain/mal/entity"
	"github.com/shaggyze/mal-cover/internal/domain/mal/repository"
	"github.com/shaggyze/mal-cover/internal/errors"
	"github.com/shaggyze/mal-cover/internal/utils"
)

// Cache contains functions for mal cache.
type Cache struct {
	cacher cache.Cacher
	repo   repository.Repository
}

// New to create new mal cache.
func New(cacher cache.Cacher, repo repository.Repository) *Cache {
	return &Cache{
		cacher: cacher,
		repo:   repo,
	}
}

// GetList to get anime/manga list from cache.
func (c *Cache) GetList(ctx context.Context, username string, mainType string, status int, genre int) (data []entity.Entry, code int, err error) {
	key := utils.GetKey("list", username, mainType)
	if c.cacher.Get(ctx, key, &data) == nil {
		return data, http.StatusOK, nil
	}

	data, code, err = c.repo.GetList(ctx, username, mainType, status, genre)
	if err != nil {
		return nil, code, stack.Wrap(ctx, err)
	}

	if err := c.cacher.Set(ctx, key, data); err != nil {
		return nil, http.StatusInternalServerError, stack.Wrap(ctx, err, errors.ErrInternalCache)
	}

	return data, code, nil
}
