package helpers

import (
	"fmt"
	"testing"
)

func TestBSTFromList(t *testing.T) {
	bst := BST{}
	bst.FromList([]int{1, 2, 5, 7, 10, 13, 14, 15, 22})
	fmt.Println(bst.Root)
}

func TestBSTInvert(t *testing.T) {
	tree := &BST{}
	tree.FromList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("Original: ", *tree)
	tree.Invert()
	fmt.Println("Inverted: ", *tree)
}
