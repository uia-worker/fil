// Guerrieri, A. (2019). Hands-On System Programming with Go. Packt Publishing Ltd.
package main

import (
	"fmt"
	"os"
  "io"
)
func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify a file")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close() // we ensure close to avoid leaks

	var (
		b = make([]byte, 16)
	)

	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			//fmt.Println(string(b[:n])) // only print what's been read
			fmt.Printf("%x", b[:n])
		}
	}

	if err != nil && err != io.EOF { // we expect an EOF
		fmt.Println("\n\nError:", err)
	}

}
