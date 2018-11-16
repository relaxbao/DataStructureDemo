package main

import (
	"errors"
)

var (
	Max_Len        = 20
	ErrNoExist     = errors.New("Element Does Not Exist!")
	ErrUnvalideLen = errors.New("The Length Is Unvalide!")
	ErrMaxLen      = errors.New("The Lenth Has Reach The Maximum!")
	ErrEmpty       = errors.New("The List Is Empty!")
)

type LineList struct {
	MaxLength int
	Length    int
	Content   []interface{}
}

type List interface {
	IsEmpty() bool
	IsFull() bool
	IndexOver() bool
	ClearList()
	GetElem(i int) (interface{}, error)
	LocateElem(value interface{}) (int, error)
	ListInsert(i int, v interface{}) error
	ListDelete(i int) error
}

// Init a Sequenatial List
func InitList() *LineList {

	return &LineList{
		MaxLength: Max_Len,
		Content:   make([]interface{}, 0, Max_Len),
	}
}

// whether the Sequenatial List is Empty
func (l *LineList) IsEmpty() bool {
	return l.Length == 0
}

// whether the Sequenatial List is full
func (l *LineList) IsFull() bool {
	return l.Length == l.MaxLength
}

func (l *LineList) IndexOver(i int) bool {
	return (i < 1 || i > l.Length)
}

// clear the List
func (l *LineList) ClearList() {
	if !l.IsEmpty() {
		l.Length = 0
		l.Content = make([]interface{}, 0, Max_Len)
	}

}

// Get the Element with index i, return the value of the Element
func (l *LineList) GetElem(i int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, ErrEmpty
	}
	if l.IndexOver(i) {
		return nil, ErrUnvalideLen
	}
	return l.Content[i-1], nil
}

// Get the Element with  value  v, return its index
func (l *LineList) LocateElem(v interface{}) (int, error) {
	if l.IsEmpty() {
		return 0, ErrEmpty
	}
	for i := 0; i < l.Length; i++ {
		if l.Content[i] == v {
			return i, nil
		} else {
			continue
		}
	}
	return 0, ErrNoExist
}

func (l *LineList) Append(v interface{}) error {
	if l.IsFull() {
		return ErrMaxLen
	}
	l.Content = append(l.Content, v)
	l.Length++
	return nil
}

func (l *LineList) Pop() (interface{}, error) {
	if l.IsEmpty() {
		return nil, ErrEmpty
	}
	v := l.Content[l.Length-1]
	l.Content = l.Content[:l.Length-2]
	l.Length--
	return v, nil
}

// Insert v where the index is index
func (l *LineList) ListInsert(index int, v interface{}) error {

	if l.IsFull() {
		return ErrMaxLen
	}

	if l.IndexOver(index) {
		return ErrUnvalideLen
	}

	// Add an empty data first to prevent the following access from crossing the boundary
	l.Append("")
	for i := l.Length - 1; i > index-1; i-- {
		//从后往前赋值，新增一个空node，然后把数据一个个后移，直到插入的位置
		//注意线性表从１开始，而切片是从０开始的
		// first add an empty node,then assign value from the end,and then move the data one by one utill the inserted position
		// Note that the Sequenatial List starts with 1,and the slice starts with 0
		l.Content[i] = l.Content[i-1]
	}
	l.Content[index-1] = v

	return nil

}

// Delete the Element with index index
func (l *LineList) ListDelete(index int) error {
	if l.IsEmpty() {
		return ErrEmpty
	}
	if l.IndexOver(index) {
		return ErrUnvalideLen
	}

	for i := index - 1; i < l.Length-1; i++ {
		l.Content[i] = l.Content[i+1]
	}
	l.Content = l.Content[:l.Length-1]
	l.Length--
	return nil
}
