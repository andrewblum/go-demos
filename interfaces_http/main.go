package main

import (
	"fmt"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	// alternative short hand
	// https://pkg.go.dev/io#Copy
	// command and hover COPY or right click go to def
	// io.Copy(os.Stdout, resp.Body)

	// lets make our own!
	// lw := logWriter{}
	// io.Copy(lw, resp.Body)
}

// satisfying the an interface isnt that meaningful, as seen below
// func (logWriter) Write(bs []byte) (int, error) {
// 	return 1, nil
// }

// because the interface writer exsist in the Go code and our logWriter
// has a receiver implenting all of the (in this case only one)
// things that it specifies, our logWriter is now said to implement that interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	// supposed to return the number of bytes written and any errors encountered
	return len(bs), nil
}
