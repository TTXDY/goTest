package main

import "fmt"

// 总结：38-40行的赋值情况，需要注意是同一个地址。t变，l 也会变

func PrintList(l ListNode) int {
	i := 1
	print("this 1list val: ")
	for ; l.Next != nil; i++ {
		fmt.Print(" ", l.Val, " ")
		l = *l.Next
	}
	fmt.Print(" ", l.Val, " ")
	return i
}

func ReturnLength(l ListNode) int {
	i := 1
	for ; l.Next != nil; i++ {
		l = *l.Next
	}
	return i
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// t是最长的那个list
	var t *ListNode
	var t2 *ListNode
	var small *ListNode

	lengthL1 := ReturnLength(*l1)
	lengthL2 := ReturnLength(*l2)
	if lengthL1 <= lengthL2 {
		small = l1
		t = l2
		t2 = l2

	} else {
		small = l2
		t = l1
		t2 = l1
	}

	for i := 0; i < max(lengthL1, lengthL2)-min(lengthL2, lengthL1); i++ {
		for j := 0; small.Next != nil; j++ {
			small = small.Next
		}
		small.Next = &ListNode{Val: 0, Next: nil}
	}
	k := 0
	for i := 0; i < max(lengthL1, lengthL2); i++ {
		t.Val = l1.Val + l2.Val + k
		if t.Val >= 10 && t.Next != nil {
			k = 1
			t.Val = t.Val - 10
		} else if t.Val >= 10 && t.Next == nil {
			k = 1
			t.Val = t.Val - 10
			t.Next = &ListNode{Val: 1, Next: nil}
			//return t2
		} else {
			k = 0
		}

		t = t.Next
		l1 = l1.Next
		l2 = l2.Next

	}
	return t2
}

func main() {
	l1 := ListNode{
		Val: 9, Next: &ListNode{
			Val: 9, Next: &ListNode{
				Val: 9, Next: &ListNode{
					Val: 9, Next: nil,
				},
			},
		},
	}
	//l2 := ListNode{
	//	Val: 8, Next: &ListNode{
	//		Val: 2, Next: &ListNode{
	//			Val: 3, Next: nil,
	//		},
	//	},
	//}
	l2 := ListNode{
		Val: 9, Next: &ListNode{
			Val: 9, Next: nil,
		},
	}
	var t *ListNode
	t = addTwoNumbers(&l1, &l2)
	PrintList(*t)
	//PrintListReturnLength(l2)
}
