package heap

import (
	"fmt"
	"testing"
)

func TestNewHeap(t *testing.T) {
	var hp = NewHeap(2, 4, 1, 3, 2, 9, 3, 8)
	fmt.Println(hp.items)
}
