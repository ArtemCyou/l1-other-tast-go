package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main()  {
	str := "snow dog sun" //строка для изменения
	fmt.Printf("Передаем функции слово \"%s\" и получаем: %s", str, reversStr(str))
}


func reversStr(str string) string {
	temp := strings.Split(str, " ") //создаю из слова слайс строк
	//переворачиваю слайс строк
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return strings.Join(temp, " ") //возвращаю слайс в строку
}