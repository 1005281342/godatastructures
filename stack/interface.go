package stack

import (
	"github.com/1005281342/godatastructures/stack/stack"
	"github.com/1005281342/godatastructures/stack/stackwithcapacity"
)

type Stack interface {
	// Push 添加一个元素
	Push(interface{}) bool
	// BatchPush 批量添加
	BatchPush(...interface{}) bool
	// Pop 弹出栈顶元素
	Pop() (interface{}, bool)
	// Top 获取栈顶元素
	Top() (interface{}, bool)
	// Empty 栈为空
	Empty() bool
	// Len 已使用大小
	Len() int
	//// 可用的
	//Usable() int
	//// 栈已满
	//Full() bool
	// 容量
	//Capacity() int
}

// NewStack 创建一个栈
// if cap <= 0 return 无固定容量的栈
// else return 固定容量的栈
func NewStack(cap int) Stack {
	if cap <= 0 {
		return stack.NewStack()
	}
	return stackwithcapacity.NewStack(cap)
}
