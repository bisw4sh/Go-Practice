package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currentRef_l1 := l1
	currentRef_l2 := l2

	l3 := &ListNode{Val: 0, Next: nil}
	currentRef_l3 := l3
	carry := 0

	for currentRef_l1 != nil || currentRef_l2 != nil || carry > 0 {

		if currentRef_l1 != nil {
			carry += currentRef_l1.Val
			currentRef_l1 = currentRef_l1.Next
		}

		if currentRef_l2 != nil {
			carry += currentRef_l2.Val
			currentRef_l2 = currentRef_l2.Next
		}

		currentRef_l3.Val = carry % 10
		carry /= 10

		if currentRef_l1 != nil || currentRef_l2 != nil || carry > 0 {
			currentRef_l3.Next = &ListNode{Val: 0, Next: nil}
			currentRef_l3 = currentRef_l3.Next
		}
	}

	return l3
}

func params_to_linked_list(params ...int) *ListNode {
	if len(params) == 0 {
		return nil
	}

	new_list := &ListNode{Val: params[0], Next: nil}
	cur_ref := new_list

	for _, elem := range params[1:] {
		cur_ref.Next = &ListNode{Val: elem, Next: nil}
		cur_ref = cur_ref.Next
	}

	return new_list
}

func main() {
	l1 := params_to_linked_list(9, 9, 9, 9, 9, 9, 9)
	l2 := params_to_linked_list(9, 9, 9, 9)

	fmt.Print("\n\nResult Iteration\n")
	// Call the addTwoNumbers function
	result := addTwoNumbers(l1, l2)

	// Print the result linked list
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
