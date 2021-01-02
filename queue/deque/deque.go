package deque

import (
	"container/list"
	"sync"
)

type Deque struct {
	// 队列
	lst  *list.List
	lock sync.RWMutex
}

func NewDeque() *Deque {
	return &Deque{lst: list.New()}
}

// Empty 是否为空
func (dq *Deque) Empty() bool {
	dq.lock.RLock()
	defer dq.lock.RUnlock()

	return 0 == dq.lst.Len()
}

// Len 当前队列长度
func (dq *Deque) Len() int {
	return dq.lst.Len()
}

// LAppend 从左边添加元素, return是否添加成功
func (dq *Deque) LAppend(v interface{}) bool {
	dq.lock.Lock()
	defer dq.lock.Unlock()

	var e = dq.lst.PushFront(v)
	return e != nil
}

// LPop 从左边移除元素, return元素值、是否移除成功
func (dq *Deque) LPop() (interface{}, bool) {
	if dq.Empty() {
		return nil, false
	}

	dq.lock.Lock()
	defer dq.lock.Unlock()

	var v = dq.lst.Remove(dq.lst.Front())
	return v, true
}

// Append 从右边添加元素, return是否添加成功
func (dq *Deque) Append(v interface{}) bool {

	dq.lock.Lock()
	defer dq.lock.Unlock()

	var e = dq.lst.PushBack(v)
	return e != nil
}

// Pop 从右边移除元素, return元素值、是否移除成功
func (dq *Deque) Pop() (interface{}, bool) {
	if dq.Empty() {
		return nil, false
	}

	dq.lock.Lock()
	defer dq.lock.Unlock()

	var v = dq.lst.Remove(dq.lst.Back())
	return v, true
}

// Head 获取最左边元素
func (dq *Deque) Head() interface{} {
	return dq.lst.Front()
}

// Tail 获取最右边元素
func (dq *Deque) Tail() interface{} {
	return dq.lst.Back()
}
