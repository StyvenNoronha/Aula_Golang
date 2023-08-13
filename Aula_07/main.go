package main

import ("fmt")

type pessoa struct{
	 nome string
	 sobrenome string
	 idade int
}

type dentista struct{
	pessoa
	especialidade string
	salario float64
}



func main(){
 fmt.Println(("Herança SQN"))
 p1 := pessoa{"João", "Silva", 30}
 fmt.Println(p1)
 p2 := dentista{p1, "Ortodontia", 10000.00}
 fmt.Println(p2)


}