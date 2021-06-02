package unionfindset

// UnionFindSet
type UnionFindSet interface {
	// Find 查找元素a的根节点
	Find(int) int
	// 合并两个集合
	// False: 说明已经在一个集合中，无需合并
	// True: 合并成功
	Union(a, b int) bool
}
