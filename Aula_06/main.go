package main

import (
	"fmt"
)

type Carro struct {
	nome string
	ano  int
	cor  string
}

func main() {
	fmt.Print("Arquivo struct")
	var car Carro
	car.nome = "uno"
	car.ano = 98
	car.cor = "azul"

	car1 := Carro{"Fusca", 99, "preto"}
	car2 := Carro{ano: 99, cor: "preto"}

	fmt.Println(car)
	fmt.Println(car1)
	fmt.Println(car2)
}
