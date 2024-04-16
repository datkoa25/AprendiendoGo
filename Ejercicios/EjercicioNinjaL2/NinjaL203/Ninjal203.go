package main

import (
	"fmt"
)

const x int = 3
const y = 4

func main() {
	fmt.Println(x, y)
	fmt.Printf("el tipo de x es: %T\n", x)
	fmt.Printf("el tipo de y es: %T\n", y)
}
