package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrIndexIllegal = errors.New("index is illegal")
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		arr    = readIntArray(reader)
		q      = arr[1]
		nums   = readInt64Array(reader)
		merge  = func(a int64, b int64) int64 {
			return a + b
		}
		seg = NewInt64SegmentTree(nums, merge)
	)
	seg.WithInt64Optimization(func(left int, right int, value int64) int64 {
		return value * int64(right-left+1)
	})
	for i := 0; i < q; i++ {
		var (
			arrString = readStrArray(reader)
			left      = str2Int(arrString[1]) - 1
			right     = str2Int(arrString[2]) - 1
		)
		switch arrString[0] {
		case "Q":
			var v interface{}
			v, _ = seg.QueryLazy(left, right)
			fmt.Println(v)
		case "C":
			var value = int64(str2Int(arrString[3]))
			seg.AddValueLazy(left, right, value)
		}
	}
}

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

func str2Int(s string) int {
	var a, _ = strconv.Atoi(s)
	return a
}

func readLine(reader *bufio.Reader) string {
	var line, _ = reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}

func readInt(reader *bufio.Reader) int {
	var num, _ = strconv.Atoi(readLine(reader))
	return num
}

func readStrArray(reader *bufio.Reader) []string {
	var line = readLine(reader)
	return strings.Split(line, " ")
}

func readIntArray(reader *bufio.Reader) []int {
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

func readInt64Array(reader *bufio.Reader) []int64 {
	var line = readLine(reader)
	var strList = strings.Split(line, " ")
	var nums = make([]int64, 0)
	var err error
	var v int
	for i := 0; i < len(strList); i++ {
		if v, err = strconv.Atoi(strList[i]); err != nil {
			continue
		}
		nums = append(nums, int64(v))
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

// Accepted
