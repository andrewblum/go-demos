package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args will contain an array with:
	// 1 - the path to the temporary file that is created when our program is compiled
	// 2 - everything typed after "go run main.go", as split by spaces

	filename := os.Args[1:]
	fmt.Println(filename)
	f, err := os.Open(filename[0])
	if err != nil {
		fmt.Println("error!", err)
		os.Exit(1)
	}
	// copy wants a writer and then a reader
	// we give it os.Stdout as the writer
	// and we give it the file we opened as the reader
	io.Copy(os.Stdout, f)
}
