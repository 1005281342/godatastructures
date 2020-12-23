package stackwithcapacity

import "sync"

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

func (stk *Stack) Capacity() int {
	return stk.capacity + 1
}

func (stk *Stack) Len() int {
	return stk.use
}

// Full栈是否已满
func (stk *Stack) Full() bool {
	return stk.use >= stk.capacity
}

// Push栈有空间时可添加
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

// Usable计算栈的可用空间
func (stk *Stack) Usable() int {

	var usable = stk.capacity - stk.use
	if usable >= 0 {
		return usable
	}
	stk.use = stk.capacity
	return 0
}

// BatchPush批量添加
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

// Empty栈是否为空
func (stk *Stack) Empty() bool {
	return stk.use == 0
}

// Pop弹出元素
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

// Top获取栈顶元素
func (stk *Stack) Top() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	stk.lock.RLock()
	defer stk.lock.RUnlock()

	return stk.slice[stk.use-1], true
}
