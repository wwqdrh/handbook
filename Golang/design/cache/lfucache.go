package cache

import (
	"container/list"
)

type LfuVc struct {
	value int
	count int
}

type LFUCache struct {
	mapKvs  map[int]LfuVc
	mapFreq map[int]*list.List
	min     int
	cap     int
}

func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		mapKvs:  make(map[int]LfuVc, capacity),
		mapFreq: make(map[int]*list.List, capacity),
		min:     0,
		cap:     capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	lfuVcElem, ok := this.mapKvs[key]
	if !ok {
		return -1
	}

	this.mapKvs[key] = LfuVc{value: lfuVcElem.value, count: lfuVcElem.count + 1}
	cnt := lfuVcElem.count

	for e := this.mapFreq[cnt].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.mapFreq[cnt].Remove(e)
			break
		}
	}
	if this.mapFreq[cnt+1] == nil {
		this.mapFreq[cnt+1] = list.New()
	}
	this.mapFreq[cnt+1].PushBack(key)
	//update min
	if cnt == this.min && this.mapFreq[cnt].Len() == 0 {
		this.min++ //由于被查的key的count由min升级到min+1，所以min+1必然非空
	}

	return lfuVcElem.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	_, ok := this.mapKvs[key]
	if ok {
		this.mapKvs[key] = LfuVc{value: value, count: this.mapKvs[key].count} //更新count留给this.Get()做
		this.Get(key)
		return
	}
	if len(this.mapKvs) >= this.cap {
		//if this.min > 1 {
		//  //满容，且都是高频元素，不执行Put
		//  return
		//}
		l := this.mapFreq[this.min]
		delete(this.mapKvs, l.Front().Value.(int))
		l.Remove(l.Front())
	}
	this.mapKvs[key] = LfuVc{value: value, count: 1}
	if this.min != 1 {
		this.mapFreq[1] = list.New()
	}
	this.mapFreq[1].PushBack(key)
	this.min = 1
}
