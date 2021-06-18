package queue

import (
	"github.com/1005281342/godatastructures/queue/deque"
	"github.com/1005281342/godatastructures/queue/ringqueue"
)

type Deque interface {
	Len() int
	Head() interface{}
	Tail() interface{}
	Empty() bool
	LAppend(x interface{}) bool
	Append(x interface{}) bool
	LPop() (interface{}, bool)
	Pop() (interface{}, bool)
}

// NewDeque
// 如果指定了容量则选择循环队列
// 否则选择双端队列
func NewDeque(cap int) Deque {
	if cap > 0 {
		return ringqueue.NewRingQueue(cap)
	}
	return deque.NewDeque()
}
