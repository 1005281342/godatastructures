package main

type WordDictionary struct {
	isWord bool
	nodes  map[byte]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{nodes: map[byte]*WordDictionary{}}
}

func (this *WordDictionary) AddWord(word string) {
	var curNode = this
	for i := 0; i < len(word); i++ {
		if _, has := curNode.nodes[word[i]]; !has {
			var t = Constructor()
			curNode.nodes[word[i]] = &t
		}
		curNode = curNode.nodes[word[i]]
	}
	curNode.isWord = true
	return
}

func (this *WordDictionary) Search(word string) bool {
	return this.Match(this, word, 0)
}

func (this *WordDictionary) Match(node *WordDictionary, word string, index int) bool {
	// 所有字符都扫描了一遍
	if index == len(word) {
		return node.isWord
	}
	if word[index:index+1] != "." {
		// 字符匹配
		if _, has := node.nodes[word[index]]; !has {
			return false
		}
		return this.Match(node.nodes[word[index]], word, index+1)
	}
	for k := range node.nodes {
		if this.Match(node.nodes[k], word, index+1) {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
