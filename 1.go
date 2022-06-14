package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

//родительская структура имеющая свои методы
type Human struct {
	Age  int
	Name string
	//..
}

//имеет встроенную структуру Human, поэтому имеет доступ к ее методам
//встраивание методов происходит автоматически
type Action struct {
	Human
}

//метод выдает имя Human
func (h *Human) Say() {
	fmt.Println("my name - ", h.Name)
}

func main() {
	a := Human{Name: "John"}

	//инициализирую структуру Action и получаю доступ к методам и полям родительской структуры Human
	b := Action{a}
	b.Say()
}
