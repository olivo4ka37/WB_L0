package service

import (
	"context"
	"github.com/olivo4ka37/WB_L0/internal/cache"
	"github.com/olivo4ka37/WB_L0/internal/models"
	"github.com/olivo4ka37/WB_L0/pkg/repository"
	"log"
)

type OrderService struct {
	repo  repository.Orders
	cache cache.MemoryCache
}

func NewOrderService(repo repository.Orders, cache cache.MemoryCache) *OrderService {
	return &OrderService{
		repo:  repo,
		cache: cache,
	}
}

func (s *OrderService) Create(ctx context.Context, orderUid string, order models.Order) error {
	s.cache.Set(&order)
	return s.repo.Create(ctx, orderUid, order)
}

func (s *OrderService) GetById(ctx context.Context, orderUid string) (models.Order, error) {
	data, err := s.cache.Get(orderUid)
	if err != nil {
		log.Println("cant get cached data (maybe there is no such key in map)", err)
	}

	if data != nil {
		return *data, nil
	}
	return s.repo.GetById(ctx, orderUid)
}
