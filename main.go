package main

import (
	"fmt"

	"github.com/AbdulahadAbduqahhorov/bootcamp/bigInt/bigint"
)

func main() {

	a, err := bigint.NewInt("-19")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("-20")
	if err != nil {
		panic(err)
	}

	// // err = a.Set("-23453678")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Println(a)

	// c := bigint.Add(a, b)
	d := bigint.Sub(a, b)
	// e := bigint.Multiply(a, b)
	// // f := bigint.Mod(a, b)
	// // fmt.Println(a)
	// // fmt.Println(b)
	// fmt.Println(c)
	fmt.Println(d)
	// fmt.Println(e)
	// // fmt.Println(f)
}
