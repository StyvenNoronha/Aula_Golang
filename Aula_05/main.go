package main

import ("fmt")

func main(){
	soma := 1+1
	sub := 1-1
	mul := 1*1
	div := 10/2
	resDiv := 10/3

	fmt.Println(soma,sub, mul, div, resDiv)

	verdade, falso := true, false

	fmt.Println(verdade && !falso)

   var num int = 1 
   var texto string

   if num > 5 {
	texto = "texto maior que 5"
   } else {
	texto = "menor que 5"
   }

   fmt.Println(texto)


}