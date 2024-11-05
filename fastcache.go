package main

import (
	"github.com/VictoriaMetrics/fastcache"
)

type FastCache struct {
	v *fastcache.Cache
}

func NewFastCache(size int) Cache {
	cache := fastcache.New(size)

	return &FastCache{
		v: cache,
	}
}

func (c *FastCache) Name() string {
	return "fastcache"
}

func (c *FastCache) Set(key string) {
	k2 := key // copy
	k := []byte(k2)
	c.v.Set(k, k)
}

func (c *FastCache) Get(key string) bool {
	k := []byte(key)
	_, ok := c.v.HasGet(k, k)
	return ok
}

func (c *FastCache) Close() {}
