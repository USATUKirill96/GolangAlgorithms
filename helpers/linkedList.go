package helpers

import "fmt"

type LinkedList struct {
	Root *LinkedListNode
}

func (ll *LinkedList) FromList(list []int) {
	for _, v := range list {
		ll.Append(v)
	}
}

func (ll *LinkedList) Append(v int) {
	if ll.Root != nil {
		ll.Root.Insert(v)
	} else {
		ll.Root = &LinkedListNode{Value: v}
	}
}

func (ll *LinkedList) String() string {
	return fmt.Sprint(ll.Root)
}

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func (n *LinkedListNode) String() string {
	return fmt.Sprintf("%v %s", n.Value, n.Next)
}

func (n *LinkedListNode) Insert(v int) {
	if n.Next == nil {
		n.Next = &LinkedListNode{Value: v}
	} else {
		n.Next.Insert(v)
	}
}
