package telerouter

import (
	"github.com/dghubble/trie"
)

type Router struct {
	*trie.PathTrie
}

func NewRouter() *Router {
	return &Router{
		PathTrie: trie.NewPathTrie(),
	}
}
