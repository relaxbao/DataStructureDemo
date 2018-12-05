package examples

import (
	"fmt"
)

// 合并 k 个排序链表，返回合并后的排序链表。
// 输入:
// [
// 	1->4->5,
// 	1->3->4,
// 	2->6
//   ]
//   输出: 1->1->2->3->4->4->5->6

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	size uint64    // 车辆数量
	head *ListNode // 车头
	tail *ListNode // 车尾
}

func MergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		//针对两个链表进行归并排序
		return merge(lists[0], lists[1])
	default:
		key := n / 2
		//数组拆分,使下一次递归的lists的长度=2

		//优化思路: mergeKLists(lists[:key]),使用Goroutine+channel进行并发合并(归并排序的特点)
		return MergeKLists([]*ListNode{MergeKLists(lists[:key]), MergeKLists(lists[key:])})
	}

}

//merge 对两个有序链表进行归并排序
func merge(left *ListNode, right *ListNode) *ListNode {
	//head: 新的链表的head指针,保持不变
	//tail: 新链表的尾指针
	var head, tail *ListNode

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	if left.Val < right.Val {
		head, tail, left = left, left, left.Next
	} else {
		head, tail, right = right, right, right.Next
	}

	//循环,直到某一个链表已遍历完
	for left != nil && right != nil {
		//找到下一个节点,添加到新链表的尾
		if left.Val < right.Val {
			tail.Next, left = left, left.Next
		} else {
			tail.Next, right = right, right.Next
		}
		//更新tail
		tail = tail.Next
	}

	//剩下的节点字节拼接到新链表尾部
	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}

	return head
}

func printl(l *ListNode) {
	if l.Next != nil {
		fmt.Printf("%v ->", l.Val)
		l = l.Next
	}
}

func (list *List) Init() {
	(*list).size = 0   // 此时链表是空的
	(*list).head = nil // 没有车头
	(*list).tail = nil // 没有车尾
}

func (list *List) Append(v int) bool {
	return list.append(&ListNode{Val: v})
}

func (list *List) append(node *ListNode) bool {
	if node == nil {
		return false
	}

	(*node).Next = nil
	// 将新元素放入单链表中
	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).Next = node
	}

	// 调整尾部位置，及链表元素数量
	(*list).tail = node // node成为新的尾部
	(*list).size++      // 元素数量增加

	return true
}

// func MergeList(l ...*list.List) {
// 	lists := l
// 	for (len){

// 		Lmin := GetMinFrontList(l...)

// 		Lmin.Remove(Lmin.Front())
// 	}

// }

// func GetMinFrontList(l ...*list.List) *list.List {

// 	min := l[0]
// 	MValue := getFrontValue(min)
// 	for _, value := range l {
// 		if getFrontValue(value) < MValue {
// 			min = value
// 		}
// 	}
// 	return min
// }

// func getFrontValue(l *list.List) int {
// 	front := l.Front()
// 	value, _ := front.Value.(int)
// 	return value
// }
