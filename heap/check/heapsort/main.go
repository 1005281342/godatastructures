package main

// 验证地址https://www.acwing.com/problem/content/840/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var d = readArray(reader)
	var t = make([]int, d[1])
	var x = readArray(reader)
	var hp = NewHeap(x...)
	for i := 0; i < d[1]; i++ {
		var v, _ = hp.Pop()
		t[i] = v
	}
	fmt.Println(nums2string(t, " "))
}

type Heap struct {
	items []int
	size  int
}

func NewHeap(items ...int) *Heap {
	var hp = &Heap{items: items, size: len(items)}
	hp.initHeap()
	return hp
}

func (h *Heap) Pop() (int, bool) {
	if h.Empty() {
		return 0, false
	}

	var ans = h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.down(0)
	return ans, true
}

func (h *Heap) Empty() bool {
	return h.Len() == 0
}

func (h *Heap) Len() int {
	return h.size
}

func (h *Heap) initHeap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Heap) down(u int) {
	var (
		t     = u
		left  = 2*u + 1
		right = left + 1
	)
	if left < h.size && h.items[left] < h.items[t] {
		t = left
	}
	if right < h.size && h.items[right] < h.items[t] {
		t = right
	}
	if t != u {
		h.items[t], h.items[u] = h.items[u], h.items[t]
		h.down(t)
	}
}

//func (h *Heap) up(u int) {
//	for {
//		var root = (u - 1) >> 1
//		if root < 0 || h.items[root] >= h.items[u] {
//			break
//		}
//		h.items[u], h.items[root] = h.items[root], h.items[u]
//		u = root
//	}
//}

func readLine(reader *bufio.Reader) string {
	var line, _ = reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

//func readInt(reader *bufio.Reader) int {
//	var num, _ = strconv.Atoi(readLine(reader))
//	return num
//}

func readArray(reader *bufio.Reader) []int {
	var line = readLine(reader)
	var strList = strings.Split(line, " ")
	var nums = make([]int, 0)
	var err error
	var v int
	for i := 0; i < len(strList); i++ {
		if v, err = strconv.Atoi(strList[i]); err != nil {
			continue
		}
		nums = append(nums, v)
	}
	return nums
}

func nums2string(x []int, sep string) string {
	var b strings.Builder
	for i := 0; i < len(x); i++ {
		b.WriteString(strconv.Itoa(x[i]))
		b.WriteString(sep)
	}
	return b.String()
}
