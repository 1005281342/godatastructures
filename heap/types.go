package heap

// Float32Heap int heap
type Float32Heap struct {
	items []float32
	size  int
}

// NewFloat32Heap new int heap
func NewFloat32Heap(items ...float32) *Float32Heap {
	var hp = &Float32Heap{items: items, size: len(items)}
	hp.initFloat32Heap()
	return hp
}

// Len len
func (h *Float32Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Float32Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Float32Heap) Top() (float32, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Float32Heap) Pop() (float32, bool) {
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
func (h *Float32Heap) Remove(index int) (float32, bool) {
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
func (h *Float32Heap) Push(v float32) {
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

func (h *Float32Heap) initFloat32Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Float32Heap) down(u int) {
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

func (h *Float32Heap) up(u int) {
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

// Float64Heap int heap
type Float64Heap struct {
	items []float64
	size  int
}

// NewFloat64Heap new int heap
func NewFloat64Heap(items ...float64) *Float64Heap {
	var hp = &Float64Heap{items: items, size: len(items)}
	hp.initFloat64Heap()
	return hp
}

// Len len
func (h *Float64Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Float64Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Float64Heap) Top() (float64, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Float64Heap) Pop() (float64, bool) {
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
func (h *Float64Heap) Remove(index int) (float64, bool) {
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
func (h *Float64Heap) Push(v float64) {
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

func (h *Float64Heap) initFloat64Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Float64Heap) down(u int) {
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

func (h *Float64Heap) up(u int) {
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

// Int32Heap int heap
type Int32Heap struct {
	items []int32
	size  int
}

// NewInt32Heap new int heap
func NewInt32Heap(items ...int32) *Int32Heap {
	var hp = &Int32Heap{items: items, size: len(items)}
	hp.initInt32Heap()
	return hp
}

// Len len
func (h *Int32Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Int32Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Int32Heap) Top() (int32, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Int32Heap) Pop() (int32, bool) {
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
func (h *Int32Heap) Remove(index int) (int32, bool) {
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
func (h *Int32Heap) Push(v int32) {
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

func (h *Int32Heap) initInt32Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Int32Heap) down(u int) {
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

func (h *Int32Heap) up(u int) {
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

// Int16Heap int heap
type Int16Heap struct {
	items []int16
	size  int
}

// NewInt16Heap new int heap
func NewInt16Heap(items ...int16) *Int16Heap {
	var hp = &Int16Heap{items: items, size: len(items)}
	hp.initInt16Heap()
	return hp
}

// Len len
func (h *Int16Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Int16Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Int16Heap) Top() (int16, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Int16Heap) Pop() (int16, bool) {
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
func (h *Int16Heap) Remove(index int) (int16, bool) {
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
func (h *Int16Heap) Push(v int16) {
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

func (h *Int16Heap) initInt16Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Int16Heap) down(u int) {
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

func (h *Int16Heap) up(u int) {
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

// Uint64Heap int heap
type Uint64Heap struct {
	items []uint64
	size  int
}

// NewUint64Heap new int heap
func NewUint64Heap(items ...uint64) *Uint64Heap {
	var hp = &Uint64Heap{items: items, size: len(items)}
	hp.initUint64Heap()
	return hp
}

// Len len
func (h *Uint64Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Uint64Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Uint64Heap) Top() (uint64, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Uint64Heap) Pop() (uint64, bool) {
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
func (h *Uint64Heap) Remove(index int) (uint64, bool) {
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
func (h *Uint64Heap) Push(v uint64) {
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

func (h *Uint64Heap) initUint64Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Uint64Heap) down(u int) {
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

func (h *Uint64Heap) up(u int) {
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

// Uint32Heap int heap
type Uint32Heap struct {
	items []uint32
	size  int
}

// NewUint32Heap new int heap
func NewUint32Heap(items ...uint32) *Uint32Heap {
	var hp = &Uint32Heap{items: items, size: len(items)}
	hp.initUint32Heap()
	return hp
}

// Len len
func (h *Uint32Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Uint32Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Uint32Heap) Top() (uint32, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Uint32Heap) Pop() (uint32, bool) {
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
func (h *Uint32Heap) Remove(index int) (uint32, bool) {
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
func (h *Uint32Heap) Push(v uint32) {
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

func (h *Uint32Heap) initUint32Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Uint32Heap) down(u int) {
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

func (h *Uint32Heap) up(u int) {
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

// Uint16Heap int heap
type Uint16Heap struct {
	items []uint16
	size  int
}

// NewUint16Heap new int heap
func NewUint16Heap(items ...uint16) *Uint16Heap {
	var hp = &Uint16Heap{items: items, size: len(items)}
	hp.initUint16Heap()
	return hp
}

// Len len
func (h *Uint16Heap) Len() int {
	return h.size
}

// Empty empty
func (h *Uint16Heap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *Uint16Heap) Top() (uint16, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *Uint16Heap) Pop() (uint16, bool) {
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
func (h *Uint16Heap) Remove(index int) (uint16, bool) {
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
func (h *Uint16Heap) Push(v uint16) {
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

func (h *Uint16Heap) initUint16Heap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Uint16Heap) down(u int) {
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

func (h *Uint16Heap) up(u int) {
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

// UintHeap int heap
type UintHeap struct {
	items []uint
	size  int
}

// NewUintHeap new int heap
func NewUintHeap(items ...uint) *UintHeap {
	var hp = &UintHeap{items: items, size: len(items)}
	hp.initUintHeap()
	return hp
}

// Len len
func (h *UintHeap) Len() int {
	return h.size
}

// Empty empty
func (h *UintHeap) Empty() bool {
	return h.Len() == 0
}

// Top top
func (h *UintHeap) Top() (uint, bool) {
	if h.Empty() {
		return 0, false
	}
	return h.items[0], true
}

// Pop pop
func (h *UintHeap) Pop() (uint, bool) {
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
func (h *UintHeap) Remove(index int) (uint, bool) {
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
func (h *UintHeap) Push(v uint) {
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

func (h *UintHeap) initUintHeap() {
	for i := h.Len() >> 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *UintHeap) down(u int) {
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

func (h *UintHeap) up(u int) {
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
