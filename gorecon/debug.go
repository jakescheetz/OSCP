package main

import (
	"fmt"
	"os"
)
func main(){
if (len(os.Args) < 1) || (os.Args[1] == "-h") {
		fmt.Println(string(len(os.Args)))
		return
	} else {
		//print banner
		fmt.Println("print banner")
	}
}
