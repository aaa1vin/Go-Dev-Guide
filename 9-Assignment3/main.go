package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	// f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0644)

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}