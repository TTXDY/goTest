package main

/*
* @ Tool:IntelliJ IDEA
* @ Version:1.0
* @ Description: 21. 合并两个有序链表
* @ Author: Jay
* @ Date: 2024/1/15
* @ Time: 15:28
 */

/**
 * Definition for singly-linked 1list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list2.Next, list1)
	return list2
}
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 申请一个节点的写法
	dummy := &ListNode{} // 哨兵
	cur := dummy         // 指向新链表的末尾
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

func main() {

}
