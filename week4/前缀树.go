package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func InitTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[r]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, r := range word {
		if next, ok := node.children[r]; !ok {
			return false
		} else {
			node = next
		}
	}
	return node.isEnd
}

// StartsWith 检查前缀树中是否有单词以指定前缀开始
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if next, ok := node.children[char]; !ok {
			return false
		} else {
			node = next
		}
	}
	return true
}

func main() {
	trie := InitTrie()
	trie.Insert("apple")

	fmt.Println(trie.Search("apple"))   // 输出: true
	fmt.Println(trie.Search("app"))     // 输出: false，因为"app"没有作为完整单词插入
	fmt.Println(trie.StartsWith("app")) // 输出: true
}
