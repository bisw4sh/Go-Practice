package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	rev := 0
	original := x
	for x > 0 {
		digit := x % 10
		rev = rev*10 + digit
		x = x / 10
	}

	return rev == original
}

func main() {
	plalindrome := isPalindrome(14341)
	fmt.Println(plalindrome)
}
