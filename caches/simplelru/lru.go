package simplelru

import (
	"container/list"
	"errors"
	"tst/caches"
)


type LruCache struct {
	size int
	backend *list.List
	items     map[interface{}]*list.Element
}

type entry struct {
	key interface{}
	value interface{}
}

func NewLRU(size int) (caches.LRUCacher, error) {
	if size <= 0 {
		return nil, errors.New("Must provide a positive size")
	}
	c := &LruCache{
		size:      size,
		backend: list.New(),
		items:     make(map[interface{}]*list.Element),
	}
	return c, nil
}

func (c *LruCache) Purge() {
	for k  := range c.items {
		delete(c.items, k)
	}
	c.backend.Init()
}

func (c *LruCache) Add(key, value interface{}) (evicted bool) {
	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.backend.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return false
	}

	// Add new item
	entry := c.backend.PushFront(&entry{key, value})
	c.items[key] = entry

	evict := c.backend.Len() > c.size
	// Verify size not exceeded
	if evict {
		c.removeOldest()
	}
	return evict
}

func (c *LruCache) Get(key interface{}) (value interface{}, ok bool){
	 if v, ok := c.items[key]; ok {
	 	c.backend.MoveToFront(v)
	 	if v.Value.(*entry) == nil {
	 		return nil, false
		}
	 	return v.Value.(*entry).value, true
	}
 	return nil, false
}

func (c *LruCache) Contains(key interface{}) (ok bool) {
	_, ok = c.items[key]
	return ok
}

func (c *LruCache) Peek(key interface{}) (value interface{}, ok bool){
	if v, ok := c.items[key]; ok {
		return v.Value.(*entry).value, true
	}
	return nil, false
}

func (c *LruCache) Remove(key interface{}) bool {
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent, key)
		return true
	}
	return false
}

func (c *LruCache) removeElement(e *list.Element, key interface{}) {
	c.backend.Remove(e)
	delete(c.items, key)
}

func (c *LruCache) RemoveOldest() (key interface{}, value interface{}, ok bool) {
	ent := c.backend.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		c.removeElement(ent, kv.key)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}
func (c *LruCache) removeOldest() {
	ent := c.backend.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		c.removeElement(ent, kv.key)
	}
}
func (c *LruCache) GetOldest() (key interface{}, value interface{}, ok bool) {
	ent := c.backend.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

func (c *LruCache) Keys() []interface{} {
	keys := make([]interface{}, len(c.items))
	i := 0
	for ent := c.backend.Back(); ent != nil; ent = ent.Prev() {
		keys[i] = ent.Value.(*entry).key
		i++
	}
	return keys
}

func (c *LruCache) Len() int {
	return c.backend.Len()
}

func (c *LruCache) Resize(size int) (evicted int) {
	diff := c.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		c.removeOldest()
	}
	c.size = size
	return diff
}


