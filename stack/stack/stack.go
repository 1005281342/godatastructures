package stack

import (
	"container/list"
	"sync"
)

type Stack struct {
	lst  *list.List
	lock sync.RWMutex
}

// NewStack
func NewStack() *Stack {
	return &Stack{lst: list.New()}
}

func (stk *Stack) Push(val interface{}) bool {
	stk.lock.Lock()
	defer stk.lock.Unlock()

	stk.lst.PushBack(val)
	return true
}

func (stk *Stack) Pop() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	stk.lock.Lock()
	defer stk.lock.Unlock()

	// 移除尾节点
	var v = stk.lst.Remove(stk.lst.Back())
	return v, false
}

// Empty栈是否为空
func (stk *Stack) Empty() bool {
	return stk.Len() == 0
}

// Top获取栈顶元素
func (stk *Stack) Top() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	stk.lock.RLock()
	defer stk.lock.RUnlock()

	return stk.lst.Back().Value, true
}

// BatchPush批量添加
func (stk *Stack) BatchPush(valList ...interface{}) bool {
	stk.lock.Lock()
	defer stk.lock.Unlock()

	for i := 0; i < len(valList); i++ {
		stk.lst.PushBack(valList[i])
	}
	return true
}

func (stk *Stack) Len() int {
	return stk.lst.Len()
}
