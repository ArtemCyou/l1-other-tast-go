package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a := 2
	b := 4
	fmt.Printf("Начальное значение: a= %d, b= %d\n", a, b) //

	//идиоматичный способ
	a, b = b, a
	fmt.Printf("Идиоматичный способ: a= %d, b= %d\n", a, b)

	//побитовое исключающее (ИЛИ)
	a = a ^ b // 100^10 = 110(6)
	b = a ^ b // 110^10 = 100(4)
	a = a ^ b // 110^100 = 010(2)
	fmt.Printf("Побитовое исключающее: a= %d, b= %d\n", a, b)
}
