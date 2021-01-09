package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x = []int{1, 0, 0}
	var y = []struct{}{{}, {}, {}}
	var z = []bool{true, false, false}
	var h = []float64{1, 0, 0}
	fmt.Println("x size", unsafe.Sizeof(x))
	fmt.Println("y size", unsafe.Sizeof(y))
	fmt.Println("z size", unsafe.Sizeof(z))
	fmt.Println("h size", unsafe.Sizeof(h))
}
