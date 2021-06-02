package unionfindset

type DSU []int

// ConstructorDSU 创建容量为n的并查集
// n指的是并查集中点的个数
func ConstructorDSU(n int) DSU {
	var dsu = make(DSU, n)
	// 初始化时每个节点的根节点就是自身
	for i := 0; i < n; i++ {
		dsu[i] = i
	}
	return dsu
}

// Find 查找a的根节点，并进行路径压缩（注意使用路径压缩的话，独立集合的根结点不能初始化为一个相同值，比如-1）
// 如`dsu[i] = -1`，这样在合并时求解是否存在环的时候所找根节点都会为-1而造成误判
func (dsu DSU) Find(a int) int {
	if dsu[a] != a {
		// 某个节点的根节点不是自身时
		// 路径压缩
		dsu[a] = dsu.Find(dsu[a])
	}
	// 获取到根节点
	return dsu[a]
}

// Union 合并两个集合
// False: 说明已经在一个集合中，无需合并
// True: 合并成功
func (dsu DSU) Union(a, b int) bool {
	var (
		aRoot = dsu.Find(a)
		bRoot = dsu.Find(b)
	)
	// 如果根节点相同说明已经在同一个集合中了
	if aRoot == bRoot {
		return false
	}
	// 将集合a加入到集合b中
	dsu[aRoot] = bRoot
	return true
}
