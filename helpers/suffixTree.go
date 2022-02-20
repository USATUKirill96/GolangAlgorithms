package helpers

import (
	"fmt"
	"strings"
)

type SuffixTree struct {
	root *suffixTreeNode
}

func (st *SuffixTree) String() string {
	return fmt.Sprint(st.root)
}

func (st *SuffixTree) Insert(word string) {

	if st.root == nil {
		st.root = newNode()
	}

	letters := strings.Split(word, "")
	letters = append(letters, "*")

	for i := range letters {
		currentNode := st.root
		for j := i; j < len(letters); j++ {
			currentNode = currentNode.getOrCreate(letters[j])
		}
	}
}

func (st *SuffixTree) Search(word string) bool {
	letters := strings.Split(word, "")
	currentNode := st.root
	var ok bool
	for _, v := range letters {
		currentNode, ok = currentNode.children[v]
		if !ok {
			return false
		}
	}
	_, ok = currentNode.children["*"]
	return ok
}

func newNode() *suffixTreeNode { return &suffixTreeNode{children: make(map[string]*suffixTreeNode)} }

type suffixTreeNode struct {
	children map[string]*suffixTreeNode
}

func (stn *suffixTreeNode) getOrCreate(value string) *suffixTreeNode {
	node, ok := stn.children[value]
	if ok {
		return node
	}
	newNode := newNode()
	stn.children[value] = newNode
	return newNode
}

func (stn *suffixTreeNode) String() string {
	return stn.string(0)
}
func (stn *suffixTreeNode) string(indents int) string {
	var children string
	indent := strings.Repeat(" ", indents)

	for key, value := range stn.children {
		children += fmt.Sprintf("\r\n %v %v: %v", indent, key, value.string(indents+2))
	}

	return children
}
