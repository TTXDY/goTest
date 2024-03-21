package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList1(l ListNode) int {
	i := 1
	print("this 1list val: ")
	for ; l.Next != nil; i++ {
		fmt.Print(" ", l.Val, " ")
		l = *l.Next
	}
	fmt.Print(" ", l.Val, " ")
	return i
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func main() {
	l1 := ListNode{
		Val: 1, Next: &ListNode{
			Val: 5, Next: &ListNode{
				Val: 6, Next: &ListNode{
					Val: 9, Next: nil,
				},
			},
		},
	}
	l2 := ListNode{
		Val: 2, Next: &ListNode{
			Val: 4, Next: &ListNode{
				Val: 7, Next: nil,
			},
		},
	}
	//mergeTwoLists(&l1, &l2)

	x := mergeTwoLists(&l1, &l2)
	//PrintList1(l1)
	PrintList1(*x)
}
