package main

import (
	"fmt"

	"github.com/AbdulahadAbduqahhorov/bootcamp/bigInt/bigint"
)

func main() {

	a, err := bigint.NewInt("12")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("0")
	if err != nil {
		panic(err)
	}

	err = a.Set("67687879789")
	if err != nil {
		panic(err)
	}
	err = b.Set("34324")
	if err != nil {
		panic(err)
	}

	c := bigint.Add(a, b)
	d := bigint.Sub(a, b)
	e := bigint.Multiply(a, b)
	d, f := bigint.ModAndDivision(a, b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("Quotient: %v Remainder: %v\n",d.Value,f.Value)
	fmt.Println(d.Value, f.Value)
}
