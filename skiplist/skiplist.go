package skiplist

// https://leetcode-cn.com/problems/design-skiplist/

import (
	"math"
	"math/rand"
	"sync"
)

const (
	headValue = math.MinInt32
)

// SkipList跳表
type SkipList struct {
	// 虚拟头节点
	head *node
	// 层级
	level int
	// 链表长度
	length int
	// 锁
	lock sync.RWMutex
}

// 节点
type node struct {
	// 节点值
	val int
	// 右节点
	right *node
	// 下一级
	down *node
}

func NewSkipList() *SkipList {
	var sp = Constructor()
	return &sp
}

func Constructor() SkipList {
	return SkipList{
		// 虚拟头节点
		head: &node{
			val: headValue,
		}, //头节点 设置为一个极小的值
		level:  1, //跳表层数（包括原链表）
		length: 1, //原链包的个数（包括虚拟头节点）
	}
}

// Search 查找是否存在target
// 由于设置了虚拟节点，那么意味着它存在的话，它必然是某个节点的右节点
func (s *SkipList) Search(target int) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.getBefore(target, s.head) != nil
}

// Add 添加元素
func (s *SkipList) Add(num int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var (
		n     = s.head
		i     int
		nodes = make([]*node, s.level+1)
	)
	for n != nil {
		for n.right != nil && n.right.val < num {
			n = n.right
		}
		nodes[i] = n
		i++
		n = n.down
	}
	i--                       // 最后一个i是没有值的要去掉
	var beforeNode = nodes[i] // 原链表的的前值
	newNode := &node{
		val:   num,
		right: beforeNode.right,
	}
	beforeNode.right = newNode
	s.length++
	// 50%概率建立索引
	for {
		if rand.Intn(2) == 0 || s.level > s.length>>6+1 {
			break
		}
		if i > 0 {
			i--
			newNode = &node{
				val:   num,
				right: nodes[i].right,
				down:  newNode,
			}
			nodes[i].right = newNode
		} else {
			// 超过最大层数建索引
			newNode = &node{
				val:  num,
				down: newNode,
			}
			s.head = &node{
				val:   headValue,
				right: newNode,
				down:  s.head,
			}
			s.level++
		}
	}
}

// Erase 移除元素
func (s *SkipList) Erase(num int) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	var (
		target = num
		before = s.getBefore(target, s.head)
	)
	if before == nil {
		return false
	}
	for {
		if before == nil {
			break
		}
		before.right = before.right.right
		before = before.down
		before = s.getBefore(num, before)
	}
	s.length--
	return true
}

// 获取目标值的前一个节点
func (s *SkipList) getBefore(target int, from *node) *node {
	var n = from
	for n != nil {

		// 如果n存在右节点
		for n.right != nil && n.right.val < target {
			n = n.right
		}
		// 找到了
		if n.right != nil && n.right.val == target {
			return n
		}
		// 没找到到下一级找
		n = n.down
	}
	// 没找到
	return nil
}
