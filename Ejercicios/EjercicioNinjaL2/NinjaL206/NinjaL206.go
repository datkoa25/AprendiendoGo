package main

import (
	"fmt"
)

const (
	año = 2024
	s1  = (año + iota)
	s2  = (año + iota)
	s3  = (año + iota)
)

func main() {
	fmt.Println(año, s1, s2, s3)
}
