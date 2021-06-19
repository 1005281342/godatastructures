package main

type Trie struct {
	isWord bool
	nodes  map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{nodes: map[byte]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var curNode = this
	for i := 0; i < len(word); i++ {
		if _, has := curNode.nodes[word[i]]; !has {
			var t = Constructor()
			curNode.nodes[word[i]] = &t
		}
		curNode = curNode.nodes[word[i]]
	}
	curNode.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var curNode = this
	for i := 0; i < len(word); i++ {
		if _, has := curNode.nodes[word[i]]; !has {
			return false
		}
		curNode = curNode.nodes[word[i]]
	}
	return curNode.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var curNode = this
	for i := 0; i < len(prefix); i++ {
		if _, has := curNode.nodes[prefix[i]]; !has {
			return false
		}
		curNode = curNode.nodes[prefix[i]]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
