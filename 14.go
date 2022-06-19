package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func main() {
	a := make(chan int) // создаю переменную с одним из заданных типов

	//определяю ее тип с помощью:
	typeSwith(a)
	fmt.Println("===============") //просто разделитель
	typeReflect(a)
	fmt.Println("===============")
	typeFMT(a)
}

func typeSwith(a interface{}) {
	fmt.Println("Определяем тип с помощью switch")

	switch v := a.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: boll")
	case chan int:
		fmt.Println("Тип переменной: channel int")
	case chan float64:
		fmt.Println("Тип переменной: channel float64")
	case chan string:
		fmt.Println("Тип переменной: channel string")
	case chan bool:
		fmt.Println("Тип переменной: channel bool")
	default:
		fmt.Printf("Неожиданный тип: %T", v)
	}
}

func typeReflect(a interface{}) {
	fmt.Println("Определяем тип с помощью reflect")

	v := reflect.TypeOf(a).Kind() //проверяю тип переменной
	notype := true //создаю булевую переменную чтобы активировать if если тип переменной не входит в наш список

	if v == reflect.Int {
		fmt.Println("Тип переменной: int")
		notype = false
	}
	if v == reflect.String {
		fmt.Println("Тип переменной: string")
		notype = false
	}
	if v == reflect.Bool {
		fmt.Println("Тип переменной: bool")
		notype = false
	}
	if v == reflect.Chan {
		fmt.Println("Тип переменной: channel")
		notype = false
	}
	if notype {
		fmt.Println("Неожиданный тип: ", v)
	}
}

func typeFMT(a interface{})  {
	fmt.Println("Определяем тип с помощью пакета fmt")
	fmt.Printf("Тип переменной: %T", a)

}
