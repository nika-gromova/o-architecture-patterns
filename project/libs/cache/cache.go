package cache

import (
	"github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}

func New() *Cache {
	return &Cache{
		cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (c *Cache) Set(key string, value any) {
	c.cache.Set(key, value, cache.NoExpiration)
}

func (c *Cache) Get(key string) (any, bool) {
	return c.cache.Get(key)
}
