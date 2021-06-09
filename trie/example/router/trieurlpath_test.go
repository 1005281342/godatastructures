package router

import (
	"testing"
)

func TestURLPathTrie(t *testing.T) {
	var myTrie = NewURLPathTrie()
	myTrie.Insert("/")
	t.Logf("search path:`/`, has: %v, want: true", myTrie.Search("/"))
	t.Logf("search path:`/cc`, has: %v, want: false", myTrie.Search("/cc"))
	myTrie.Insert("/cc")
	t.Logf("search path:`/cc`, has: %v, want: true", myTrie.Search("/cc"))
}
