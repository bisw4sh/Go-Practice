package main

import (
	"fmt"
	 "os"
	 "log"
)

func main(){
file, _ := os.Create("output.txt")
sucsess , errmsg := fmt.Fprint(file,"This seems to be the way to file into a file")
fmt.Println(sucsess, errmsg)
file.Close()
// DeleteFile()
}


//Deletes the file 
func DeleteFile(){
	e := os.Remove("output.txt")

    if e != nil { 
        log.Fatal(e) 
    }
	fmt.Println("File has been deleted")
} 