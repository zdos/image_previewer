package lruCache

import "sync"

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	mx       sync.RWMutex
	queue    List
	items    map[string]*ListItem
}

func (c *lruCache) Set(k string, value interface{}) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	v, ok := c.items[k]
	if ok {
		v.Value = value
		c.queue.MoveToFront(v)
		c.items[k] = c.queue.Front()
	} else {
		if c.capacity == c.queue.Len() {
			for key, value := range c.items {
				if value == c.queue.Back() {
					delete(c.items, key)
				}
			}
			c.queue.Remove(c.queue.Back())
		}
		c.queue.PushFront(value)
		c.items[k] = c.queue.Front()
	}
	return ok
}

func (c *lruCache) Get(k string) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	if v, ok := c.items[k]; ok {
		c.queue.MoveToFront(v)
		return v.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	for k := range c.items {
		delete(c.items, k)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[string]*ListItem, capacity),
	}
}
