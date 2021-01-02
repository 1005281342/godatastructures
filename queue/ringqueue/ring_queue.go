package ringqueue

import (
	"sync"
)

type RingQueue struct {
	front  int           // 指向队列头部第1个有效数据的位置
	rear   int           // 指向队列尾部（即最后1个有效数据）的下一个位置，即下一个从队尾入队元素的位置
	length int           // 队列长度，非容量，slice动态库容
	nums   []interface{} // 队列元素
	// lock
	lock sync.RWMutex
}

// NewRingQueue
func NewRingQueue(k int) *RingQueue {
	var length = k + 1
	return &RingQueue{
		length: length,
		nums:   make([]interface{}, length),
	}
}

// Len
func (q *RingQueue) Len() int {

	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.front > q.rear {
		return q.rear - q.front + q.length
	}
	return q.rear - q.front
}

// Head
func (q *RingQueue) Head() interface{} {
	if q.Empty() {
		return nil
	}

	return q.nums[q.front]
}

// Tail
func (q *RingQueue) Tail() interface{} {
	if q.Empty() {
		return nil
	}
	var pos int // 其实就是rear-1 //pos = (q.rear - 1 + q.length) % q.length
	if q.rear > 0 {
		pos = q.rear - 1
	} else if q.rear == 0 {
		pos = q.length - 1
	}
	return q.nums[pos]
}

// LAppend
func (q *RingQueue) LAppend(x interface{}) bool {
	if q.IsFull() {
		return false
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	if q.front == 0 {
		q.front = q.length - 1
	} else {
		q.front -= 1
	}
	q.nums[q.front] = x
	return true
}

// Append
func (q *RingQueue) Append(x interface{}) bool {
	if q.IsFull() {
		return false
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	q.nums[q.rear] = x
	q.rear++ // q.rear = (q.rear + 1) % q.length // rear指针后移一位
	if q.rear == q.length {
		q.rear = 0
	}
	return true
}

// Pop
func (q *RingQueue) Pop() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	if q.rear == 0 {
		q.rear = q.length - 1
	} else {
		q.rear -= 1
	}
	// q.rear指向队列尾部（即最后1个有效数据）的下一个位置，即下一个从队尾入队元素的位置
	// 上一个被Pop的元素即为q.rear
	return q.nums[q.rear], true
}

// LPop
func (q *RingQueue) LPop() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	var v = q.nums[q.front]
	if q.front == q.length-1 { // front指针后移一位 // q.front = (q.front + 1) % q.length
		q.front = 0
	} else {
		q.front += 1
	}
	return v, true
}

// Empty 队列是否为空
func (q *RingQueue) Empty() bool {
	return q.front == q.rear
}

// IsFull 队列是否已满
func (q *RingQueue) IsFull() bool {
	return (q.rear+1)%q.length == q.front
}
