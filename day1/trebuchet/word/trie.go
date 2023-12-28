package word

import (
	"slices"

	"github.com/Fabulous-Fadz/aoc23/internal"
)

const (
	validBaseIndex = 0
	validHighIndex = 25
)

type node struct {
	children [26]*node
	isEnd    bool
}

type Trie struct {
	root *node
}

func New() *Trie {
	return &Trie{
		root: &node{
			children: [26]*node{},
		},
	}
}

func (t Trie) Insert(word string) {
	currentNode, wordLength := t.root, len(word)
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if currentNode.children[index] == nil {
			currentNode.children[index] = &node{
				children: [26]*node{},
			}
		}
		currentNode = currentNode.children[index]
	}

	currentNode.isEnd = true
}

func (t Trie) InsertReverse(word string) {
	foo := []byte(word)

	slices.Reverse(foo)
	t.Insert(string(foo))
}

func (t Trie) Contains(word string) bool {
	currentNode, wordLength := t.root, len(word)
	for i := 0; i < wordLength; i++ {
		index := int(word[i] - 'a')
		if index < validBaseIndex || index > validHighIndex {
			return false
		}
		if currentNode.children[index] == nil {
			return false
		}
		currentNode = currentNode.children[index]
	}

	return currentNode.isEnd
}

func (t Trie) ContainsPrefix(word string) bool {
	currentNode, wordLength := t.root, len(word)
	for i := 0; i < wordLength; i++ {
		index := int(word[i] - 'a')
		if index < validBaseIndex || index > validHighIndex {
			return false
		}
		if currentNode.children[index] == nil {
			return false
		}
		currentNode = currentNode.children[index]
	}

	return true
}

func (t Trie) ContainsSuffix(word string) bool {
	currentNode, lastChar := t.root, len(word)-1
	for i := lastChar; i >= 0; i-- {
		if internal.IsDigit(word[i]) {
			return false
		}
		child := currentNode.children[word[i]-'a']
		if child == nil {
			return false
		}
		currentNode = child
	}

	return true
}
