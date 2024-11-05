package main

import (
	"github.com/coocood/freecache"
)

type FreeCache struct {
	v *freecache.Cache
}

func NewFreeCache(size int) Cache {
	cache := freecache.NewCache(size)

	return &FreeCache{
		v: cache,
	}
}

func (c *FreeCache) Name() string {
	return "freecache"
}

func (c *FreeCache) Set(key string) {
	k := []byte(key)
	c.v.Set(k, k, 0)
}

func (c *FreeCache) Get(key string) bool {
	_, err := c.v.Get([]byte(key))
	return err == nil
}

func (c *FreeCache) Close() {}
