package main

import (
	"sync"
)

type SyncMap struct {
	v sync.Map
}

func NewSyncMap(size int) Cache {
	return &SyncMap{}
}

func (c *SyncMap) Name() string {
	return "syncmap"
}

func (c *SyncMap) Set(key string) {
	c.v.Store(key, key)
}

func (c *SyncMap) Get(key string) bool {
	_, ok := c.v.Load(key)
	return ok
}

func (c *SyncMap) Close() {}
