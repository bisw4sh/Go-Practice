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

		//Only list 1 nodes sum
		if currentRef_l1 != nil && currentRef_l2 == nil {

			currentRef_l3.Val = carry + currentRef_l1.Val
			carry = currentRef_l3.Val / 10
			currentRef_l3.Val = currentRef_l3.Val % 10

			currentRef_l1 = currentRef_l1.Next

		} else if currentRef_l1 == nil && currentRef_l2 != nil {

			currentRef_l3.Val = carry + currentRef_l2.Val
			carry = currentRef_l3.Val / 10
			currentRef_l3.Val = currentRef_l3.Val % 10

			currentRef_l2 = currentRef_l2.Next

		} else if currentRef_l1 == nil && currentRef_l2 == nil && carry > 0 {
			currentRef_l3.Val = carry
			carry = currentRef_l3.Val / 10
			currentRef_l3.Val = currentRef_l3.Val % 10
		} else {
			currentRef_l3.Val = carry + currentRef_l1.Val + currentRef_l2.Val
			carry = currentRef_l3.Val / 10
			currentRef_l3.Val = currentRef_l3.Val % 10

			currentRef_l1 = currentRef_l1.Next
			currentRef_l2 = currentRef_l2.Next

		}

		if currentRef_l1 != nil || currentRef_l2 != nil || carry > 0 {
			new_l3_node := &ListNode{Val: 0, Next: nil}
			currentRef_l3.Next = new_l3_node
			currentRef_l3 = new_l3_node
		}

	}

	return l3 //change to the sum linkedlist
}

func params_to_linked_list(params ...int) *ListNode {
	new_list := &ListNode{Val: 0, Next: nil}
	cur_ref := new_list

	for idx, elem := range params {
		cur_ref.Val = elem

		if idx != len(params)-1 {
			NextNode := &ListNode{Val: 0, Next: nil}
			cur_ref.Next = NextNode
			cur_ref = NextNode
		}

	}

	return new_list
}

func main() {
	l1 := params_to_linked_list(9, 9, 9, 9, 9, 9, 9)
	l2 := params_to_linked_list(9, 9, 9, 9)

	fmt.Print("\n\nResult Itereation\n")
	// Call the addTwoNumbers function
	result := addTwoNumbers(l1, l2)

	// Print the result linked list
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
