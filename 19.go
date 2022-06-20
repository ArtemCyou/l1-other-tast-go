package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	str := "главрыба"
	fmt.Printf("Передаем функции слово \"%s\" и получаем: %s", str, reverseStr(str))



	// sort.Sort(sort.Reverse(sort.StringSlice(temp)))
	//fmt.Println(temp)
}

func reverseStr(str string) string {
	temp := strings.Split(str, "") //создаю из слова слайс строк
	//переворачиваю слайс строк
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return strings.Join(temp, "") //возвращаю слайс в строку
}
