package main

import (
	"fmt"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.

//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}

//func main() {
//	someFunc()
//}

var justString string

func main() {
	someFunc()
	goodFunc()
}

func someFunc() {
	v := createHugeString(1 << 10) //создаем длинную строку

	//если мы используем такую реализацию то потеряем часть информации
	//результат вывода: привет привет привет привет привет привет привет прив�
	justString = v[:100]
	fmt.Printf("Некорректная реализация: %v\n", justString)
}

func goodFunc() {
	v := createHugeString(1 << 10) //создаем длинную строку

	temp := []rune(v) //преобразовываем строку в массив рун
	//возвращаем строку, теперь мы обрежем слово, но не символ так как забираем не 100 байт, а 100 рун
	justString := string(temp[:100])

	fmt.Printf("Корректная реализация: %v\n", justString)
}

func createHugeString(size uint64) (line string) {
	//в цикле в переменную 1024 раза добавим слово "привет"
	for i := 0; uint64(i) < size; i++ {
		line += fmt.Sprint("привет ") // слово занимает 13 байт (каждый символ 2 байт плюс пробел 1 байт)
	}
	return
}
