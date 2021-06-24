package heap

// Int64Heap int heap
type Int64Heap struct {
	items []int64
	size  int
}

// NewInt64Heap new int heap
func NewInt64Heap(items ...int64) *Int64Heap {
	var hp = &Int64Heap{items: items, size: len(items)}
	hp.initInt64Heap()
	return hp
}

// Len len
func (h *Int64Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Int64Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Int64Heap) Top() (int64, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Int64Heap) Pop() (int64, bool) {
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
func (h *Int64Heap) Remove(index int) (int64, bool) {
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
func (h *Int64Heap) Push(v int64) {
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

func (h *Int64Heap) initInt64Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Int64Heap) down(u int) {
	var (
		t     = u
		left  = 2*u + 1
		right = left + 1
	)
	// 在根节点、左节点、右节点三个节点中选择最大的节点
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

func (h *Int64Heap) up(u int) {
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
