package trie

// Trie
type Trie interface {
	Insert(string)
	Search(string) bool
}
