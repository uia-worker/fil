package main

import (
	"fmt"
	"bytes"
)
func main() {

	var dikt = []string{
`  Frihed, hil!`,
`Et Suk vi Norges Frihed bringe
   Sang og Smiil
Skal være Sukkets Vinge.`,
`See idag sig reise kan den Arme,
   Raabe: "Jordens
   Fyrste! Nordens
   Mand har Diadem!"`,
	}

	var b = bytes.NewBuffer(make([]byte, 16))

	for i := range dikt {
		b.Reset()
		b.WriteString(dikt[i])
		fmt.Println("Length:", b.Len(), "\tCapacity:", b.Cap())
	}

}
