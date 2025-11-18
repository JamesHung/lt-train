package implementtrieprefixtree

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Fatalf("expected apple to be found")
	}
	if trie.Search("app") {
		t.Fatalf("expected app to be absent before insertion")
	}
	if !trie.StartsWith("app") {
		t.Fatalf("expected prefix app to exist")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Fatalf("expected app to be found after insertion")
	}

	if !trie.Search("appxxx") {
		t.Fatalf("expected app to be found after insertion")
	}
}
