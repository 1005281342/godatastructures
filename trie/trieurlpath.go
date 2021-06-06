package trie

import (
	"strings"
)

const slash = "/"

// NodeURLPath trie node
type NodeURLPath struct {
	// isPath
	isPath bool
	// NodeURLPaths 某个节点下挂载的子节点
	urlPaths map[string]*NodeURLPath
}

// NewURLPathTrie trie构造函数
func NewURLPathTrie() *NodeURLPath {
	// init
	return &NodeURLPath{urlPaths: map[string]*NodeURLPath{}}
}

// Insert 插入一个节点  /aa/bb/cc
func (n *NodeURLPath) Insert(path string) {
	var (
		subPaths = strings.Split(path, slash)
		curNode  = n
	)
	for i := 0; i < len(subPaths); i++ {
		if _, has := curNode.urlPaths[subPaths[i]]; !has {
			curNode.urlPaths[subPaths[i]] = &NodeURLPath{urlPaths: map[string]*NodeURLPath{}}
		}
		curNode = curNode.urlPaths[subPaths[i]]
	}
	// 路径的最后一个节点
	curNode.isPath = true
}

// Search 查找
func (n *NodeURLPath) Search(path string) bool {
	var (
		subPaths = strings.Split(path, slash)
		curNode  = n
	)
	for i := 0; i < len(subPaths); i++ {
		if _, has := curNode.urlPaths[subPaths[i]]; !has {
			return false
		}
		curNode = curNode.urlPaths[subPaths[i]]
	}
	return curNode.isPath
}
