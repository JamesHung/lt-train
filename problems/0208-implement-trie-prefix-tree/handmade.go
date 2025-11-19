package implementtrieprefixtree

import "fmt"

type Trie struct {
	isEnd    bool
	children map[rune]*Trie
}

const debugTrie = true

func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	if debugTrie {
		fmt.Printf("Insert %q:\n", word)
	}
	current := t
	for _, c := range word {
		if debugTrie {
			fmt.Printf("  at %p processing %q\n", current, string(c))
		}
		if current.children == nil {
			current.children = make(map[rune]*Trie)
		}

		if current.children[c] == nil {
			current.children[c] = &Trie{children: make(map[rune]*Trie)}
		}

		current = current.children[c]
	}

	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.find(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}

func (t *Trie) find(s string) *Trie {
	if debugTrie {
		fmt.Printf("Find %q:\n", s)
	}
	current := t
	for _, ch := range s {
		if debugTrie {
			fmt.Printf("  at %p looking for %q\n", current, string(ch))
		}
		if current.children == nil {
			if debugTrie {
				fmt.Printf("    no children -> fail\n")
			}
			return nil
		}
		current = current.children[ch]
		if current == nil {
			if debugTrie {
				fmt.Printf("    child missing -> fail\n")
			}
			return nil
		}
	}

	if debugTrie {
		fmt.Printf("  success at %p end=%v\n", current, current.isEnd)
	}
	return current
}
