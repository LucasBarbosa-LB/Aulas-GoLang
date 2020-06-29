package main

import (
	"fmt"
	"math"
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64 é inferido pelo compilador.

	//forma reduzida de criar uma variavel
	area := PI * math.Pow(raio, 2)
	fmt.Println("Á Area do circulo é igual a ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = false, true

	fmt.Println(e, f)

	g, h, i := true, false, "opaa"

	fmt.Println(g, h, i)

}
