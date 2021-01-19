package skiplist

import (
	"fmt"
	"math"
	"testing"
)

func TestExample(t *testing.T) {
	var sp = Constructor()
	sp.Add(1)
	sp.Add(3)
	if !sp.Search(3) {
		t.Fail()
	}

	if sp.Search(2) {
		t.Fail()
	}
	fmt.Println(sp.Search(math.MinInt32))
	sp.Add(math.MinInt32)
	fmt.Println(sp.Search(math.MinInt32))
}
