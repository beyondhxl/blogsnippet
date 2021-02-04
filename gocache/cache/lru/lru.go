package lru

import (
	"container/list"
)

// Cache is a LRU cache. It is not safe for concurrent access.
type CaChe struct {
	maxBytes int64 // 允许使用的最大内存
	nbytes   int64 // 当前已使用的内存
	ll       *list.List
	cache    map[string]*list.Element

	// optional and executed when an entry is purged
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *CaChe {
	return &CaChe{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *CaChe) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele) // 双向链表作为队列，队首队尾是相对的，在这里约定 front 为队尾
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *CaChe) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		if c.nbytes >= int64(len(kv.key))+int64(kv.value.Len()) {
			c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		}
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *CaChe) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key: key, value: value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *CaChe) Len() int {
	return c.ll.Len()
}
