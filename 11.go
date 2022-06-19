package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	//Гуглим диаграмму Венна, чтобы понять, что от нас требуют
	subsetOne := []int{1, 8, 3, 14, 5, 9, 7, 85, 77, 10}
	subsetTwo := []int{80, 9, 8, 50, 12, 13, 14, 3}

	res := cross(subsetOne, subsetTwo)       // нахожу пересечения с помощью перебора O(n^2)
	fmt.Println("Пересечение (sort): ", res) //вывожу пересечения

	res = crossMap(subsetOne, subsetTwo)    // нахожу пересечения с помощью хэш-таблицы(мапки) O(n)
	fmt.Println("Пересечение (map): ", res) //вывожу пересечения

}

func cross(a, b []int) []int {
	crArr := make([]int, 0) //инициализирую слайс для пересечений

	for _, valA := range a { //в цикле перебираю значения первого множества
		for _, valB := range b { //в цикле перебираю значения второго множества
			if valA == valB { //сравниваю значения первого множества с каждым значением из второго множества
				crArr = append(crArr, valA) //добавляю в слайс найденные пересечения
			}
		}
	}
	return crArr
}

func crossMap(a, b []int) []int {
	crMap := make(map[int]struct{}) //инициализирую мапку с пустой структурой так как она ничего не весит
	crArr := make([]int, 0)         //инициализирую слайс для пересечений

	//заполняю мапку ключами из первого множества
	for _, valA := range a {
		crMap[valA] = struct{}{}
	}

	//сравниваю ключи мапки со значениями из второго множества, если нахожу добавляю в слайс
	for _, valB := range b {
		if _, ok := crMap[valB]; ok {
			crArr = append(crArr, valB)
		}
	}

	return crArr
}
