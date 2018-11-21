package main

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []interface{}
	}{
		{"index = 1", 1, []interface{}{"k", "1", "2", "3"}},
		{"index = 3(MaxLen)", 3, []interface{}{"1", "2", "k", "3"}},
		{"index = 4(IndexOver)", 4, []interface{}{"1", "2", "3"}},
		{"index = 2", 2, []interface{}{"1", "k", "2", "3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := InitListTest()
			err := list.ListInsert(tt.args, "k")
			if err != nil {
				fmt.Println("err:", err)
			}
			if got := list.Content; !SliceEqualBCE(got, tt.want) {
				t.Errorf("TestDelete Failed got = %v , want = %v", got, tt.want)
			}

		})
	}

}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []interface{}
	}{
		{"index = 1", 1, []interface{}{"2", "3"}},
		{"index = 3(MaxLen)", 3, []interface{}{"1", "2"}},
		{"index = 4(IndexOver)", 4, []interface{}{"1", "2", "3"}},
		{"index = 2", 2, []interface{}{"1", "3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := InitListTest()
			_, err := list.ListDelete(tt.args)
			if err != nil {
				fmt.Println("err:", err)
			}
			if got := list.Content; !SliceEqualBCE(got, tt.want) {
				t.Errorf("TestDelete Failed got = %v , want = %v", got, tt.want)
			}
		})
	}
}

func InitListTest() *LineList {
	list := InitList()
	list.Append("1")
	list.Append("2")
	list.Append("3")
	return list
}

//
func SliceEqualBCE(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
