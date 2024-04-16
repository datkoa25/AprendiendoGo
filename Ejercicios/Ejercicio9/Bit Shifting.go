package main

import (
	"fmt"
)

const (
	// _ es igual blank identifier y sirve para desechar el primer número de iota

	_  = iota
	kb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {

	a := 4

	//verbo %d para expresar la variable en decimal y %b para expresar la variable en binario

	fmt.Printf("%d\t\t%b\n", a, a)

	//<< o >> sirven para un shift de los bits de posición hacia la izquierda o derecha

	b := a << 1

	fmt.Printf("%d\t\t%b\n", b, b)

	//ejercicio de memoria

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)

}
