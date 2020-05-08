package simplelfu

import (
	"container/list"
)


type entry struct {
	key interface{}
	value interface{}
	score int
}

type LfuCache struct {
	size int
	backend *list.List
	items     map[interface{}]*list.Element
}



func (c *LfuCache) Add(key, value interface{}) bool {
	// Checking does element already exists
	if ent, ok := c.items[key]; ok {
		// if exists, we increase its score, move to next position if its score is greater then next ones score
		c.increaseScoreAndMove(ent)
		ent.Value.(*entry).value = value
		return false
	}
	// If element does not exists remove element with lower score in case len limit is faced

	evict := c.backend.Len() >= c.size
	if evict {
		c.removeLowestScore()
	}
	//adding new element to the back of array
	entry := c.backend.PushBack(&entry{key, value, 0})
	c.items[key] = entry

	return evict
}

func (c *LfuCache) Get(key interface{}) (value interface{}, ok bool) {
	if v, ok := c.items[key]; ok {
		c.increaseScoreAndMove(v)
		if v.Value.(*entry) == nil {
			return nil ,false
		}
		return v.Value.(*entry).value, true
	}
	return nil, false
}

func (c *LfuCache) increaseScoreAndMove(e *list.Element) {
	nextGreater := c.findNextGreatestScore(e)
	if nextGreater != nil {
		c.backend.MoveBefore(e, nextGreater)
	} else {
		c.backend.MoveToFront(e)
	}
}

func (c *LfuCache) findNextGreatestScore(e *list.Element) *list.Element {
	// if no next element, means our element have greater score
	if e.Next() == nil {
		return nil
	}
	if e.Value.(*entry).score >= e.Next().Value.(*entry).score {
		return c.findNextGreatestScore(e.Next())
	} else {
		return e.Next()
	}
}

func (c *LfuCache) removeLowestScore() {
	ent := c.backend.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		c.backend.Remove(ent)
		delete(c.items, kv.key)
	}
}