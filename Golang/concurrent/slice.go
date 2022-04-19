package concurrent

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

// 实现高性能的并发安全的切片读写

// 1、写只需要往后append
// 2、读的时候只需要从现有的range
type SafeSlice struct {
	data  []interface{}
	mutex sync.Mutex
}

type SafeSliceByNode struct {
	head   *sliceNode
	tail   *sliceNode
	length int64
	mutex  sync.Mutex
}

type sliceNode struct {
	Val  interface{}
	Next *sliceNode
}

func NewSafeSlice() *SafeSlice {
	return &SafeSlice{
		data:  []interface{}{},
		mutex: sync.Mutex{},
	}
}

func NewSafeSliceByNode() *SafeSliceByNode {
	return &SafeSliceByNode{
		head:   nil,
		tail:   nil,
		length: 0,
		mutex:  sync.Mutex{},
	}
}

// 并发安全的sliceappend
func (s *SafeSlice) SafeSliceAppend(value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, value)
}

// 并发安全的sliceread
// v1版本：直接遍历，使用range访问
// 测试存在datarace
func (s *SafeSlice) SafeSliceReadV1() string {
	res := []string{}
	for _, item := range s.data {
		res = append(res, fmt.Sprint(item))
	}
	return strings.Join(res, ",")
}

// v2版本: 获取长度，使用下标访问
// 测试存在race
func (s *SafeSlice) SafeSliceReadV2() string {
	length := len(s.data)
	res := []string{}
	for i := 0; i < length; i++ {
		res = append(res, fmt.Sprint(s.data[i]))
	}
	return strings.Join(res, ",")
}

// v3版本: 先进行一次复制，然后遍历访问
// 存在datarace
func (s *SafeSlice) SafeSliceReadV3() string {
	newData := s.data[:]
	res := []string{}
	for _, item := range newData {
		res = append(res, fmt.Sprint(item))
	}
	return strings.Join(res, ",")
}

func (s *SafeSliceByNode) SafeSliceAppend(value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cur := &sliceNode{
		Val: value,
	}
	if s.length == 0 {
		s.head = cur
		s.tail = cur
		atomic.AddInt64(&s.length, 1)
	} else {
		s.tail.Next = cur
		s.tail = cur
		atomic.AddInt64(&s.length, 1)
	}
}

func (s *SafeSliceByNode) SafeSliceRead() string {
	length := atomic.LoadInt64(&s.length)
	res := []string{}

	cur := s.head
	for i := 0; i < int(length); i++ {
		res = append(res, fmt.Sprint(cur.Val))
		cur = cur.Next
	}
	return strings.Join(res, ",")
}
