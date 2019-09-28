package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Orange", "Plum", "Pomegranate", "Grape"}

	// some := fruits[1:3]
	// some = append([]string{}, some...)
	// some = append(some, "***Banana***")

	some := fruits[ /*start*/ 1: /*len*/ 3: /*cap*/ 3]
	fmt.Printf("fruits => %v\n", fruits)
	fmt.Printf("some => %v\n", some)
	fmt.Printf("cap(some) => %v\n", cap(some))
	fmt.Printf("len(some) => %v\n", len(some))

	some = append(some, "***1***")

	fmt.Printf("fruits => %v\n", fruits)
	fmt.Printf("some => %v\n", some)
	fmt.Printf("cap(some) => %v\n", cap(some))
	fmt.Printf("len(some) => %v\n", len(some))
}
