package main

func main() {

}

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

func (root *Node) search(word string) bool {
	node := root
	for i, ch := range word {
		if _, ok := node.children[ch]; ok {
			node = node.children[ch]
		} else if ch == '.' {
			if len(node.children) == 0 {
				return false
			}
			for _, child := range node.children {
				if child.search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			return false
		}
	}

	return node.isEnd
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{&Node{children: make(map[rune]*Node)}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &Node{children: make(map[rune]*Node)}
		}

		node = node.children[ch]
	}

	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	node := this.root
	return node.search(word)

}
