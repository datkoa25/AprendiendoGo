package main

import (
	"fmt"
)

var x int = 1
var y int = 8

func main() {
	if x > y {
		fmt.Println(x, "es mayor que ", y)
		fmt.Println(x >= y) //Operador de comparación - Logica booleana
	} else {
		fmt.Println(x, "es menor que ", y)
		fmt.Println(x <= y) //Operador de comparación - Logica booleana
	}
}
