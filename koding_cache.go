package main

import (
	"github.com/koding/cache"
)

type KodingCache struct {
	v cache.Cache
}

func NewKodingCache(size int) Cache {
	cache := cache.NewLRU(size)
	return &KodingCache{
		v: cache,
	}
}

func (c *KodingCache) Name() string {
	return "koding_cache"
}

func (c *KodingCache) Set(key string) {
	c.v.Set(key, key)
}

func (c *KodingCache) Get(key string) bool {
	_, err := c.v.Get(key)
	return err == nil
}

func (c *KodingCache) Close() {}
