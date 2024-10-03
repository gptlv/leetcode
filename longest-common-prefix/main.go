package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	trie := NewTrie()

	for _, str := range strs {
		trie.Insert(str)
	}

	prefix := []rune{}
	cur := trie.root

	for {
		if len(cur.children) != 1 || cur.isEndOfWord {
			break
		}

		for char, node := range cur.children {
			prefix = append(prefix, char)
			cur = node
		}
	}

	return string(prefix)
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			currentNode.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		currentNode = currentNode.children[char]
	}

	currentNode.isEndOfWord = true
}
