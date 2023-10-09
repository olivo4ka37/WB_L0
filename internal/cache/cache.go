package cache

import (
	"errors"
	"github.com/olivo4ka37/WB_L0/internal/models"
	"sync"
)

type MemoryCache struct {
	sync.RWMutex
	Cache map[string]*models.Order
}

func NewCache() *MemoryCache {
	c := &MemoryCache{}
	c.Cache = make(map[string]*models.Order)
	return c
}

func (c *MemoryCache) Set(order *models.Order) error {
	c.Lock()
	defer c.Unlock()

	c.Cache[order.OrderUID] = order

	return nil
}

func (c *MemoryCache) Get(key string) (*models.Order, error) {
	c.RLock()
	data, ex := c.Cache[key]
	defer c.RUnlock()

	if !ex {
		return nil, errors.New("no such element in mememory cache")
	}

	return data, nil
}
