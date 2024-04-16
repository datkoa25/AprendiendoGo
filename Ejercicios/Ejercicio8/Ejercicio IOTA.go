package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

//iota se reinicia cada que la palabra clave "const" sea usada

const (
	d = iota
	e = iota
	f = iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
