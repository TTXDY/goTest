package main

func save(l *ListNode) (map[int]int, int) {
	var valueMap = make(map[int]int)
	i := 0
	for l != nil {
		valueMap[i] = l.Val
		l = l.Next
		i++
	}
	return valueMap, i
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var valueMap = make(map[int]int)
	var length int
	// 暂存head头指针
	var first *ListNode
	first = head
	valueMap, length = save(head)

	for i := length - 1; first != nil; i-- {
		first.Val = valueMap[i]
		first = first.Next
	}
	return head
}

func main() {
	l1 := ListNode{
		Val: 8, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: nil,
			},
		},
	}

	reverseList(&l1)
}
