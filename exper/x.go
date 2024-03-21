package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Prin(l *ListNode) {
	for l != nil {
		fmt.Println("val:", l.Val)
		l = l.Next
	}
}

func main() {
	var l1 *ListNode
	l1 = &ListNode{
		Val: 8, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: nil,
			},
		},
	}
	var l2 *ListNode
	l2 = l1
	for l1 != nil {
		fmt.Println("val:", l1.Val)
		l1 = l1.Next
	}
	fmt.Println(l2.Val)
	m := make(map[int]int)
	m[0]++
	fmt.Println(m)

}
