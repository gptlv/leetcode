package main

func main() {

}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, ch := range word {
		if _, ok := cur.children[ch]; !ok {
			cur.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, ch := range word {
		if child, ok := cur.children[ch]; ok {
			cur = child
		} else {
			return false
		}
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, ch := range prefix {
		if child, ok := cur.children[ch]; !ok {
			return false
		} else {
			cur = child
		}
	}
	return true
}
