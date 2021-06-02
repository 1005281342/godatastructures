package main

import "github.com/1005281342/godatastructures/unionfindset"

func findRedundantDirectedConnection(edges [][]int) []int {
	var (
		dsu      = unionfindset.ConstructorDSU(1010)
		entryMap = make(map[int][]int)
		flag     bool
		nodeY    int
	)
	// 统计入度
	for i := 0; i < len(edges); i++ {
		entryMap[edges[i][1]] = append(entryMap[edges[i][1]], edges[i][0])
		if len(entryMap[edges[i][1]]) == 2 {
			flag = true
			nodeY = edges[i][1]
			break
		}
	}

	if flag {
		// 存在入度不小于2的点
		for i := 0; i < len(edges); i++ {
			if edges[i][0] == entryMap[nodeY][1] && edges[i][1] == nodeY {
				continue
			}
			dsu.Union(edges[i][0], edges[i][1])
		}

		if !dsu.Union(entryMap[nodeY][1], nodeY) {
			return []int{entryMap[nodeY][1], nodeY}
		}
		return []int{entryMap[nodeY][0], nodeY}
	}

	// 不存在
	for i := 0; i < len(edges); i++ {
		if !dsu.Union(edges[i][0], edges[i][1]) {
			// 合并失败，说明构成了环
			return []int{edges[i][0], edges[i][1]}
		}
	}
	return []int{}
}
