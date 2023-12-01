package main

import "fmt"

func main() {
	var operand1, operand2 int

	fmt.Println("Enter two operands: ")
	fmt.Scanln(&operand1, &operand2)

	var add, sub, mul, div int
	add, sub, mul, div = ArithmeticFunc(operand1, operand2)

	fmt.Printf("Addition: %d\nSubtraction: %d\nMultiplication: %d\nDivision: %d\n", add, sub, mul, div)
}


func ArithmeticFunc(op1 int, op2 int) (int, int, int, int){
return op1 + op2, op1 - op2, op1 * op2, op1 / op2
}