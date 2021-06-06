package trie

type Trie interface {
	Insert(string)
	Search(string) bool
}
