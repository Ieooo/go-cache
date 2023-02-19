package lru

import (
	"cache/server/value"
	"container/list"
	"sync"
)

type kv struct {
	key   string
	value value.Value
}

type Cache struct {
	sync.Mutex

	maxBytes  int64
	usedBytes int64
	m         map[string]*list.Element
	list      *list.List
}

func NewCache(maxBytes int64) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		m:        make(map[string]*list.Element),
		list:     list.New(),
	}
}

// get value according to key
// return nil if key is not exists
func (c *Cache) Get(key string) value.Value {
	c.Lock()
	defer c.Unlock()

	if v, ok := c.m[key]; ok {
		c.list.MoveToFront(v)
		return v.Value.(*kv).value
	}
	return nil
}

// set value by key
func (c *Cache) Set(key string, val value.Value) {
	c.Lock()
	if v, ok := c.m[key]; ok {
		c.usedBytes += val.Len() - v.Value.(*kv).value.Len()
		kv := v.Value.(*kv)
		kv.value = val
		c.list.MoveToFront(v)
	} else {
		e := c.list.PushFront(&kv{key: key, value: val})
		c.m[key] = e
		c.usedBytes += val.Len()
	}
	c.Unlock()

	for c.maxBytes > 0 && c.usedBytes > c.maxBytes {
		c.removeOld()
	}
}

// delete value according to key
func (c *Cache) Del(key string) {
	c.Lock()
	defer c.Unlock()

	if v, ok := c.m[key]; ok {
		c.list.Remove(v)
		delete(c.m, key)
		c.usedBytes -= v.Value.(*kv).value.Len()
	}
}
func (c *Cache) Scan() map[string]value.Value {
	c.Lock()
	defer c.Unlock()

	resMap := make(map[string]value.Value, len(c.m))
	for k, v := range c.m {
		resMap[k] = v.Value.(*kv).value.(value.StrValue)
	}
	return resMap
}

func (c *Cache) removeOld() {
	c.Lock()
	defer c.Unlock()

	e := c.list.Back()
	if e != nil {
		c.list.Remove(e)
		kv := e.Value.(*kv)
		delete(c.m, kv.key)
		c.usedBytes -= kv.value.Len()
	}
}

func (c *Cache) Len(key string) int64 {
	return int64(c.list.Len())
}
