package segmenttree

// Int64Merger 用户自定义区间内操作逻辑
type Int64Merger func(int64, int64) int64

type Int64Optimization func(int, int, int64) int64

// Int64SegmentTree 线段树
type Int64SegmentTree struct {
	data         []int64
	tree         []int64
	lazy         []int64
	merge        Int64Merger
	optimization Int64Optimization
}

// NewInt64SegmentTree new Int64SegmentTree
func NewInt64SegmentTree(array []int64, merge Int64Merger) *Int64SegmentTree {
	var seg = &Int64SegmentTree{
		data:  array,
		tree:  make([]int64, 4*len(array)),
		lazy:  make([]int64, 4*len(array)),
		merge: merge,
	}
	seg.buildInt64SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithInt64Optimization 指定Int64Optimization来自定义区间[left, right]更新value的计算操作
func (s *Int64SegmentTree) WithInt64Optimization(f Int64Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Int64SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Int64SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildInt64SegmentTree 建立线段树
func (s *Int64SegmentTree) buildInt64SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildInt64SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildInt64SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Int64SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Int64SegmentTree) Get(idx int) (int64, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Int64SegmentTree) Query(queryLeft int, queryRight int) (int64, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Int64SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) int64 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Int64SegmentTree) QueryLazy(queryLeft int, queryRight int) (int64, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Int64SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) int64 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Int64SegmentTree) Set(index int, e int64) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Int64SegmentTree) set(idx int, left int, right int, index int, e int64) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Int64SegmentTree) AddValueLazy(addLeft int, addRight int, value int64) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Int64SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value int64) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Int64SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// Float32Merger 用户自定义区间内操作逻辑
type Float32Merger func(float32, float32) float32

type Float32Optimization func(int, int, float32) float32

// Float32SegmentTree 线段树
type Float32SegmentTree struct {
	data         []float32
	tree         []float32
	lazy         []float32
	merge        Float32Merger
	optimization Float32Optimization
}

// NewFloat32SegmentTree new Float32SegmentTree
func NewFloat32SegmentTree(array []float32, merge Float32Merger) *Float32SegmentTree {
	var seg = &Float32SegmentTree{
		data:  array,
		tree:  make([]float32, 4*len(array)),
		lazy:  make([]float32, 4*len(array)),
		merge: merge,
	}
	seg.buildFloat32SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithFloat32Optimization 指定Float32Optimization来自定义区间[left, right]更新value的计算操作
func (s *Float32SegmentTree) WithFloat32Optimization(f Float32Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Float32SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Float32SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildFloat32SegmentTree 建立线段树
func (s *Float32SegmentTree) buildFloat32SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildFloat32SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildFloat32SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Float32SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Float32SegmentTree) Get(idx int) (float32, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Float32SegmentTree) Query(queryLeft int, queryRight int) (float32, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Float32SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) float32 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Float32SegmentTree) QueryLazy(queryLeft int, queryRight int) (float32, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Float32SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) float32 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Float32SegmentTree) Set(index int, e float32) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Float32SegmentTree) set(idx int, left int, right int, index int, e float32) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Float32SegmentTree) AddValueLazy(addLeft int, addRight int, value float32) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Float32SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value float32) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Float32SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// Float64Merger 用户自定义区间内操作逻辑
type Float64Merger func(float64, float64) float64

type Float64Optimization func(int, int, float64) float64

// Float64SegmentTree 线段树
type Float64SegmentTree struct {
	data         []float64
	tree         []float64
	lazy         []float64
	merge        Float64Merger
	optimization Float64Optimization
}

// NewFloat64SegmentTree new Float64SegmentTree
func NewFloat64SegmentTree(array []float64, merge Float64Merger) *Float64SegmentTree {
	var seg = &Float64SegmentTree{
		data:  array,
		tree:  make([]float64, 4*len(array)),
		lazy:  make([]float64, 4*len(array)),
		merge: merge,
	}
	seg.buildFloat64SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithFloat64Optimization 指定Float64Optimization来自定义区间[left, right]更新value的计算操作
func (s *Float64SegmentTree) WithFloat64Optimization(f Float64Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Float64SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Float64SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildFloat64SegmentTree 建立线段树
func (s *Float64SegmentTree) buildFloat64SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildFloat64SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildFloat64SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Float64SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Float64SegmentTree) Get(idx int) (float64, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Float64SegmentTree) Query(queryLeft int, queryRight int) (float64, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Float64SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) float64 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Float64SegmentTree) QueryLazy(queryLeft int, queryRight int) (float64, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Float64SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) float64 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Float64SegmentTree) Set(index int, e float64) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Float64SegmentTree) set(idx int, left int, right int, index int, e float64) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Float64SegmentTree) AddValueLazy(addLeft int, addRight int, value float64) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Float64SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value float64) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Float64SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// Int32Merger 用户自定义区间内操作逻辑
type Int32Merger func(int32, int32) int32

