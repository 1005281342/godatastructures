package main

type MapSum struct {
	value int
	nodes map[byte]*MapSum
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{nodes: map[byte]*MapSum{}}
}

func (this *MapSum) Insert(key string, val int) {
	var curNode = this
	for i := 0; i < len(key); i++ {
		if _, has := curNode.nodes[key[i]]; !has {
			var t = Constructor()
			curNode.nodes[key[i]] = &t
		}
		curNode = curNode.nodes[key[i]]
	}
	curNode.value = val
}

func (this *MapSum) Sum(prefix string) int {
	// 1. find prefix
	var (
		node *MapSum
		has  bool
	)
	if node, has = this.PrefixNode(prefix); !has {
		return 0
	}
	// 2. cnt value
	return this.Count(node)
}

// Count 扫描节点累计值
func (this *MapSum) Count(node *MapSum) int {
	var cnt = node.value
	for _, subNode := range node.nodes {
		cnt += this.Count(subNode)
	}
	return cnt
}

func (this *MapSum) PrefixNode(prefix string) (*MapSum, bool) {
	var curNode = this
	for i := 0; i < len(prefix); i++ {
		if _, has := curNode.nodes[prefix[i]]; !has {
			return nil, false
		}
		curNode = curNode.nodes[prefix[i]]
	}
	return curNode, true
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
