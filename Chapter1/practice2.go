// +build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Print("index->", i)
		fmt.Print(" : ")
		fmt.Print("value->", arg)
		fmt.Println("")
	}
}
