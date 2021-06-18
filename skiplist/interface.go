package skiplist

type ISkipList interface {
	// Search 查找是否存在target
	Search(target int) bool
	// Add 添加元素
	Add(num int)
	// Erase 移除元素
	Erase(num int) bool
}
