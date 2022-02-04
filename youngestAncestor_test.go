package main

import (
	"USATUKirill96/leetcode/helpers"
	"fmt"
	"math"
	"testing"
)

func TestYoungestAncestor(t *testing.T) {
	tree := &helpers.BST{}
	tree.FromList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(tree)
	res := YoungestAncestor(tree, 1, 4)
	fmt.Println(res)
}

func YoungestAncestor(tree *helpers.BST, firstChild, secondChild int) *helpers.Node {
	var track func(*helpers.Node, int) []*helpers.Node
	track = func(n *helpers.Node, target int) []*helpers.Node {
		if n.Key == target {
			return []*helpers.Node{}
		} else if target < n.Key {
			return append([]*helpers.Node{n}, track(n.Left, target)...)
		} else {
			return append([]*helpers.Node{n}, track(n.Right, target)...)
		}
	}

	firstChildTrack := track(tree.Root, firstChild)
	secondChildTrack := track(tree.Root, secondChild)

	var youngestAncestor *helpers.Node

	for i := 0; i < int(math.Min(float64(len(firstChildTrack)), float64(len(secondChildTrack)))); i++ {
		if firstChildTrack[i].Key != secondChildTrack[i].Key {
			return youngestAncestor
		}
		youngestAncestor = firstChildTrack[i]
	}
	return youngestAncestor
}
