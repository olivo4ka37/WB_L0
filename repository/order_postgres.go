package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/olivo4ka37/WB_L0/internal/cache"
	"github.com/olivo4ka37/WB_L0/internal/models"
	"log"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

func (o *OrderPostgres) Create(ctx context.Context, orderUid string, order models.Order) error {

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		log.Println("cant marshal order")
		return err
	}
	_, err = o.db.Exec("INSERT INTO orders (order_uid, order_info) values ($1, $2)", orderUid, jsonOrder)
	if err != nil {
		log.Println("cant create order in db")
		return err
	}
	return nil
}

func (o *OrderPostgres) GetById(ctx context.Context, orderUid string) (models.Order, error) {
	var order models.Order
	row := o.db.QueryRow("SELECT order_info FROM orders WHERE order_uid = $1", orderUid)
	err := row.Scan(&order)
	//err := o.db.QueryRow("SELECT order_info FROM orders WHERE order_uid = $1", orderUid).Scan(&order)
	if err == sql.ErrNoRows {
		log.Println("no order with this id")
		return models.Order{}, err
	}
	if err != nil {
		log.Println("cant select order")
		return models.Order{}, err
	}

	return order, nil
}

func (o *OrderPostgres) RestoreCache(ctx context.Context, cache *cache.MemoryCache) error {
	rows, err := o.db.Query("SELECT order_info FROM orders")
	if err != nil {
		log.Println("cant restore cache from postgres")
		return err
	}
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order); err != nil {
			log.Println("cant scan values from db to go struct")
			return err
		}
		if err = cache.Set(&order); err != nil {
			log.Println("cant set data into cache")
			return err
		}
	}

	return nil
}
