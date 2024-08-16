package main

import (
	"fmt"
)

type Stack struct {
	items []string
}

func push(stack *Stack, s string) {
	stack.items = append(stack.items, s)
}

func pop(stack *Stack) {
	if len(stack.items) == 0 {
		return
	}
	stack.items = stack.items[:len(stack.items)-1]
}

func isValid(s string) bool {
	var string_stack Stack

	//Filled the stack
	for _, bracket := range s {

		switch {
		case string(bracket) == ")":
			if len(string_stack.items) < 1 || string_stack.items[len(string_stack.items)-1] != "(" {
				return false
			} else {
				pop(&string_stack)
			}
		case string(bracket) == "}":
			if len(string_stack.items) < 1 || string_stack.items[len(string_stack.items)-1] != "{" {
				return false
			} else {
				pop(&string_stack)
			}
		case string(bracket) == "]":
			if len(string_stack.items) < 1 || string_stack.items[len(string_stack.items)-1] != "[" {
				return false
			} else {
				pop(&string_stack)
			}
		default:
			push(&string_stack, string(bracket))
		}

		// fmt.Println(string_stack.items)
	}

	return len(string_stack.items) == 0
}

func main() {
	s1 := isValid("()")
	s2 := isValid("()[]{}")
	s3 := isValid("(]")
	s4 := isValid("([)]")
	s5 := isValid("{[]}")
	s6 := isValid("{[]}")
	s7 := isValid("[")
	fmt.Println(s1, s2, s3, s4, s5, s6, s7)

	// fmt.Println(s5)

}
