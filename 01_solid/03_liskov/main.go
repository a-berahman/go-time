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
	PrintSum(2, 3, adder.Int{})
	PrintSum(2, 3, adder.Double{})
}
