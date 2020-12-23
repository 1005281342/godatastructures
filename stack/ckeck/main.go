package main

import (
	"fmt"
	"github.com/1005281342/godatastructures/stack"
	"math"
)

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())

}

// https://leetcode-cn.com/problems/min-stack/
// 执行用时：16ms,在所有Go提交中击败了95.97% 的用户
// 内存消耗：8.3MB,在所有Go提交中击败了92.03%的用户
type MinStack struct {
	stk stack.Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	//return MinStack{stk: stack.NewStack(10000)}
	return MinStack{stk: stack.NewStack(-1)}
}

func (m *MinStack) Push(x int) {

	var y = x
	if y > m.GetMin() {
		y = m.GetMin()
	}
	m.stk.Push([2]int{x, y})
}

func (m *MinStack) Pop() {
	m.stk.Pop()
}

func (m *MinStack) Top() int {
	v, _ := m.stk.Top()
	return v.([2]int)[0]
}

func (m *MinStack) GetMin() int {
	if m.stk.Empty() {
		return math.MaxInt32
	}

	v, _ := m.stk.Top()
	return v.([2]int)[1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(x)
 * obj.Pop()
 * param_3 := obj.Top()
 * param_4 := obj.GetMin()
 */

//基于链表 执行用时：20ms,在所有Go提交中击败了81.70%的用户
//内存消耗：9.7MB,在所有Go提交中击败了13.12%的用户
