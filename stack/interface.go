package stack

import (
	"github.com/1005281342/godatastructures/stack/stack"
	"github.com/1005281342/godatastructures/stack/stackwithcapacity"
)

type Stack interface {
	// 添加一个元素
	Push(interface{}) bool
	// 批量添加
	BatchPush(...interface{}) bool
	// 弹出栈顶元素
	Pop() (interface{}, bool)
	// 获取栈顶元素
	Top() (interface{}, bool)

	// 栈为空
	Empty() bool
	// 已使用大小
	Len() int
	//// 可用的
	//Usable() int
	//// 栈已满
	//Full() bool
	// 容量
	//Capacity() int
}

func NewStack(cap int) Stack {

	if cap <= 0 {
		return stack.NewStack()
	}
	return stackwithcapacity.NewStack(cap)
}
