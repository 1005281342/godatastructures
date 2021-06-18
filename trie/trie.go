package trie

// Type 前缀树类型
type Type int32

const (
	// EmHashTrie 哈希前缀树
	EmHashTrie = iota
	// EmListTrie 列表前缀树
	EmListTrie
	// EmArrayTrie 数组前缀树
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
