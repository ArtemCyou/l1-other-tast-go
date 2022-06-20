package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	a := "abcd"
	b := "abCdefAaf"
	c := "aabcd"
	d := "aAbcd"

	fmt.Println("Проверяем, что все символы в строке уникальные:")
	fmt.Println("abcd — ", unicSlise(a))
	fmt.Println("abCdefAaf — ", unicSlise(b))
	fmt.Println("aabcd — ", unicSlise(c))
	fmt.Println("aAbcd — ", unicSlise(d))
}

func unicSlise(str string) bool {
	//привожу все символы в строке к одному регистру
	strLow := strings.ToLower(str)
	//разбиваю страку на слайс
	strSl := strings.Split(strLow, "")

	//перебираю слайс в циклах и сравниваю его сам с собой
	for i := range strSl {
		for j := i + 1; j < len(strSl); j++ { //увеличиваю индекс элемента на 1
			if strSl[i] == strSl[j] {
				return false
			}
		}
	}
	return true
}
