package main

import (
	"sync"
)

type MutexMap struct {
	l sync.RWMutex
	v sync.Map
}

func NewMutexMap(size int) Cache {
	return &MutexMap{}
}

func (c *MutexMap) Name() string {
	return "rwmutex_map"
}

func (c *MutexMap) Set(key string) {
	c.l.Lock()
	c.v.Store(key, key)
	c.l.Unlock()
}

func (c *MutexMap) Get(key string) bool {
	c.l.RLock()
	_, ok := c.v.Load(key)
	c.l.RUnlock()
	return ok
}

func (c *MutexMap) Close() {}
