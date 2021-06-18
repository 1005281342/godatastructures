package trie

// ArrayTrie 数组前缀树
type ArrayTrie struct {
	isWord   bool
	children [26]*ArrayTrie
}

var _ Trie = (*ArrayTrie)(nil)

// NewArrayTrie new arrayTrie
func NewArrayTrie() *ArrayTrie {
	return &ArrayTrie{children: [26]*ArrayTrie{}}
}

// Insert 往前缀树中添加一个元素word
func (a *ArrayTrie) Insert(word string) {
	var curNode = a
	for i := 0; i < len(word); i++ {
		if curNode.children[word[i]-'a'] == nil {
			curNode.children[word[i]-'a'] = NewArrayTrie()
		}
		curNode = curNode.children[word[i]-'a']
	}
	curNode.isWord = true
}

// Search 查找前缀树中是否元素word
func (a *ArrayTrie) Search(word string) bool {
	var _, has = a.search(word)
	return has
}

func (a *ArrayTrie) search(word string) (*ArrayTrie, bool) {
	var curNode = a
	for i := 0; i < len(word); i++ {
		if curNode.children[word[i]-'a'] == nil {
			return nil, false
		}
		curNode = curNode.children[word[i]-'a']
	}
	return curNode, curNode.isWord
}

// HasPrefix 查询前缀树中是否存在前缀prefix
func (a *ArrayTrie) HasPrefix(prefix string) bool {
	var curNode = a
	for i := 0; i < len(prefix); i++ {
		if curNode.children[prefix[i]-'a'] == nil {
			return false
		}
		curNode = curNode.children[prefix[i]-'a']
	}
	return true
}
