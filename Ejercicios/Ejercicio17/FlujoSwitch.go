package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 == 4:
		fmt.Println("No deberia imprimir")
	case 3 == 3:
		fmt.Println("Deberia imprimir")
		//fallthorugh <- ejecuta el codigo del siguiente caso asi no sea verdadero
	case 4 == 5:
		fmt.Println("No deberia imprimir 2")
	default:
		fmt.Println("Ningun caso es verdadero")
	}

}
