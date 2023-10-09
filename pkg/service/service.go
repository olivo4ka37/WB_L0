package service

import (
	"context"
	"github.com/olivo4ka37/WB_L0/internal/cache"
	"github.com/olivo4ka37/WB_L0/internal/models"
	"github.com/olivo4ka37/WB_L0/pkg/repository"
)

type Orders interface {
	Create(ctx context.Context, orderUid string, order models.Order) error
	GetById(ctx context.Context, orderUid string) (models.Order, error)
}

type Service struct {
	Repo  Orders
	Cache *cache.MemoryCache
}

func NewService(repos *repository.Repository, cache *cache.MemoryCache) *Service {
	return &Service{
		Repo:  repos,
		Cache: cache,
	}
}