type Int32Optimization func(int, int, int32) int32

// Int32SegmentTree 线段树
type Int32SegmentTree struct {
	data         []int32
	tree         []int32
	lazy         []int32
	merge        Int32Merger
	optimization Int32Optimization
}

// NewInt32SegmentTree new Int32SegmentTree
func NewInt32SegmentTree(array []int32, merge Int32Merger) *Int32SegmentTree {
	var seg = &Int32SegmentTree{
		data:  array,
		tree:  make([]int32, 4*len(array)),
		lazy:  make([]int32, 4*len(array)),
		merge: merge,
	}
	seg.buildInt32SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithInt32Optimization 指定Int32Optimization来自定义区间[left, right]更新value的计算操作
func (s *Int32SegmentTree) WithInt32Optimization(f Int32Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Int32SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Int32SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildInt32SegmentTree 建立线段树
func (s *Int32SegmentTree) buildInt32SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildInt32SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildInt32SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Int32SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Int32SegmentTree) Get(idx int) (int32, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Int32SegmentTree) Query(queryLeft int, queryRight int) (int32, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Int32SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) int32 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Int32SegmentTree) QueryLazy(queryLeft int, queryRight int) (int32, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Int32SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) int32 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Int32SegmentTree) Set(index int, e int32) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Int32SegmentTree) set(idx int, left int, right int, index int, e int32) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Int32SegmentTree) AddValueLazy(addLeft int, addRight int, value int32) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Int32SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value int32) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Int32SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// Int16Merger 用户自定义区间内操作逻辑
type Int16Merger func(int16, int16) int16

type Int16Optimization func(int, int, int16) int16

// Int16SegmentTree 线段树
type Int16SegmentTree struct {
	data         []int16
	tree         []int16
	lazy         []int16
	merge        Int16Merger
	optimization Int16Optimization
}

// NewInt16SegmentTree new Int16SegmentTree
func NewInt16SegmentTree(array []int16, merge Int16Merger) *Int16SegmentTree {
	var seg = &Int16SegmentTree{
		data:  array,
		tree:  make([]int16, 4*len(array)),
		lazy:  make([]int16, 4*len(array)),
		merge: merge,
	}
	seg.buildInt16SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithInt16Optimization 指定Int16Optimization来自定义区间[left, right]更新value的计算操作
func (s *Int16SegmentTree) WithInt16Optimization(f Int16Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Int16SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Int16SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildInt16SegmentTree 建立线段树
func (s *Int16SegmentTree) buildInt16SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildInt16SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildInt16SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Int16SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Int16SegmentTree) Get(idx int) (int16, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Int16SegmentTree) Query(queryLeft int, queryRight int) (int16, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Int16SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) int16 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Int16SegmentTree) QueryLazy(queryLeft int, queryRight int) (int16, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Int16SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) int16 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Int16SegmentTree) Set(index int, e int16) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Int16SegmentTree) set(idx int, left int, right int, index int, e int16) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Int16SegmentTree) AddValueLazy(addLeft int, addRight int, value int16) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Int16SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value int16) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Int16SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// IntMerger 用户自定义区间内操作逻辑
type IntMerger func(int, int) int

type IntOptimization func(int, int, int) int

// IntSegmentTree 线段树
type IntSegmentTree struct {
	data         []int
	tree         []int
	lazy         []int
	merge        IntMerger
	optimization IntOptimization
}

// NewIntSegmentTree new IntSegmentTree
func NewIntSegmentTree(array []int, merge IntMerger) *IntSegmentTree {
	var seg = &IntSegmentTree{
		data:  array,
		tree:  make([]int, 4*len(array)),
		lazy:  make([]int, 4*len(array)),
		merge: merge,
	}
	seg.buildIntSegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithIntOptimization 指定IntOptimization来自定义区间[left, right]更新value的计算操作
func (s *IntSegmentTree) WithIntOptimization(f IntOptimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *IntSegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *IntSegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildIntSegmentTree 建立线段树
func (s *IntSegmentTree) buildIntSegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildIntSegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildIntSegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *IntSegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *IntSegmentTree) Get(idx int) (int, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *IntSegmentTree) Query(queryLeft int, queryRight int) (int, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *IntSegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) int {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *IntSegmentTree) QueryLazy(queryLeft int, queryRight int) (int, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *IntSegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) int {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *IntSegmentTree) Set(index int, e int) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *IntSegmentTree) set(idx int, left int, right int, index int, e int) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *IntSegmentTree) AddValueLazy(addLeft int, addRight int, value int) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *IntSegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value int) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *IntSegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// Uint64Merger 用户自定义区间内操作逻辑
type Uint64Merger func(uint64, uint64) uint64

