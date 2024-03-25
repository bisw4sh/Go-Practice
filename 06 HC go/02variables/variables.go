package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var user string = "Biswash"
	surname := "Dhungana"
	binaryVal := false
	manyVal := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(user, surname)
	fmt.Println(binaryVal, manyVal)

	var eightbitint uint8 = 255
	var sixteenbitint uint16 = 65535
	var thirthytwobitint uint32 = 4294967295
	var sixtyfourbitint uint64 = 18446744073709551615
	fmt.Printf("Max of 8 bit Integer -> %v\n",eightbitint)
	fmt.Printf("Max of 16 bit Integer -> %v\n",sixteenbitint)
	fmt.Printf("Max of 32 bit Integer -> %v\n",thirthytwobitint)
	fmt.Printf("Max of 64 bit Integer -> %v\n",sixtyfourbitint)

	var z complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Value of a complex number z -> %v", z)


	var minflo float32 = 2.2E-308
	var maxflo float32 = 1.7E+38
	var maxflo64 float64 = 1.7E+38

	fmt.Printf("Min Float -> %v %T and Max Float -> %v %T\n",minflo, minflo,maxflo, maxflo)
	fmt.Printf("Min Float -> %v %T and Max Float -> %v %T\n",minflo, minflo,maxflo, maxflo)
	fmt.Printf("Float 64 -> %v %T\n",maxflo64, maxflo64)
}