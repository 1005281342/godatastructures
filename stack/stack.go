package stack

type Stack interface {
	// 添加一个元素
	Push(interface{}) bool
	// 批量添加
	BatchPush(...interface{}) bool
	// 弹出栈顶元素
	Pop() (interface{}, bool)
	// 获取栈顶元素
	Top() (interface{}, bool)
	// 容量
	Capacity() int
	// 已使用大小
	Len() int
	// 可用的
	Usable() int
	// 栈为空
	Empty() bool
	// 栈已满
	Full() bool
}

type Stk struct {
	slice []interface{}
	// 从0开始记
	capacity int
	// 从0开始记
	use int
}

func NewStack(cap int) Stack {

	if cap <= 0 {
		panic("cap <= 0")
	}
	return &Stk{capacity: cap - 1, slice: make([]interface{}, cap, cap), use: 0}
}

func (stk *Stk) Capacity() int {
	return stk.capacity + 1
}

func (stk *Stk) Len() int {
	return stk.use + 1
}

// Full栈是否已满
func (stk *Stk) Full() bool {
	return stk.use >= stk.capacity
}

// Push栈有空间时可添加
func (stk *Stk) Push(val interface{}) bool {

	if stk.Full() {
		return false
	}
	stk.slice[stk.use] = val
	stk.use++
	return true
}

// Usable计算栈的可用空间
func (stk *Stk) Usable() int {
	var usable = stk.capacity - stk.use
	if usable >= 0 {
		return usable
	}
	// TODO 什么情况会走到这个逻辑
	stk.use = stk.capacity
	return 0
}

// BatchPush批量添加
func (stk *Stk) BatchPush(valList ...interface{}) bool {

	if len(valList) > stk.Usable() {
		return false
	}

	for i := 0; i < len(valList); i++ {
		stk.slice[stk.use] = valList[i]
		stk.use++
	}
	return true
}

// Empty栈是否为空
func (stk *Stk) Empty() bool {
	return stk.use == 0
}

// Pop弹出元素
func (stk *Stk) Pop() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}

	var v = stk.slice[stk.use]
	stk.use--
	return v, true
}

// Top获取栈顶元素
func (stk *Stk) Top() (interface{}, bool) {
	if stk.Empty() {
		return nil, false
	}
	return stk.slice[stk.use-1], true
}
