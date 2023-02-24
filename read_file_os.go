package main

import (
	"fmt"
	"os"
  "log"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify a path.")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
    // bruker en standard logger, som kaller log.Print() med påfølgende os.Exit(1)
    // output er standard error fil
    log.Fatal(err)
	}
	os.Stdout.Write(data)
}
