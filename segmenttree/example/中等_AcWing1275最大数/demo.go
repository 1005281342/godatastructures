package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		arr    = readIntArray(reader)
		m      = arr[0]
		p      = arr[1]
		merge  = func(a interface{}, b interface{}) interface{} {
			if ToInt(a) >= ToInt(b) {
				return a
			}
			return b
		}
		seg = New(make([]interface{}, m), merge)
		idx = 0
		cnt = 0
	)
	for i := 0; i < m; i++ {
		var (
			arrString = readStrArray(reader)
			query     = arrString[0]
		)
		switch query {
		case "A":
			var t = ConvT(Str2Int(arrString[1]), p)
			seg.Set(idx, t)
			idx++
			cnt++
		case "Q":
			var t = Str2Int(arrString[1])
			var v, _ = seg.Query(cnt-t, cnt-1)
			var tv = ToInt(v)
			lastAns = tv
			fmt.Println(tv)
		}
	}
}

func Str2Int(a string) int {
	var v, _ = strconv.Atoi(a)
	return v
}

var lastAns = 0

func ConvT(t int, p int) int {
	return (t + lastAns) % p
}

func ToInt(a interface{}) int {
	if a == nil {
		return 0
	}
	var v, ok = a.(int)
	if !ok {
		return 0
	}
	return v
}

var (
	// ErrIndexIllegal 索引非法
	ErrIndexIllegal = errors.New("index is illegal")
)

// Merger 用户自定义区间内操作逻辑
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
	// merger
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

func nums2string(x []int, sep string) string {
	var b strings.Builder
	for i := 0; i < len(x); i++ {
		b.WriteString(strconv.Itoa(x[i]))
		b.WriteString(sep)
	}
	return b.String()
}

// https://www.acwing.com/problem/content/1277/
