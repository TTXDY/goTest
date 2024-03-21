package main

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	list := make([]int, 0)
	// head中所有元素添加到list中
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	length := len(list)
	for i, _ := range list[:length/2] {
		if list[i] != list[length-1-i] {
			return false
		}
	}
	return true
}
