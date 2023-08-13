package main

import (
	"fmt"
)

func main() {

	fmt.Println(("Array and Slices "))

	var array1 [5]int
	array1[0] = 2
	array1[1] = 4
	array1[2] = 6
	array1[3] = 8
	array1[4] = 10

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"} //inicializando o array com valores
	array3 := [...]int{1, 2, 3, 4, 5} ///pouco fexivel
    fmt.Println("isso e um array")
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19} //inicializando o slice com valores
	slice = append(slice, 20)                             //adicionando um valor ao slice
	fmt.Println("isso e um slice")
	fmt.Println(slice)
    slice2 := array2[1:3] //pegando um pedaço do array
	fmt.Println(slice2)

	//array internos

	fmt.Println("isso e um array interno")
	slice3 := make([]float32, 10, 11) //criando um slice interno
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 5)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 10)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 15)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println(slice3)
	

}
