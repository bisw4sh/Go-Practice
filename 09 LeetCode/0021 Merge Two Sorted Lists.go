package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged_list := &ListNode{}
	head_merged_list := merged_list

	for {
		if list1 == nil {
			head_merged_list.Next = list2
			break
		} else if list2 == nil {
			head_merged_list.Next = list1
			break
		}

		if list1.Val < list2.Val {
			head_merged_list.Next = list1
			list1 = list1.Next
		} else {
			head_merged_list.Next = list2
			list2 = list2.Next
		}
		head_merged_list = head_merged_list.Next
	}

	return merged_list.Next
}

func main() {

	// Create list1: 1 -> 3 -> 5
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 3}
	list1.Next.Next = &ListNode{Val: 5}

	// Create list2: 2 -> 4 -> 6
	list2 := &ListNode{Val: 2}
	list2.Next = &ListNode{Val: 4}
	list2.Next.Next = &ListNode{Val: 6}

	// Merge the two lists
	mergedList := mergeTwoLists(list1, list2)

	// Print the merged list
	fmt.Println("Merged List:")

	for mergedList != nil {
		fmt.Println(mergedList.Val)
		mergedList = mergedList.Next
	}
}
