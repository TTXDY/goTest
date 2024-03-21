package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description:83. 删除排序链表中的重复元素 7min
* @ Author: Jay
* @ Date: 2024/1/15
* @ Time: 15:46
 */

/**
 * Definition for singly-linked 1list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := &ListNode{}
	cur = head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			//temp := &ListNode{}
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {

}
