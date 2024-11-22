package main

import "fmt"
import "os"

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	os.Exit(255)
}