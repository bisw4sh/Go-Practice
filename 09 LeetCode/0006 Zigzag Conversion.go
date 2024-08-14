package main

import "fmt"

func convert(s string, numRows int) string {
	// Return as it is if numRows is 1
	if numRows == 1 {
		return s
	}

	// Initialize the matrix to store the zigzag pattern
	wMatrix := make([]string, numRows)
	currRow := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		wMatrix[currRow] += string(s[i])

		// Toggle direction when reaching the first or last row
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}

		// Move to the next row
		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	// Concatenate all rows to get the final string
	ret_string := ""
	for _, row := range wMatrix {
		ret_string += row
	}

	return ret_string
}

func main() {
	result := convert("PAYPALISHIRING", 3)
	fmt.Printf("The returned string: %s\n", result)
}
