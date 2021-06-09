package trie

// Type 前缀树类型
type Type int32

const (
	EmHashTrie = iota
	EmListTrie
	EmArrayTrie
)

// NewTrie 创建一个Trie
func NewTrie(t Type) Trie {
	switch t {
	case EmHashTrie:
		return NewHashTrie()
	case EmListTrie:
		return NewListTrie()
	case EmArrayTrie:
		return NewArrayTrie()
	default:
		// HashTrie较为通用
		return NewHashTrie()
	}
}
