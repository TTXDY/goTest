package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 24. 两两交换链表中的节点 40min
* @ Author: Jay
* @ Date: 2024/1/15
* @ Time: 16:23
 */

/**
 * Definition for singly-linked 1list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	res := &ListNode{}
	Nhead := &ListNode{}
	Nhead.Next = head
	res = Nhead
	for Nhead != nil {
		if Nhead.Next.Next != nil && Nhead.Next != nil {
			tempN := &ListNode{}
			tempNN := &ListNode{}
			tempNNN := &ListNode{}
			tempN = Nhead.Next
			tempNN = Nhead.Next.Next
			tempNNN = Nhead.Next.Next.Next
			Nhead.Next = tempNN
			Nhead.Next.Next = tempN
			Nhead.Next.Next.Next = tempNNN
			Nhead = Nhead.Next.Next
		}
		if Nhead.Next != nil && Nhead.Next.Next != nil {
			continue
		} else if Nhead.Next == nil || Nhead.Next.Next == nil {
			break
		}
	}
	return res.Next
}

func main() {
	thelist := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	swapPairs(thelist)
}
