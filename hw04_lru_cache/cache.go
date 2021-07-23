package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	Len() int
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, isItem := c.items[key]

	if isItem {
		item.Value = cacheItem{key, value}
		c.queue.MoveToFront(item)
		c.items[key] = item
	} else {
		newItem := c.queue.PushFront(cacheItem{key, value})
		c.items[key] = newItem
	}

	if c.queue.Len() > c.capacity {
		lastItem := c.queue.Back()
		delete(c.items, lastItem.Value.(cacheItem).key)
		c.queue.Remove(lastItem)
	}
	return isItem
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, isItem := c.items[key]

	var value interface{}

	if isItem {
		value = item.Value.(cacheItem).value
		c.queue.MoveToFront(item)
	}

	return value, isItem
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func (c *lruCache) Len() int {
	return c.queue.Len()
}
