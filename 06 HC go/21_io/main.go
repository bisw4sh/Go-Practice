package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("Hello,", input)

	fmt.Println("Enter your age:")
	strInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	ageInput, err := strconv.ParseFloat(strings.TrimSpace(strInput), 64)
	if err != nil {
		fmt.Println("Error parsing age:", err)
		return
	}

	fmt.Printf("The age is %.2f\n", ageInput)
}