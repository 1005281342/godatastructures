package trie

// HashTrie Hash前缀树
type HashTrie struct {
	isWord   bool
	children map[byte]*HashTrie
}

var _ Trie = (*HashTrie)(nil)

// NewHashTrie new hashTrie
func NewHashTrie() *HashTrie {
	return &HashTrie{children: make(map[byte]*HashTrie)}
}

// Insert 往前缀树中添加一个元素word
func (h *HashTrie) Insert(word string) {
	var curNode = h
	for i := 0; i < len(word); i++ {
		if _, has := curNode.children[word[i]]; !has {
			curNode.children[word[i]] = NewHashTrie()
		}
		curNode = curNode.children[word[i]]
	}
	curNode.isWord = true
}

// Search 查找前缀树中是否元素word
func (h *HashTrie) Search(word string) bool {
	var _, has = h.search(word)
	return has
}

func (h *HashTrie) search(word string) (*HashTrie, bool) {
	var curNode = h
	for i := 0; i < len(word); i++ {
		if curNode.children[word[i]] == nil {
			return nil, false
		}
		curNode = curNode.children[word[i]]
	}
	return curNode, curNode.isWord
}

// HasPrefix 查询前缀树中是否存在前缀prefix
func (h *HashTrie) HasPrefix(prefix string) bool {
	var curNode = h
	for i := 0; i < len(prefix); i++ {
		if curNode.children[prefix[i]] == nil {
			return false
		}
		curNode = curNode.children[prefix[i]]
	}
	return true
}
