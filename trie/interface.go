package trie

type Trie interface {
	Insert(string)
	Search(string) bool
	HasPrefix(string) bool
	Delete(string) bool
}
