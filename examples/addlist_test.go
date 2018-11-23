package examples

import (
	"container/list"
	"fmt"
	"testing"
)

// > 给定两个非空链表来表示两个非负整数,位数按照逆序方式存储,他们的每个节点只存储单个数字,将两数相加返回一个新的链表.你可以假设除了数字0之外,这两个数字不会以0开头.
// 示例:
// 输入:(2->4->3) + (5->6->4)
// 输出:7->0->8
// 原因:342+465 = 807
func TestAddList(t *testing.T) {
	l1 := list.New()
	l2 := list.New()

	l1.PushBack(2)
	l1.PushBack(4)
	l1.PushBack(3)

	fmt.Println("l1 =====")
	printList(l1)

	l2.PushBack(5)
	l2.PushBack(6)
	l2.PushBack(4)
	fmt.Println("l2 =====")
	printList(l2)

	lnew := AddLists(l1, l2)

	fmt.Println("lnew =====")
	printList(lnew)
}
