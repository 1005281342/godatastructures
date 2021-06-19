package main

import (
	"sync"
)

// 遇到"abc"则出栈，最后校验栈是否为空
func isValid(s string) bool {
	var (
		stk   = NewStack(len(s) + 10)
		value interface{}
		c     byte
		ok    bool
	)
	for i := 0; i < len(s); i++ {
		if stk.Len() < 2 {
			stk.Push(s[i])
			continue
		}
		if s[i] != 'c' {
			stk.Push(s[i])
			continue
		}
		value, _ = stk.Top()
		if c, ok = value.(byte); !ok {
			return false
		}
		if c == 'b' {
			stk.Pop()
		} else {
			stk.Push('c')
			continue
		}
		value, _ = stk.Top()
		if c, ok = value.(byte); !ok {
			return false
		}
		if c == 'a' {
			stk.Pop()
			continue
		} else {
			stk.Push('b')
			stk.Push('c')
		}
	}
	return stk.Len() == 0
}

type Stack struct {
	slice []interface{}
	// 从0开始记
	capacity int
	// 从0开始记
	use  int
	lock sync.RWMutex
}

func NewStack(capacity int) *Stack {
	return &Stack{slice: make([]interface{}, capacity), capacity: capacity - 1, use: 0}
}

// Capacity 栈容量大小
func (stk *Stack) Capacity() int {
	return stk.capacity + 1
}

// Len 栈中元素个数
func (stk *Stack) Len() int {
	return stk.use
}

// Full 栈是否已满
func (stk *Stack) Full() bool {
	return stk.use >= stk.capacity
}

// Push 栈有空间时可添加
func (stk *Stack) Push(val interface{}) bool {
	if stk.Full() {
		return false
	}

	stk.lock.Lock()
	defer stk.lock.Unlock()
	stk.slice[stk.use] = val
	stk.use++
	return true
}

// Usable 计算栈的可用空间
func (stk *Stack) Usable() int {
	var usable = stk.capacity - stk.use
	if usable >= 0 {
		return usable
	}
	stk.use = stk.capacity
	return 0
}

// BatchPush 批量添加
func (stk *Stack) BatchPush(valList ...interface{}) bool {
	if len(valList) > stk.Usable() {
		return false
	}

	stk.lock.Lock()
	defer stk.lock.Unlock()
	for i := 0; i < len(valList); i++ {
		stk.slice[stk.use] = valList[i]
		stk.use++
	}
	return true
}

// Empty 栈是否为空
func (stk *Stack) Empty() bool {
	return stk.use == 0
}

// Pop 弹出元素
func (stk *Stack) Pop() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	stk.lock.Lock()
	defer stk.lock.Unlock()
	var v = stk.slice[stk.use]
	stk.use--
	return v, true
}

// Top 获取栈顶元素
func (stk *Stack) Top() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	stk.lock.RLock()
	defer stk.lock.RUnlock()
	return stk.slice[stk.use-1], true
}

//type Stack struct {
//	lst  *list.List
//	lock sync.RWMutex
//}
//
//// NewStack 创建一个栈
//func NewStack() *Stack {
//	return &Stack{lst: list.New()}
//}
//
//// Push 元素入栈
//func (stk *Stack) Push(val interface{}) bool {
//	stk.lock.Lock()
//	defer stk.lock.Unlock()
//	stk.lst.PushBack(val)
//	return true
//}
//
//// Pop 元素出栈
//func (stk *Stack) Pop() (interface{}, bool) {
//	if stk.Empty() {
//		return nil, false
//	}
//
//	stk.lock.Lock()
//	defer stk.lock.Unlock()
//	// 移除尾节点
//	var v = stk.lst.Remove(stk.lst.Back())
//	return v, false
//}
//
//// Empty 栈是否为空
//func (stk *Stack) Empty() bool {
//	return stk.Len() == 0
//}
//
//// Top 获取栈顶元素
//func (stk *Stack) Top() (interface{}, bool) {
//	if stk.Empty() {
//		return nil, false
//	}
//
//	stk.lock.RLock()
//	defer stk.lock.RUnlock()
//	return stk.lst.Back().Value, true
//}
//
//// BatchPush 批量添加
//func (stk *Stack) BatchPush(valList ...interface{}) bool {
//	stk.lock.Lock()
//	defer stk.lock.Unlock()
//	for i := 0; i < len(valList); i++ {
//		stk.lst.PushBack(valList[i])
//	}
//	return true
//}
//
//// Len 获取栈中元素个数
//func (stk *Stack) Len() int {
//	return stk.lst.Len()
//}

// https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions/
