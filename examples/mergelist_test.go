package examples

import (
	"fmt"
	"testing"
)

func TestMergelist(t *testing.T) {
	var l1, l2, l3 List
	l1.Init()
	l2.Init()
	l3.Init()

	l1.Append(1)
	l1.Append(4)
	l1.Append(5)
	l1.printl()

	l2.Append(2)
	l2.Append(3)
	l2.Append(8)
	l2.printl()

	l3.Append(1)
	l3.Append(6)
	l3.Append(9)
	l3.printl()

	var lists []*ListNode
	lists = append(lists, l1.head)
	lists = append(lists, l2.head)
	lists = append(lists, l3.head)
	newnode := mergeKLists(lists)
	printlistnode(newnode)
}

func (l *List) printl() {
	first := l.head

	for first.Next != nil {
		fmt.Printf("%v - > ", first.Val)
		first = first.Next
	}
	fmt.Printf("%v\n", first.Val)
}

func printlistnode(n *ListNode) {
	for n.Next != nil {
		fmt.Printf("%v - > ", n.Val)
		n = n.Next
	}
	fmt.Printf("%v\n", n.Val)
}
