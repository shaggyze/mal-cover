package repository

import (
	"context"

	"github.com/shaggyze/mal-cover/internal/domain/mal/entity"
)

// Repository contains functions for mal domain.
type Repository interface {
	GetList(ctx context.Context, username, mainType string, status, genre) ([]entity.Entry, int, error)
}
