package main

import (
	"github.com/1005281342/godatastructures/unionfindset"
)

func findRedundantConnection(edges [][]int) []int {
	var dsu = unionfindset.ConstructorDSU(2*len(edges) + 10)
	for i := 0; i < len(edges); i++ {
		if !dsu.Union(edges[i][0], edges[i][1]) {
			// 合并失败，说明构成了环
			if edges[i][0] > edges[i][1] {
				return []int{edges[i][1], edges[i][0]}
			}
			return []int{edges[i][0], edges[i][1]}
		}
	}
	return []int{}
}
