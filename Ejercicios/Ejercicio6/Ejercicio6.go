package main

import (
	"fmt"
)

func main() {

	s1 := "Hola mundo"
	s2 := `Esta es una linea 
	partida.`

	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("El tipo de s2 es: %T\n", s2)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T", bs)

	//Unicode code point ej: U+20AC

	fmt.Println("")

	//for funciona para crear un bucle y i es igual a iterar o en otras palabras repetir
	//for i, para inicializar, i marcar el condicional y i++ para declarar el post, i++ = "i = i+1"

	for i := 0; i < len(s1); i++ {

		fmt.Printf("%#U ", s1[i])
	}

	fmt.Println("")

	for i, v := range s1 {

		fmt.Printf("En el indice %d el valor es %v\n", i, v)
	}

	fmt.Printf("Con el verbo %q inidico que se imprima el string %s", "%s", s1)

}
