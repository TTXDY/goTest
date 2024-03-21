package main

import "slices"

func hasCycle3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var flag bool
	over := make([]*ListNode, 0)
	over = append(over, head)
	for head.Next != nil {
		// 如果不在里面 说明不是环形，返回false， slices.Contains(over, head.Next)判断为false
		if slices.Contains(over, head.Next) {
			flag = true
			return flag
		} else {
			flag = false
		}
		over = append(over, head.Next)
		head = head.Next

	}
	return flag
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var first *ListNode
	var second *ListNode
	first = head
	second = head.Next
	for first != second {
		if second == nil || second.Next == nil {
			return false
		}
		first = first.Next
		second = second.Next.Next
	}
	return true
}

// best

func hasCycle1(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
func main() {

}
