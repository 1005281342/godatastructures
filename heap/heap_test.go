package heap

import (
	"fmt"
	"testing"
)

func TestNewIntHeap(t *testing.T) {
	var hp = NewIntHeap(2, 4, 1, 3, 2, 9, 3, 8)
	fmt.Println(hp.items)
}
