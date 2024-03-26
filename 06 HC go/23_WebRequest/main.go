package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const URL = "http://blogs.biswashdhungana.com.np"

func getReq(){
	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
	panic(err)
	}

	fmt.Println(string(databytes))
}

func postReq(){
	const url = "http://localhost:3000"
 
	reqBody := strings.NewReader(`
	{
		"description": "I wanted to add this 4"
	}`)

	response, err := http.Post(url, "application/json", reqBody)

	if err != nil {
	panic(err)
	}

	defer response.Body.Close()

	content , err := io.ReadAll(response.Body)

	if err != nil {
	panic(err)
	}

	fmt.Println(string(content))

}

func main() {

getReq()
postReq()
}