type Uint64Optimization func(int, int, uint64) uint64

// Uint64SegmentTree 线段树
type Uint64SegmentTree struct {
	data         []uint64
	tree         []uint64
	lazy         []uint64
	merge        Uint64Merger
	optimization Uint64Optimization
}

// NewUint64SegmentTree new Uint64SegmentTree
func NewUint64SegmentTree(array []uint64, merge Uint64Merger) *Uint64SegmentTree {
	var seg = &Uint64SegmentTree{
		data:  array,
		tree:  make([]uint64, 4*len(array)),
		lazy:  make([]uint64, 4*len(array)),
		merge: merge,
	}
	seg.buildUint64SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithUint64Optimization 指定Uint64Optimization来自定义区间[left, right]更新value的计算操作
func (s *Uint64SegmentTree) WithUint64Optimization(f Uint64Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Uint64SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Uint64SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildUint64SegmentTree 建立线段树
func (s *Uint64SegmentTree) buildUint64SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildUint64SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildUint64SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Uint64SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Uint64SegmentTree) Get(idx int) (uint64, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Uint64SegmentTree) Query(queryLeft int, queryRight int) (uint64, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Uint64SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) uint64 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Uint64SegmentTree) QueryLazy(queryLeft int, queryRight int) (uint64, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Uint64SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) uint64 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Uint64SegmentTree) Set(index int, e uint64) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Uint64SegmentTree) set(idx int, left int, right int, index int, e uint64) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Uint64SegmentTree) AddValueLazy(addLeft int, addRight int, value uint64) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Uint64SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value uint64) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Uint64SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// Uint32Merger 用户自定义区间内操作逻辑
type Uint32Merger func(uint32, uint32) uint32

type Uint32Optimization func(int, int, uint32) uint32

// Uint32SegmentTree 线段树
type Uint32SegmentTree struct {
	data         []uint32
	tree         []uint32
	lazy         []uint32
	merge        Uint32Merger
	optimization Uint32Optimization
}

// NewUint32SegmentTree new Uint32SegmentTree
func NewUint32SegmentTree(array []uint32, merge Uint32Merger) *Uint32SegmentTree {
	var seg = &Uint32SegmentTree{
		data:  array,
		tree:  make([]uint32, 4*len(array)),
		lazy:  make([]uint32, 4*len(array)),
		merge: merge,
	}
	seg.buildUint32SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithUint32Optimization 指定Uint32Optimization来自定义区间[left, right]更新value的计算操作
func (s *Uint32SegmentTree) WithUint32Optimization(f Uint32Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Uint32SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Uint32SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildUint32SegmentTree 建立线段树
func (s *Uint32SegmentTree) buildUint32SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildUint32SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildUint32SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Uint32SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Uint32SegmentTree) Get(idx int) (uint32, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Uint32SegmentTree) Query(queryLeft int, queryRight int) (uint32, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Uint32SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) uint32 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Uint32SegmentTree) QueryLazy(queryLeft int, queryRight int) (uint32, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Uint32SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) uint32 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Uint32SegmentTree) Set(index int, e uint32) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Uint32SegmentTree) set(idx int, left int, right int, index int, e uint32) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Uint32SegmentTree) AddValueLazy(addLeft int, addRight int, value uint32) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Uint32SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value uint32) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Uint32SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// Uint16Merger 用户自定义区间内操作逻辑
type Uint16Merger func(uint16, uint16) uint16

type Uint16Optimization func(int, int, uint16) uint16

// Uint16SegmentTree 线段树
type Uint16SegmentTree struct {
	data         []uint16
	tree         []uint16
	lazy         []uint16
	merge        Uint16Merger
	optimization Uint16Optimization
}

