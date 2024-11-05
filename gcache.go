package main

import (
	"github.com/bluele/gcache"
)

type Gcache struct {
	v gcache.Cache
}

func NewGcache(size int) Cache {
	cache := gcache.New(size).LRU().Build()

	return &Gcache{
		v: cache,
	}
}

func (c *Gcache) Name() string {
	return "gcache"
}

func (c *Gcache) Set(key string) {
	c.v.Set(key, key)
}

func (c *Gcache) Get(key string) bool {
	_, err := c.v.Get(key)
	return err == nil
}

func (c *Gcache) Close() {}
