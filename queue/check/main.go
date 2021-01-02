package main

import (
	"fmt"
	"github.com/1005281342/godatastructures/queue"
)

func main() {

	var dq = queue.NewDeque(-1)
	fmt.Println(dq.Append(1))
	fmt.Println(dq.Append(2))
	fmt.Println(dq.Append(3))
	fmt.Println(dq.Append(4))
	fmt.Println(dq.Pop())
	fmt.Println(dq.LPop())
}
