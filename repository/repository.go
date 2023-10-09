package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/olivo4ka37/WB_L0/internal/models"
)

type Orders interface {
	Create(ctx context.Context, orderUid string, order models.Order) error
	GetById(ctx context.Context, orderUid string) (models.Order, error)
}

type Repository struct {
	Orders
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Orders: NewOrderPostgres(db),
	}
}
