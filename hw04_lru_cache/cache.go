// Package hw04lrucache реализует LRU-кэш на основе двусвязного списка.
package hw04lrucache

import "sync"

// Key — ключ записи в кэше.
type Key string

// Cache — LRU-кэш ограниченной ёмкости.
type Cache interface {
	// Set сохраняет value по ключу key.
	// Если ключ уже есть: обновляет значение, переносит элемент в начало очереди и возвращает true.
	// Если ключа нет и ёмкость кэша <= 0: новый элемент не добавляется, возвращается false.
	// Если ключа нет и кэш заполнен: удаляется последний элемент очереди (LRU) вместе с данными
	// по его ключу, затем добавляется новый элемент в начало очереди; возвращается false.
	// Если ключа нет и есть место: элемент добавляется в начало очереди; возвращается false.
	Set(key Key, value interface{}) bool

	// Get возвращает значение по ключу и true, если ключ есть; иначе nil и false.
	// При успешном обращении элемент считается использованным и переносится в начало очереди.
	Get(key Key) (interface{}, bool)

	// Clear удаляет все элементы из кэша (очередь и словарь ключей).
	Clear()
}

// listVal — ключ и значение кэша в ListItem.Value узла очереди (listKey нужен при вытеснении LRU).
type listVal struct {
	listKey Key
	listVal interface{}
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// NewCache создаёт LRU-кэш вместимостью capacity элементов.
// При capacity <= 0 новые ключи методом Set добавлены не будут (существующие ключи при этом не появятся).
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

	if item, ok := c.items[key]; ok {
		item.Value.(*listVal).listVal = value
		c.queue.MoveToFront(item)
		return true
	}

	if c.capacity <= 0 {
		return false
	}

	if len(c.items) >= c.capacity {
		back := c.queue.Back()
		if back != nil {
			payload := back.Value.(*listVal)
			delete(c.items, payload.listKey)
			c.queue.Remove(back)
		}
	}

	item := c.queue.PushFront(&listVal{listKey: key, listVal: value})
	c.items[key] = item
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(item)
	return item.Value.(*listVal).listVal, true
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
