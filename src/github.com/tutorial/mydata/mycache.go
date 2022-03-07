package mydata

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	mutex sync.RWMutex
	data map[string]*list.Element
	help *list.List
	cap int
}

type entry struct {
	key string
	value interface{}
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		data: make(map[string]*list.Element),
		help: list.New(),
		cap: cap,
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// key 不存在
	ele, ok := c.data[key]
	if !ok {
		return nil, false
	} 
	
	c.help.MoveToFront(ele)
	return ele.Value.(*entry).value, true
}


func (c *LRUCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// key 已经存在
	if ele, ok := c.data[key]; ok {
		c.help.MoveToFront(ele)
		c.data[key].Value.(*entry).value = value
		return
	} 
	
	ele := c.help.PushFront(&entry{key: key, value: value})
	// 到达容量上限
	if c.cap != 0 && c.help.Len() > c.cap {
		old := c.help.Back()
		c.help.Remove(old)
		delete(c.data, old.Value.(*entry).key)
	}
	c.data[key] = ele
}

func (c *LRUCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// key 不存在
	ele, ok := c.data[key]
	if !ok {
		return
	} 
	
	c.help.Remove(ele)
	delete(c.data, ele.Value.(*entry).key)
}
