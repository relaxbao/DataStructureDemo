package examples

import (
	"container/list"
	"fmt"
)

// > 给定两个非空链表来表示两个非负整数,位数按照逆序方式存储,他们的每个节点只存储单个数字,将两数相加返回一个新的链表.你可以假设除了数字0之外,这两个数字不会以0开头.
// 示例:
// 输入:(2->4->3) + (5->6->4)
// 输出:7->0->8
// 原因:342+465 = 807

func AddLists(l1, l2 *list.List) *list.List {
	l := list.New()

	max, min := SortList(l1, l2)
	Emax := max.Back()
	Emin := min.Back()
	temp := 0
	for i := 0; i < min.Len(); i++ {
		vmax, _ := Emax.Value.(int)
		vmin, _ := Emin.Value.(int)
		v := (temp + vmax + vmin) % 10
		l.PushBack(v)
		temp = (vmax + vmin) / 10
		Emax = Emax.Prev()
		Emin = Emin.Prev()
	}

	if max.Len() > min.Len() {
		for i := 0; i < (max.Len() - min.Len()); i++ {
			v, _ := Emax.Value.(int)
			l.PushBack(v)
			Emax = Emax.Prev()
		}
	}

	return l
}

func printList(l *list.List) {
	e := l.Front()
	fmt.Printf("%v", e.Value)
	e = e.Next()
	for i := 1; i < l.Len(); i++ {
		fmt.Printf(" - > %v", e.Value)
		e = e.Next()
	}
	fmt.Println()
}

// 将l1,l2 排序,位数大的排前面
func SortList(l1, l2 *list.List) (max, min *list.List) {

	if l1.Len() >= l2.Len() {
		return l1, l2
	} else {
		return l2, l1
	}
}
