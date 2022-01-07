package main

import (
	"fmt"

	"github.com/Primaginary/codenote/pakje"
)

func main() {
	fmt.Println(pakje.PrintVar())

	// Arrays
	var tabel [4]string
	tabel[0] = "Hoppa1"
	tabel[1] = "Hoppa2"
	tabel[2] = "Hoppa3"
	tabel[3] = "Hoppa4"
	fmt.Println(tabel)

	tabel2 := [3]string{"aap", "noot", "mies"}
	fmt.Println(tabel2)

	tabel3 := [2][5]string{{"a", "A", "AA", "AAA", "AAA"}, {"b", "B", "BB", "BBB", "BBBB"}}
	fmt.Println(tabel3)
	fmt.Println(tabel3[1][3])

}
