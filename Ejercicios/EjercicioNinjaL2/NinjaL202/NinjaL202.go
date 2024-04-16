package main

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (42 <= 50)
	c := (42 >= 40)
	d := (42 != 41)
	e := (42 < 46)
	f := (42 > 35)

	fmt.Println(a, b, c, d, e, f)

}
