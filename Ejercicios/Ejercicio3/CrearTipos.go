package main

import (
	"fmt"
)

var a int

type dinero int

var b dinero

func main() {

	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	/*error - Recuerda que Go es un lenguaje estatico

	a = b
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	*/

	//Como usar la conversión con sintaxis T(x) donde T es un tipo y X es una expresión que puede ser convertida en el tipo T

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
