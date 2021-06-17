package disjointsets

type DSUCount struct {
	dsu   DSU
	count []int
}

// ConstructorDSUCount ...
func ConstructorDSUCount(n int) *DSUCount {
	var (
		count = make([]int, n)
		dsu   = ConstructorDSU(n)
	)
	for i := 0; i < n; i++ {
		count[i] = 1
	}
	return &DSUCount{dsu: dsu, count: count}
}

// Find ...
func (d *DSUCount) Find(a int) int {
	return d.dsu.Find(a)
}

// Union ...
func (d *DSUCount) Union(a, b int) bool {
	// 记录合并前a集合的个数
	var cntA = d.count[d.Find(a)]
	if !d.dsu.Union(a, b) {
		return false
	}
	d.count[d.Find(b)] += cntA
	return true
}

// Count ...
func (d *DSUCount) Count(a int) int {
	return d.count[d.Find(a)]
}
