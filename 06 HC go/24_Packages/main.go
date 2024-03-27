package main

import (
	"fmt"
	f "packages/father"
	c "packages/father/child"
)

func main() {
	f.Father()
	c.Child()
	printing := fmt.Sprintf("Father is %s and Child is %s\n", f.F, c.C)
	fmt.Println(printing)
}
