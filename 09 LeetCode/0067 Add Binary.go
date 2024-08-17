package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	alen := len(a)
	blen := len(b)

	if alen > blen {
		str_to_append := strings.Repeat("0", alen-blen)
		b = str_to_append + b
	} else {
		str_to_append := strings.Repeat("0", blen-alen)
		a = str_to_append + a
	}

	result := ""
	carry := "0"

	for i := len(a) - 1; i >= 0; i-- {
		switch {
		case string(a[i]) == "0" && string(b[i]) == "0":
			if carry == "0" {
				result = "0" + result
			} else {
				result = "1" + result
				carry = "0"
			}
		case (string(a[i]) == "0" && string(b[i]) == "1") || (string(a[i]) == "1" && string(b[i]) == "0"):
			if carry == "0" {
				result = "1" + result
			} else {
				result = "0" + result
			}
		case string(a[i]) == "1" && string(b[i]) == "1":

			if carry == "0" {
				result = "0" + result
				carry = "1"
			} else {
				result = "1" + result
			}
		default:
		}
	}

	if carry == "1" {
		result = "1" + result
		return result
	}
	return result
}

func main() {
	result_1 := addBinary("11", "1")
	fmt.Println(result_1)

	result_2 := addBinary("1010", "1011")
	fmt.Println(result_2)
}
