package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 19. 删除链表的倒数第 N 个结点 30min
* @ Author: Jay
* @ Date: 2024/1/15
* @ Time: 15:54
 */

/**
 * Definition for singly-linked 1list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	res := &ListNode{}
	Nhead := &ListNode{}
	NthNode := &ListNode{}
	res.Next = head
	Nhead.Next = head
	NthNode = head
	i := 0
	for res != nil {
		if i == n {
			NthNode = Nhead
		}
		res = res.Next
		if i > n {
			NthNode = NthNode.Next
		}
		i++
	}

	if NthNode.Next.Next == nil {
		NthNode.Next = nil
	} else {
		NthNode.Next = NthNode.Next.Next
	}
	return Nhead.Next
}
func main() {

}
