package main

import (
	"container/list"
	"errors"
)

type LRU struct {
	size      int
	evictions *list.List
	items     map[string]string
}

type entry struct {
	key string
}

func NewLru(size int) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("Should provide a positive size")
	}
	c := &LRU{
		size:      size,
		evictions: list.New(),
		items:     make(map[string]string),
	}
	return c, nil
}

func (l *LRU) Length() int {
	return len(l.items)
}

func (c *LRU) Add(key, val string) bool {
	if _, ok := c.items[key]; ok {
		return false
	}
	if c.evictions.Len() < c.size {
		ent := &entry{key}
		c.evictions.PushFront(ent)
		c.items[key] = val
		return true
	} else {
		oldest := c.evictions.Back()
		// remove it from list and map
		c.evictions.Remove(oldest)
		delete(c.items, oldest.Value.(*entry).key)

		ent := &entry{key}
		c.evictions.PushFront(ent)
		c.items[key] = val
		return true
	}
}

func (c *LRU) Get(key string) (result interface{}) {
	if ent, ok := c.items[key]; ok {
		return ent
	}
	return
}
