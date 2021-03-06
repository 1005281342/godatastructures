package main

import "fmt"

// 验证地址https://leetcode-cn.com/problems/top-k-frequent-elements/

// TopKFrequent
func main() {
	topKFrequent([]int{1, 2, 3, 4, 6}, 3)
}

func topKFrequent(nums []int, k int) []int {
	var mp = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	fmt.Println(mp)
	var hp = NewHeap()
	for xk, v := range mp {
		hp.Push(Value{priority: -v, v: xk})
		hp.Debug()
	}
	var ans = make([]int, k)
	for i := 0; i < k; i++ {
		var xv, _ = hp.Pop()
		hp.Debug()
		ans[i] = xv
	}
	return ans
}

type Value struct {
	priority int
	v        int
}

type Heap struct {
	items []Value
	size  int
}

func NewHeap() *Heap {
	var hp = &Heap{items: make([]Value, 0), size: 0}
	hp.initHeap()
	return hp
}

func (h *Heap) Debug() {
	fmt.Println(h.items[:h.size])
}

func (h *Heap) Len() int {
	return h.size
}

func (h *Heap) Empty() bool {
	return h.Len() == 0
}

func (h *Heap) Top() (int, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0].v, true
}

func (h *Heap) Pop() (int, bool) {
	if h.Empty() {
		return 0, false
	}

	var ans = h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.down(0)
	return ans.v, true
}

func (h *Heap) Remove(index int) (int, bool) {
	if h.Empty() {
		return 0, false
	}
	var ans = h.items[index]
	h.items[index] = h.items[h.size-1]
	h.size--
	h.down(index)
	h.up(index)
	return ans.v, true
}

// Push
func (h *Heap) Push(v Value) {
	if h.size < len(h.items) {
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
	if left < h.size && h.items[left].priority < h.items[t].priority {
		t = left
	}
	if right < h.size && h.items[right].priority < h.items[t].priority {
		t = right
	}
	if t != u {
		h.items[t], h.items[u] = h.items[u], h.items[t]
		h.down(t)
	}
}

func (h *Heap) up(u int) {

	for {
		var root = (u - 1) >> 1
		if root < 0 || h.items[root].priority <= h.items[u].priority {
			break
		}
		h.items[u], h.items[root] = h.items[root], h.items[u]
		u = root
	}
}
