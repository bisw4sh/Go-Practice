package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello!, my name is Biswash Dhungana"

	file, err := os.Create("./content.txt")
	
	if err != nil {
		fmt.Println("Error Occured")
		panic(err)
	}

	length , err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The length is %d\n", length)
	readFile()
	defer file.Close()

}

func readFile(){
	fileByte, err := os.ReadFile("./content.txt")

		if err != nil {
		fmt.Println("Error Occured")
		panic(err)
	}

	fmt.Println(string(fileByte))
}