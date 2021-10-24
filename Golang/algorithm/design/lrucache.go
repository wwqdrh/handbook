package design

import (
	"container/list"
	"errors"
	"sync"
)

// 最久未使用算法（LRU）：最久没有访问的内容作为替换对象
// 新增
// 查询
// 删除

type LruCache struct {
	max   int
	l     *list.List
	cache map[interface{}]*list.Element
	mu    *sync.Mutex // 锁 用于并发安全
}

type node struct {
	Key interface{}
	Val interface{}
}

func NewLruCache(len int) *LruCache {
	return &LruCache{
		max:   len,
		l:     list.New(),
		cache: make(map[interface{}]*list.Element),
		mu:    new(sync.Mutex),
	}
}

func (c *LruCache) Add(key interface{}, val interface{}) error {
	// 在缓存中 表示已经有了 将其从队列中移到队首
	// 不在缓存中加入缓存 并将元素加到队首
	// 若缓存超长 清理队尾缓存数据
	if c.l == nil {
		return errors.New("not init NewLru")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if e, ok := c.cache[key]; ok {
		e.Value.(*node).Val = val
		c.l.MoveToFront(e)
		return nil
	}

	ele := c.l.PushFront(&node{
		Key: key,
		Val: val,
	})
	c.cache[key] = ele
	if c.max != 0 && c.l.Len() > c.max {
		if e := c.l.Back(); e != nil {
			c.l.Remove(e)
			node := e.Value.(*node)
			delete(c.cache, node.Key)
		}
	}
	return nil
}

func (c *LruCache) Get(key interface{}) (val interface{}, ok bool) {
	// 如果缓存中有 将其加入到队首 并返回结果
	// 如果不存在 则返回 nil, false
	if c.cache == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if ele, ok := c.cache[key]; ok {
		c.l.MoveToFront(ele)
		return ele.Value.(*node).Val, true
	}
	return nil, false
}

func (c *LruCache) GetAll() []*node {
	c.mu.Lock()
	defer c.mu.Unlock()
	var data []*node
	for _, v := range c.cache {
		data = append(data, v.Value.(*node))
	}
	return data
}

func (c *LruCache) Del(key interface{}) {
	if c.cache == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.cache[key]; ok {
		if e := c.l.Back(); e != nil {
			c.l.Remove(e)
			delete(c.cache, key)
		}
	}
}
