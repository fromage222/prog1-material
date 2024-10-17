package main

import (
	"fmt"
	"prog1-material/code-vorlesung/greet"
)

func main() {
	m1 := greet.Greet("Kurs")
	fmt.Println(m1)

	m2 := greet.Greet("DHBW")
	fmt.Println(m2)
}
