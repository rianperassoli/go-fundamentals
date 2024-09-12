package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("You should send a filename as a param")
		os.Exit(1)
	}
	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}

	io.Copy(os.Stdout, f)
}
