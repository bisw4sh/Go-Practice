package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	for length := len(s); length > 0; length-- { // Start from the longest possible length
		for i := 0; i <= len(s)-length; i++ { // Ensure the substring doesn't exceed the string bounds
			if isPalindrome(s[i : i+length]) {
				return s[i : i+length] // Return early as soon as a palindrome is found
			}
		}
	}
	return "" // If no palindrome is found (edge case)
}

func main() {
	result1 := longestPalindrome("babad")
	result2 := longestPalindrome("cbbd")
	fmt.Println(result1)
	fmt.Println(result2)

}
