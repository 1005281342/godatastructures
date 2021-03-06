package heap

type IntHeap struct {
	items []int
	size  int
}

func NewIntHeap(items ...int) *IntHeap {
	var hp = &IntHeap{items: items, size: len(items)}
	hp.initIntHeap()
	return hp
}

func (h *IntHeap) Len() int {
	return h.size
}

func (h *IntHeap) Empty() bool {
	return h.Len() == 0
}

func (h *IntHeap) Top() (int, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

func (h *IntHeap) Pop() (int, bool) {
	if h.Empty() {
		return 0, false
	}

	var ans = h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.down(0)
	return ans, true
}

func (h *IntHeap) Remove(index int) (int, bool) {
	if h.Empty() {
		return 0, false
	}
	var ans = h.items[index]
	h.items[index] = h.items[h.size-1]
	h.size--
	h.down(index)
	h.up(index)
	return ans, true
}

// Push
func (h *IntHeap) Push(v int) {
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

func (h *IntHeap) up(u int) {
	for {
		var root = (u - 1) >> 1
		if root < 0 || h.items[root] <= h.items[u] {
			break
		}
		h.items[u], h.items[root] = h.items[root], h.items[u]
		u = root
	}
}
