package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 328. 奇偶链表
* @ Author: Jay
* @ Date: 2024/1/15
* @ Time: 17:18
 */

/**
 * Definition for singly-linked 1list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main() {
	thelist := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	oddEvenList(thelist)
}
