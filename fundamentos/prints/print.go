package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	a := 1
	b := 1.999
	c := false
	d := "Opa"

	fmt.Printf("\n%d \n%2.f, %f \n%t \n%s", a, b, b, c, d)

}
