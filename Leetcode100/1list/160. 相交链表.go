package main

import (
	"slices"
)

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	over := make([]*ListNode, 0)
	// 存储headA 所有节点
	for headA != nil {
		over = append(over, headA)
		headA = headA.Next
	}
	//var temp *ListNode
	//temp = headB
	for headB != nil {
		if slices.Contains(over, headB) {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	over := make(map[*ListNode]int)
	// 存储headA 所有节点
	for headA != nil {
		over[headA] = 0
		headA = headA.Next
	}
	for headB != nil {
		if over[headB] == 0 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
