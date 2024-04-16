package main

import (
	"fmt"
)

const a = 200

// %d para decimal, %b para binario, %x para hexadecimal, %#x para hexadecimal con el 0 previo
func main() {

	fmt.Printf("%d\n%b\n%#x", a, a, a)

}
