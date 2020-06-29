package main

import "fmt"

func main() {
	a := 3
	b := 2
	fmt.Println("soma =>", a+b)
	fmt.Println("subtração =>", a-b)
	fmt.Println("divisão=>", a*b)
	fmt.Println("multiplicação =>", a/b)

	//bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("OU", a|b)
	fmt.Println("XOR =>", a^b)

}
