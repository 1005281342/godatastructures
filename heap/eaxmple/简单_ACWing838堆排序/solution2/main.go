package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 验证地址 https://www.acwing.com/problem/content/840/
func main() {
	var reader = bufio.NewReader(os.Stdin)
	var d = readArray(reader)
	var t = make([]int, d[1])
	var x = readArray(reader)
	var hp = NewIntHeap(x...)
	for i := 0; i < d[1]; i++ {
		var v, _ = hp.Pop()
		t[i] = v
	}
	fmt.Println(nums2string(t, " "))
}

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

// IntHeap int heap
type IntHeap struct {
	items []int
	size  int
}

// NewIntHeap new int heap
func NewIntHeap(items ...int) *IntHeap {
	var hp = &IntHeap{items: items, size: len(items)}
	hp.initIntHeap()
	return hp
}

// Len len
func (h *IntHeap) Len() int {
	return h.size
}

// Empty empty
func (h *IntHeap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *IntHeap) Top() (int, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *IntHeap) Pop() (int, bool) {
	if h.Empty() {
		return 0, false
	}

	// 缓存从堆中移除的节点值
	var ans = h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.down(0)
	return ans, true
}

// Remove remove
func (h *IntHeap) Remove(index int) (int, bool) {
	if h.Empty() {
		return 0, false
	}
	var ans = h.items[index]
	h.items[index] = h.items[h.size-1]
	h.size--
	// 实际上down或up只会执行一个
	h.down(index)
	h.up(index)
	return ans, true
}

// Push push
func (h *IntHeap) Push(v int) {
	if h.size < len(h.items) {
		// 不需要扩容
		h.size++
		h.items[h.size-1] = v
		h.up(h.size - 1)
		return
	}

	// append
	h.items = append(h.items, v)
	h.size = len(h.items)
	h.up(h.size - 1)
}

func (h *IntHeap) initIntHeap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *IntHeap) down(u int) {
	var (
		t     = u
		left  = 2*u + 1
		right = left + 1
	)
	if left < h.size && h.items[left] < h.items[t] {
		// 左节点值大于根节点
		t = left
	}
	if right < h.size && h.items[right] < h.items[t] {
		// 右节点值大于根节点
		t = right
	}
	if t != u {
		h.items[t], h.items[u] = h.items[u], h.items[t]
		h.down(t)
	}
}

func (h *IntHeap) up(u int) {
	for {
		var root = (u - 1) >> 1
		if root < 0 || h.items[root] <= h.items[u] {
			break
		}
		// 走到这里意味着h.items[root]>h.items[u]，因此交换节点值
		h.items[u], h.items[root] = h.items[root], h.items[u]
		u = root
	}
}
