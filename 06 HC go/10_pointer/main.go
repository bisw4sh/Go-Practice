package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	b := &v
	b.X = 1e9

	//To access the field X of a struct when we have the struct pointer 
	//! b we could write (*b).X. However, that notation is cumbersome, so the language permits us instead to write just b.X, without the explicit dereference.
	fmt.Println(v)
}