package main

import (
	"fmt"

	"github.com/AbdulahadAbduqahhorov/bootcamp/bigInt/bigint"
)

func main() {

	a, err := bigint.NewInt("42342534636475898890")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("34534546457657")
	if err != nil {
		panic(err)
	}

	// err = a.Set("34324")
	// if err != nil {
	// 	panic(err)
	// }
	// c := bigint.Add(a, b)
	// d := bigint.Sub(a, b)
	// e := bigint.Multiply(a, b)
	f := bigint.Mod(a, b)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	fmt.Println(f)
}
