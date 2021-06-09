package trie

// ListTrie 列表前缀树
type ListTrie struct {
	isWord   bool
	char     byte
	children []*ListTrie
}

var _ Trie = (*ListTrie)(nil)

// NewListTrie new listTrie
func NewListTrie() *ListTrie {
	return new(ListTrie)
}

func (a *ListTrie) Insert(word string) {
	var (
		curNode = a
		tNode   *ListTrie
	)
	for i := 0; i < len(word); i++ {
		if tNode = curNode.find(word[i]); tNode == nil {
			curNode.children = append(curNode.children, &ListTrie{char: word[i]})
		}
		curNode = curNode.find(word[i])
	}
	curNode.isWord = true
}

func (a *ListTrie) find(c byte) *ListTrie {
	for i := 0; i < len(a.children); i++ {
		if a.children[i].char == c {
			return a.children[i]
		}
	}
	return nil
}

func (a *ListTrie) Search(word string) bool {
	var _, has = a.search(word)
	return has
}

func (a *ListTrie) search(word string) (*ListTrie, bool) {
	var curNode = a
	for i := 0; i < len(word); i++ {
		curNode = curNode.find(word[i])
		if curNode == nil {
			return nil, false
		}
	}
	return curNode, curNode.isWord
}

func (a *ListTrie) HasPrefix(prefix string) bool {
	var curNode = a
	for i := 0; i < len(prefix); i++ {
		curNode = curNode.find(prefix[i])
		if curNode == nil {
			return false
		}
	}
	return true
}

func (a *ListTrie) Delete(word string) bool {
	var node, has = a.search(word)
	if !has {
		return false
	}
	node.isWord = false
	return true
}
