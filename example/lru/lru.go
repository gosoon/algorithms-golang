package lru

import (
	"fmt"
	"sync"
)

type Interface interface {
	Get(key string) (*info, bool)
	Put(key string, val string) (bool, error)
}

// LRUCache xxx
type LRUCache struct {
	data map[string]*listItem
	head *listItem
	size int64
	mu   sync.RWMutex
	ct   int64
}

type listItem struct {
	info *info
	prev *listItem
	next *listItem
}

type info struct {
	key   string
	value string
}

func newListItem(key string, val string) *listItem {
	return &listItem{info: &info{key: key, value: val}}
}

func NewLRUCache(size int64) *LRUCache {
	return &LRUCache{
		data: make(map[string]*listItem),
		size: size,
	}
}

// Get xxx
func (lru *LRUCache) Get(key string) (*info, bool) {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	if item, found := lru.data[key]; found {
		return item.info, true
	} else {
		return nil, false
	}
}

// Put xxx
func (lru *LRUCache) Put(key string, val string) (bool, error) {
	lru.mu.Lock()
	lru.mu.Unlock()

	_, ok := lru.data[key]
	if !ok {
		// check cache is full,if full and delete the least recently used data
		if ok := lru.isFull(); ok {
			lru.removeLast()
		}

		listItem := newListItem(key, val)
		lru.data[key] = listItem
		lru.ct++
	} else {
		lru.data[key].info.value = val
	}
	if err := lru.moveItemToHead(key); err != nil {
		return false, err
	}
	return true, nil
}

func (lru *LRUCache) removeLast() (bool, error) {
	h := lru.head

	if h.next != nil {
		h = h.next
	}

	key := h.info.key
	_, ok := lru.data[key]
	if !ok {
		return false, fmt.Errorf("%v not found", key)
	}
	if err := lru.removeItemFromList(key); err != nil {
		return false, err
	}
	lru.ct--
	return true, nil
}

func (lru *LRUCache) isFull() bool {
	return lru.ct == lru.size
}

func (lru *LRUCache) moveItemToHead(key string) error {
	cur, ok := lru.data[key]
	if !ok {
		return fmt.Errorf("%v not found", key)
	}

	if cur == lru.head {
		return nil
	}

	if cur.prev != nil {
		cur.prev.next = cur.next
	}

	if cur.next != nil {
		cur.next.prev = cur.prev
	}

	if lru.head != nil {
		lru.head.prev = cur
	}

	cur.next = lru.head
	cur.prev = nil
	lru.head = cur
	return nil
}

func (lru *LRUCache) removeItemFromList(key string) error {
	cur, ok := lru.data[key]
	if !ok {
		return fmt.Errorf("%v not found", key)
	}

	if cur.prev != nil {
		cur.prev.next = cur.next
	}
	if cur.next != nil {
		cur.next.prev = cur.prev
	}

	if cur == lru.head {
		lru.head.next = cur.next
	}

	delete(lru.data, key)
	return nil
}

// List xxx
func (lru *LRUCache) List() {
	h := lru.head
	for h != nil {
		fmt.Println(h.info)
		h = h.next
	}
}
