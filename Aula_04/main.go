package main

import("fmt")


func soma(n1 int8, n2 int8) int8{
	return n1 + n2 
}

func calculo(n1, n2 int8)(int8, int8, int8, int8){	
	somar := n1 + n2
	sub := n1 - n2
	mul := n1 * n2
	div := n1 / n2
	return somar, sub, mul, div
}

func main(){
	soma := soma(10,2)

	fmt.Println(soma)

	var f = func ( txt string )  {
		fmt.Println("nada")
		fmt.Println(txt)
	}
	f("texto")


	somar,_,_,_ := calculo(10, 2)
	//somar, sub, mul, div := calculo(10, 2)
	//fmt.Println(somar, sub, mul, div)
	fmt.Println(somar)
}