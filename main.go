package main

import (
	"fmt"
	"strconv"

	"github.com/Primaginary/codenote/pakje"
)

func main() {
	fmt.Println(pakje.PrintVar())

	// Arrays
	fmt.Println(" **** ARRAYS ******")
	fmt.Println(" An array is predefined length and cannot be appended. Capacity and length are the same. ")
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

	for x := 0; x < len(tabel3); x++ {
		fmt.Println(tabel3[x])
	}

	// Slices
	fmt.Println(" **** SLICES ******")
	fmt.Println(" Slices are always based on arrays yet can be appended. The 'underlying' array will change (but not be appended).")
	tabel4 := [5]string{"aap", "noot", "mies", "toon", "bank"}
	var slyce []string = tabel4[:]
	fmt.Println(slyce)
	slyce = tabel4[1:3]
	fmt.Println(slyce)

	slyce = append(slyce, "ratel") // Slice appended, original array not; a new underlying array is created.
	fmt.Println(slyce)
	fmt.Println(len(slyce))
	fmt.Println(tabel4)
	fmt.Println(len(tabel4))

	slyce[2] = "nieuw-ratel" // Slice changed, array too
	fmt.Println(slyce)
	fmt.Println(tabel4)

	for p := 1; p < 100; p++ {
		slyce = append(slyce, strconv.Itoa(p))
	}
	fmt.Println(slyce)

	slyce[2] = "nieuw-ratel-twee" // Slice changed, but array no longer as connection was lost when slice appended array size
	fmt.Println(slyce)
	fmt.Println(cap(slyce))
	fmt.Println(len(slyce))
	fmt.Println(tabel4)

}
