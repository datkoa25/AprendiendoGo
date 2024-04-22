package main

import (
	"fmt"
)

func main() {
	fmt.Printf("2 == 2 && 3 == 3\t%v\n", 2 == 2 && 3 == 3)
	fmt.Printf("true && false\t\t%v\n", true && false)
	fmt.Printf("5 == 5 || 3 == 3\t%v\n", 5 == 5 || 3 == 3)
	fmt.Printf("true || false\t\t%v\n", true || false)
	fmt.Printf("true\t\t\t\t%v\n", true)
}
