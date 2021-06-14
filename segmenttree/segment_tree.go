package segmenttree

import (
	"errors"
)

type Merge func(interface{}, interface{}) interface{}

// SegmentTree 线段树
type SegmentTree struct {
	data   []interface{}
	tree   []interface{}
	merger Merge
}

// New new SegmentTree
func New(array []interface{}, merger Merge) *SegmentTree {
	var seg = &SegmentTree{
		data:   array,
		tree:   make([]interface{}, 4*len(array)),
		merger: merger,
	}
	seg.buildSegmentTree(0, 0, seg.Size()-1)
	return seg
}

func (s *SegmentTree) leftChild(idx int) int {
	return (idx << 1) + 1
}

func (s *SegmentTree) rightChild(idx int) int {
	return (idx << 1) + 2
}

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

func (s *SegmentTree) Size() int {
	return len(s.data)
}

func (s *SegmentTree) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= s.Size() {
		return nil, errors.New("index is illegal")
	}
	return s.data[idx], nil
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
