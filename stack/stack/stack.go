package stack

import (
	"container/list"
	"sync"
)

type Stack struct {
	lst  *list.List
	lock sync.RWMutex
}

// NewStack 创建一个栈
func NewStack() *Stack {
	return &Stack{lst: list.New()}
}

// Push 元素入栈
func (stk *Stack) Push(val interface{}) bool {
	stk.lock.Lock()
	defer stk.lock.Unlock()
	stk.lst.PushBack(val)
	return true
}

// Pop 元素出栈
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

// Empty 栈是否为空
func (stk *Stack) Empty() bool {
	return stk.Len() == 0
}

// Top 获取栈顶元素
func (stk *Stack) Top() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	stk.lock.RLock()
	defer stk.lock.RUnlock()
	return stk.lst.Back().Value, true
}

// BatchPush 批量添加
func (stk *Stack) BatchPush(valList ...interface{}) bool {
	stk.lock.Lock()
	defer stk.lock.Unlock()
	for i := 0; i < len(valList); i++ {
		stk.lst.PushBack(valList[i])
	}
	return true
}

// Len 获取栈中元素个数
func (stk *Stack) Len() int {
	return stk.lst.Len()
}
