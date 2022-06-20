package main

import (
	"fmt"
	"sort"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

//создаю свой тип для сортировки
type myType struct {
	who   string
	place int
}

func main() {
	arrA := []int{4, 2, 3, 1, 5}
	fmt.Println("Неотсортированный массив: ", arrA)

	//реализация быстрой сортировки go
	quickSort(arrA, 0, len(arrA)-1)
	fmt.Printf("Отсортированный массив (quickSort): %v\n\n", arrA)

	//сортировка встроенными методами языка
	otherSort(arrA)
}

func quickSort(arr []int, low, high int)  {
	if low < high {

		// Находим индекс в массиве для разворота массива
		var pivot = part(arr, low, high)

		// Применяем стратегию "разделяй и властвуй"
		// для сортировки левой и правой частей массива
		// в соответствии с положением pivot

		// левая часть массива
		quickSort(arr, low, pivot)

		// правая часть массива
		quickSort(arr, pivot + 1, high)
	}
}

func part(arr []int, low, high int) int {
	// Выбираем первые элемент массива в качестве начального значения
	var pivot = arr[low]

	var i = low
	var j = high

	for i < j {

		// Увеличиваем значение индекса "i"
		//до тех пор, пока крайние левые значения не станут меньше или равны pivot
		for arr[i] <= pivot && i < high {
			i++
		}

		// Уменьшаем значение индекса "j"
		//до тех пор, пока самые правые значения не будут больше, чем значение pivot
		for arr[j] > pivot && j > low {
			j--
		}

		// Поменяем местами значения
		// in the index i and j of an array only if i is less than j
		if i < j {
			//var temp = arr[i]
			arr[i], arr[j] = arr[j], arr[i]
			//arr[j] = temp
		}
	}

	// Заменяем элемент в позиции j на нижнюю границу массива
	// pivot перемещаем в позицию j
	arr[low] = arr[j]
	arr[j] = pivot

	// Возвращаем позицию индекса pivot
	return j
}

func otherSort(arrA []int)  {
	//для сортировки использую встроенный метод sort.Ints()
	sort.Ints(arrA) //так-же можно отсортировать string, float
	fmt.Println("Отсортированный массив встроенным методом языка: ", arrA)

	fmt.Println("------Сортировка своего типа-------")

	//Если нужно отсортировать свой тип, воспользуемся методом sort.Slice
	arrB := []myType{{"Бернар Арно", 3}, {"Джефф Безос", 1}, {"Илон Маск", 2}}
	fmt.Println("Неотсортированный массив: ", arrB)

	//метод принимает массив для сортировки и функцию
	sort.Slice(arrB, func(i, j int) bool {
		return arrB[i].place < arrB[j].place
	})

	fmt.Println("Отсортированный массив встроенным методом языка: ", arrB)
}
