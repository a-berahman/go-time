package main

import (
	"fmt"

	"github.com/a-berahman/go-time/01_solid/03_liskov/adder"
)

type Adder interface {
	Add(int, int) int
}

func PrintSum(a, b int, adder Adder) {
	fmt.Printf("%d + %d = %d", a, b, adder.Add(a, b))
}

func main() {
	PrintSum(1, 2, adder.Int{}) // prints: "1 + 2 = 3"

	// This line will trigger a compile-time error:
	//  cannot use adder.Double literal (type adder.Double) as type Adder
	//  in argument to PrintSum: adder.Double does not implement Adder
	//  (wrong type for Add method)
	//      have Add(float64, float64) float64
	//      want Add(int, int) int
	//PrintSum(1, 2, adder.Double{})
}
