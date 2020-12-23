package stackwithinfinite

import (
	"container/list"
	"sync"
)

type Stack struct {
	lst  *list.List
	tail *list.Element
	// 从0开始记
	use  int
	lock sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{lst: list.New()}
}

func (stk *Stack) Push(val interface{}) bool {
	stk.lock.Lock()
	defer stk.lock.Unlock()

	stk.tail = stk.lst.PushBack(val)
	stk.use++
	return true
}

func (stk *Stack) Pop() (interface{}, bool) {
	if stk.tail == nil {
		return nil, false
	}

	stk.lock.Lock()
	defer stk.lock.Unlock()

	// 尾节点的前缀节点
	var tt = stk.tail.Prev()
	// 移除尾节点
	var v = stk.lst.Remove(stk.tail)
	stk.use--
	// 更新尾节点
	stk.tail = tt
	return v, false
}

// Empty栈是否为空
func (stk *Stack) Empty() bool {
	return stk.use == 0
}

// Top获取栈顶元素
func (stk *Stack) Top() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	stk.lock.RLock()
	defer stk.lock.RUnlock()

	return stk.tail.Value, true
}

// BatchPush批量添加
func (stk *Stack) BatchPush(valList ...interface{}) bool {

	stk.lock.Lock()
	defer stk.lock.Unlock()

	for i := 0; i < len(valList)-1; i++ {
		stk.lst.PushBack(valList[i])
	}
	stk.tail = stk.lst.PushBack(valList[len(valList)-1])

	stk.use += len(valList)
	return true
}

func (stk *Stack) Len() int {
	return stk.use
}
