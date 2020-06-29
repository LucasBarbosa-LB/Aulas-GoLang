package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv32, comprarSorvete, comprarTv50
}

func main() {

	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("TV50: %t, Tv32 %t, Sorvete: %t Sorvete: %t ", tv50, tv32, sorvete, !sorvete)
}
