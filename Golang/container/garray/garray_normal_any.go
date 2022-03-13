package garray

import "sync"

type Array struct {
	mu    sync.RWMutex
	array []interface{}
}

func New() *Array {
	return NewArraySize(0, 0)
}

func NewArraySize(size int, cap int) *Array {
	return &Array{
		mu:    sync.RWMutex{},
		array: make([]interface{}, size, cap),
	}
}

func NewArrayRange(start, end, step int) *Array {
	return nil
}

func NewFrom(array []interface{}) *Array {
	return nil
}

func NewFromCopy(array []interface{}) *Array { return nil }

func NewArrayFrom(array []interface{}) *Array { return nil }

func NewArrayFromCopy(array []interface{}) *Array { return nil }

func (a *Array) At(index int) (value interface{}) { return nil }

func (a *Array) Get(index int) (value interface{}, found bool) { return nil, false }

func (a *Array) Set(index int, value interface{}) error { return nil }

func (a *Array) SetArray(array []interface{}) *Array { return nil }

func (a *Array) Replace(array []interface{}) *Array { return nil }

func (a *Array) Sum() (sum int) { return 0 }

func (a *Array) SortFunc(less func(v1, v2 interface{}) bool) *Array {
	return nil
}

func (a *Array) InsertBefore(index int, value interface{}) error { return nil }

func (a *Array) InsertAfter(index int, value interface{}) error { return nil }

func (a *Array) Remove(index int) (value interface{}, found bool) { return nil, false }

func (a *Array) doRemoveWithoutLock(index int) (value interface{}, found bool) { return nil, false }

func (a *Array) RemoveValue(value interface{}) bool { return false }

func (a *Array) PushLeft(value ...interface{}) *Array { return nil }

func (a *Array) PushRight(value ...interface{}) *Array { return nil }

func (a *Array) PopRand() (value interface{}, found bool) { return nil, false }

func (a *Array) PopRands(size int) []interface{} { return nil }

func (a *Array) PopLeft() (value interface{}, found bool)

func (a *Array) PopRight() (value interface{}, found bool)

func (a *Array) PopLefts(size int) []interface{}

func (a *Array) PopRights(size int) []interface{}
func (a *Array) Range(start int, end ...int) []interface{}
func (a *Array) SubSlice(offset int, length ...int) []interface{}
func (a *Array) Append(value ...interface{}) *Array

func (a *Array) Len() int
func (a *Array) Slice() []interface{}

func (a *Array) Interfaces() []interface{}

func (a *Array) Clone() (newArray *Array)

func (a *Array) Clear() *Array

func (a *Array) Contains(value interface{}) bool

func (a *Array) Search(value interface{}) int
func (a *Array) Unique() *Array

func (a *Array) LockFunc(f func(array []interface{})) *Array

func (a *Array) RLockFunc(f func(array []interface{})) *Array

func (a *Array) Merge(array interface{}) *Array

func (a *Array) Fill(startIndex int, num int, value interface{}) error

func (a *Array) Chunk(size int) [][]interface{}

func (a *Array) Pad(size int, val interface{}) *Array

func (a *Array) Rand() (value interface{}, found bool)

func (a *Array) Rands(size int) []interface{}

func (a *Array) Shuffle() *Array

func (a *Array) Reverse() *Array

func (a *Array) Join(glue string) string

func (a *Array) CountValues() map[interface{}]int

func (a *Array) Iterator(f func(k int, v interface{}) bool)

func (a *Array) IteratorAsc(f func(k int, v interface{}) bool)

func (a *Array) IteratorDesc(f func(k int, v interface{}) bool)

func (a *Array) String() string

func (a Array) MarshalJSON() ([]byte, error)

func (a *Array) UnmarshalJSON(b []byte) error

func (a *Array) UnmarshalValue(value interface{}) error

func (a *Array) FilterNil() *Array

func (a *Array) FilterEmpty() *Array

func (a *Array) Walk(f func(value interface{}) interface{}) *Array

func (a *Array) IsEmpty() bool