// NewUint16SegmentTree new Uint16SegmentTree
func NewUint16SegmentTree(array []uint16, merge Uint16Merger) *Uint16SegmentTree {
	var seg = &Uint16SegmentTree{
		data:  array,
		tree:  make([]uint16, 4*len(array)),
		lazy:  make([]uint16, 4*len(array)),
		merge: merge,
	}
	seg.buildUint16SegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithUint16Optimization 指定Uint16Optimization来自定义区间[left, right]更新value的计算操作
func (s *Uint16SegmentTree) WithUint16Optimization(f Uint16Optimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *Uint16SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *Uint16SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildUint16SegmentTree 建立线段树
func (s *Uint16SegmentTree) buildUint16SegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildUint16SegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildUint16SegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *Uint16SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *Uint16SegmentTree) Get(idx int) (uint16, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *Uint16SegmentTree) Query(queryLeft int, queryRight int) (uint16, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Uint16SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) uint16 {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *Uint16SegmentTree) QueryLazy(queryLeft int, queryRight int) (uint16, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *Uint16SegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) uint16 {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *Uint16SegmentTree) Set(index int, e uint16) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *Uint16SegmentTree) set(idx int, left int, right int, index int, e uint16) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *Uint16SegmentTree) AddValueLazy(addLeft int, addRight int, value uint16) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *Uint16SegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value uint16) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *Uint16SegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// UintMerger 用户自定义区间内操作逻辑
type UintMerger func(uint, uint) uint

type UintOptimization func(int, int, uint) uint

// UintSegmentTree 线段树
type UintSegmentTree struct {
	data         []uint
	tree         []uint
	lazy         []uint
	merge        UintMerger
	optimization UintOptimization
}

// NewUintSegmentTree new UintSegmentTree
func NewUintSegmentTree(array []uint, merge UintMerger) *UintSegmentTree {
	var seg = &UintSegmentTree{
		data:  array,
		tree:  make([]uint, 4*len(array)),
		lazy:  make([]uint, 4*len(array)),
		merge: merge,
	}
	seg.buildUintSegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithUintOptimization 指定UintOptimization来自定义区间[left, right]更新value的计算操作
func (s *UintSegmentTree) WithUintOptimization(f UintOptimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *UintSegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *UintSegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildUintSegmentTree 建立线段树
func (s *UintSegmentTree) buildUintSegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildUintSegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildUintSegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *UintSegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *UintSegmentTree) Get(idx int) (uint, error) {
	if idx < 0 || idx >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *UintSegmentTree) Query(queryLeft int, queryRight int) (uint, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *UintSegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) uint {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *UintSegmentTree) QueryLazy(queryLeft int, queryRight int) (uint, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return 0, ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *UintSegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) uint {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return 0
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *UintSegmentTree) Set(index int, e uint) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *UintSegmentTree) set(idx int, left int, right int, index int, e uint) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *UintSegmentTree) AddValueLazy(addLeft int, addRight int, value uint) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *UintSegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value uint) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *UintSegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != 0 {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = 0
	}
}

// StringMerger 用户自定义区间内操作逻辑
type StringMerger func(string, string) string

type StringOptimization func(int, int, string) string

// StringSegmentTree 线段树
type StringSegmentTree struct {
	data         []string
	tree         []string
	lazy         []string
	merge        StringMerger
	optimization StringOptimization
}

// NewStringSegmentTree new StringSegmentTree
func NewStringSegmentTree(array []string, merge StringMerger) *StringSegmentTree {
	var seg = &StringSegmentTree{
		data:  array,
		tree:  make([]string, 4*len(array)),
		lazy:  make([]string, 4*len(array)),
		merge: merge,
	}
	seg.buildStringSegmentTree(0, 0, seg.Size()-1)
	return seg
}

// WithStringOptimization 指定StringOptimization来自定义区间[left, right]更新value的计算操作
func (s *StringSegmentTree) WithStringOptimization(f StringOptimization) {
	s.optimization = f
}

