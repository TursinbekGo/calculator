package main

import (
	"fmt"

	"module_name/calc"
)

func main() {
	fmt.Println(calc.Add(50, 10))
	fmt.Println(calc.Mult(50, 10))
	fmt.Println(calc.Div(50, 10))
}
