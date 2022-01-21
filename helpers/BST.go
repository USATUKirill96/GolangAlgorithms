package helpers

import (
	"fmt"
	"sort"
	"strings"
)

type BST struct {
	Root *Node
}

func (bst *BST) Insert(v int) {
	if bst.Root == nil {
		bst.Root = &Node{Key: v}
	} else {
		bst.Root.Insert(v)
	}
}
func (bst *BST) String(indents int) string {
	return fmt.Sprint(bst.Root)
}

// FromList Forms a BST with minimal depth
func (bst *BST) FromList(values []int) {
	sort.Ints(values)

	var insert func([]int)
	insert = func(branch []int) {
		if len(branch) == 0 {
			return
		}
		middle := branch[len(branch)/2]
		bst.Insert(middle)
		left := branch[:len(branch)/2]
		insert(left)
		right := branch[len(branch)/2+1:]
		insert(right)
	}
	insert(values)
}

func (bst *BST) Invert() {
	var invert func(*Node)
	invert = func(node *Node) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		invert(node.Left)
		invert(node.Right)
	}
	invert(bst.Root)
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (node *Node) Insert(v int) {
	if v >= node.Key {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = &Node{Key: v}
		}
	} else {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = &Node{Key: v}
		}
	}
}

func (node *Node) Get(v int) *Node {
	switch {
	case v > node.Key:
		if node.Right != nil {
			return node.Right.Get(v)
		}
	case v < node.Key:
		if node.Left != nil {
			return node.Left.Get(v)
		}
	case v == node.Key:
		return node
	default:
		return nil
	}
	return nil
}

func (node *Node) String() string {
	return node.string(0)
}
func (node *Node) string(indents int) string {
	var left, right string
	indent := strings.Repeat(" ", indents)
	if node.Left != nil {
		left = fmt.Sprintf("\r\n %v Left: %v", indent, node.Left.string(indents+2))
	}
	if node.Right != nil {
		right = fmt.Sprintf("\r\n %v Right: %v", indent, node.Right.string(indents+2))
	}
	if left == "" && right == "" {
		return fmt.Sprintf("\r\n %v Leaf: %v", indent, node.Key)
	}
	return fmt.Sprintf(
		"node: %v %v %v",
		node.Key, left, right,
	)
}
