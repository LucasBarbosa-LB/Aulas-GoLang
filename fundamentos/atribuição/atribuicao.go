package main

import "fmt"

func main() {
	//tipos de atribuição
	i := 3
	fmt.Println(i)
	i += 3
	fmt.Println(i)
	i -= 3
	fmt.Println(i)
	i /= 3
	fmt.Println(i)
	i *= 3
	fmt.Println(i)

	//atribuir valores a varias variaveis
	x, y := 1, 3
	fmt.Println(x, y)
}