// leftChild 左子树index
func (s *StringSegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *StringSegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildStringSegmentTree 建立线段树
func (s *StringSegmentTree) buildStringSegmentTree(idx int, left int, right int) {
	if left == right {
		s.tree[idx] = s.data[left]
		return
	}
	var (
		leftTreeIdx  = s.leftChild(idx)
		rightTreeIdx = s.rightChild(idx)
		mid          = left + (right-left)/2
	)
	// 创建左子树的线段树
	s.buildStringSegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildStringSegmentTree(rightTreeIdx, mid+1, right)
	// merge
	s.tree[idx] = s.merge(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *StringSegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *StringSegmentTree) Get(idx int) (string, error) {
	if idx < 0 || idx >= s.Size() {
		return "", ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *StringSegmentTree) Query(queryLeft int, queryRight int) (string, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return "", ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *StringSegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) string {
	if left == queryLeft && right == queryRight {
		// 命中所查找区间
		return s.tree[idx]
	}

	var (
		mid = left + (right-left)/2
		// 计算左右子节点下标
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.query(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.query(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// QueryLazy 懒惰查询
func (s *StringSegmentTree) QueryLazy(queryLeft int, queryRight int) (string, error) {
	if queryLeft > queryRight {
		queryRight, queryLeft = queryLeft, queryRight
	}
	if queryLeft < 0 || queryRight >= s.Size() {
		return "", ErrIndexIllegal
	}
	return s.queryLazy(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *StringSegmentTree) queryLazy(idx int, left int, right int, queryLeft int, queryRight int) string {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > queryRight || right < queryLeft {
		return ""
	}
	if queryLeft <= left && right <= queryRight {
		// 在所查找区间范围内
		return s.tree[idx]
	}
	if queryLeft >= mid+1 {
		// 所查询区间在右半部分
		return s.queryLazy(rightIdx, mid+1, right, queryLeft, queryRight)
	}
	if queryRight <= mid {
		// 所查询区间在左半部分
		return s.queryLazy(leftIdx, left, mid, queryLeft, queryRight)
	}
	// 所查询区间分布在左右区间
	return s.merge(
		// 查询左部分
		s.queryLazy(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.queryLazy(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *StringSegmentTree) Set(index int, e string) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *StringSegmentTree) set(idx int, left int, right int, index int, e string) {
	if left == right {
		// 命中节点，更新元素值
		s.tree[idx] = e
		return
	}

	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
		mid      = left + (right-left)/2
	)
	if index <= mid {
		// idx在左边
		s.set(leftIdx, left, mid, index, e)
	} else {
		// idx在右边
		s.set(rightIdx, mid+1, right, index, e)
	}
	// merge
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

// AddValueLazy
// 给[addLeft....addRight]位置的值都加上value
// 注意这里的更新值是在原来值的基础上增加或者减少，而不是把这个区间内的值都赋值为 x，区间更新和单点更新不同
// 这里的区间更新关注的是变化，单点更新关注的是定值
// 当然区间更新也可以都更新成定值，如果只区间更新成定值，那么 lazy 更新策略需要变化，merge 策略也需要变化，这里暂不详细讨论
func (s *StringSegmentTree) AddValueLazy(addLeft int, addRight int, value string) error {
	if addLeft > addRight {
		addRight, addLeft = addLeft, addRight
	}
	if addLeft < 0 || addRight >= s.Size() {
		return ErrIndexIllegal
	}
	s.addValueLazy(0, 0, s.Size()-1, addLeft, addRight, value)
	return nil
}

func (s *StringSegmentTree) addValueLazy(idx int, left int, right int, addLeft int, addRight int, value string) {
	// 处理懒惰更新
	s.pushDown(idx, left, right)

	var (
		mid      = left + (right-left)>>1
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	if left > right || left > addRight || right < addLeft {
		return
	}
	if addLeft <= left && right <= addRight {
		// 正好在一个区间内区间
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, value))
		} else {
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], value)
			}
		}
		if left != right {
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], value)
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], value)
		}
		return
	}
	// 需要分别更新左右区间
	s.addValueLazy(leftIdx, left, mid, addLeft, addRight, value)
	s.addValueLazy(rightIdx, mid+1, right, addLeft, addRight, value)
	s.tree[idx] = s.merge(s.tree[leftIdx], s.tree[rightIdx])
}

func (s *StringSegmentTree) pushDown(idx int, left int, right int) {
	var (
		leftIdx  = s.leftChild(idx)
		rightIdx = s.rightChild(idx)
	)
	// 处理懒惰更新
	if s.lazy[idx] != "" {
		if s.optimization != nil {
			s.tree[idx] = s.merge(s.tree[idx], s.optimization(left, right, s.lazy[idx]))
		} else {
			// 懒惰更新根节点 如果是区间和则等同于：s.tree[idx] += (right-left+1)*s.lazy[idx]
			for i := 0; i < right-left+1; i++ {
				s.tree[idx] = s.merge(s.tree[idx], s.lazy[idx])
			}
		}
		if left != right {
			// 懒惰更新子节点 如果是区间求和则等同于：s.lazy[leftIdx] += s.lazy[idx]
			s.lazy[leftIdx] = s.merge(s.lazy[leftIdx], s.lazy[idx])
			s.lazy[rightIdx] = s.merge(s.lazy[rightIdx], s.lazy[idx])
		}
		// 消除懒惰更新标志
		s.lazy[idx] = ""
	}
}
