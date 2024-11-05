package main

import (
	"github.com/allegro/bigcache"
)

type BigCache struct {
	v *bigcache.BigCache
}

func NewBigCache(size int) Cache {
	config := bigcache.DefaultConfig(0)
	config.MaxEntrySize = 100
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(err)
	}

	return &BigCache{
		v: cache,
	}
}

func (c *BigCache) Name() string {
	return "bigcache"
}

func (c *BigCache) Set(key string) {
	c.v.Set(key, []byte(key))
}

func (c *BigCache) Get(key string) bool {
	_, err := c.v.Get(key)
	return err == nil
}

func (c *BigCache) Close() {}
