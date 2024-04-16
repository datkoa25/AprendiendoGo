package main

import (
	"fmt"
)

//Go no tiene ciclos "While", solo tenemos ciclos de tipo For

func main() {
	// for init; condition; post {}
	for i := 0; i <= 100; i++ {

		fmt.Println(i)
	}

}
