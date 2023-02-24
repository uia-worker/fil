// Guerrieri, A. (2019). Hands-On System Programming with Go. Packt Publishing Ltd.
package main

import (
	"fmt"
	"os"
  "io"
	"bufio"
)
func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify a file")
		return
	}

	fd, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer fd.Close() // we ensure close to avoid leaks

	r := bufio.NewReader(fd)

	var rowCount int
	// Denne sløyfen går inntil feil er nil
	for err == nil {
		var enlinje []byte
    for toolong := true; err ==  nil && toolong; {
			enlinje, toolong, err = r.ReadLine()
			if err == nil {
				fmt.Print(string(enlinje))
			}
		}
		// Hver gang toolong er usant, en linje er komplett avlest
		if err == nil {
			fmt.Println()
			rowCount++
		}
	}
	// Hvis det oppstår en annen feil enn at det er slutt på filen
	if err != nil && err != io.EOF {
		fmt.Println("\nError:", err)
		return
	}
	fmt.Println("\nRow count:", rowCount)


}
