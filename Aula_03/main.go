package main

import "fmt"

func main() {
	var num int8 = 100
	var num1 int16 = 100
	var num2 int32 = 100
	var num3 int64 = 100
	var num4 int = 100

	var num5 int8 = 2
	var num6 float32 = 2.5
	fmt.Println(num5 + int8(num6))

	fmt.Println(num)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

}