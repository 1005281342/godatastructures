package segmenttree

import (
	"errors"
)

var (
	ErrIndexIllegal = errors.New("index is illegal")
)

type Merger func(interface{}, interface{}) interface{}

// SegmentTree 线段树
type SegmentTree struct {
	data   []interface{}
	tree   []interface{}
	merger Merger
}

// New new SegmentTree
func New(array []interface{}, merger Merger) *SegmentTree {
	var seg = &SegmentTree{
		data:   array,
		tree:   make([]interface{}, 4*len(array)),
		merger: merger,
	}
	seg.buildSegmentTree(0, 0, seg.Size()-1)
	return seg
}

// leftChild 左子树index
func (s *SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

// rightChild 右子树index
func (s *SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

// buildSegmentTree 建立线段树
func (s *SegmentTree) buildSegmentTree(idx int, left int, right int) {
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
	s.buildSegmentTree(leftTreeIdx, left, mid)
	// 创建右子树的线段树
	s.buildSegmentTree(rightTreeIdx, mid+1, right)

	s.tree[idx] = s.merger(s.tree[leftTreeIdx], s.tree[rightTreeIdx])
}

// Size 线段树元素数
func (s *SegmentTree) Size() int {
	return len(s.data)
}

// Get 索引原数组值
func (s *SegmentTree) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= s.Size() {
		return nil, ErrIndexIllegal
	}
	return s.data[idx], nil
}

// Query 区间查询
func (s *SegmentTree) Query(queryLeft int, queryRight int) (interface{}, error) {
	if queryLeft < 0 || queryRight < 0 || queryLeft >= s.Size() || queryRight >= s.Size() {
		return nil, ErrIndexIllegal
	}
	return s.query(0, 0, s.Size()-1, queryLeft, queryRight), nil
}

func (s *SegmentTree) query(idx int, left int, right int, queryLeft int, queryRight int) interface{} {
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
	return s.merger(
		// 查询左部分
		s.query(leftIdx, left, mid, queryLeft, mid),
		// 查询右部分
		s.query(rightIdx, mid+1, right, mid+1, queryRight),
	)
}

// Set 更新元素值
func (s *SegmentTree) Set(index int, e interface{}) error {
	if index < 0 || index >= s.Size() {
		return ErrIndexIllegal
	}
	// 更新数组元素值
	s.data[index] = e
	// 更新tree元素值
	s.set(0, 0, s.Size()-1, index, e)
	return nil
}

func (s *SegmentTree) set(idx int, left int, right int, index int, e interface{}) {
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
	// merger
	s.tree[idx] = s.merger(s.tree[leftIdx], s.tree[rightIdx])
}
