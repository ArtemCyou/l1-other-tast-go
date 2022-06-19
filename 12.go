package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main()  {
	subset := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Собственное множество: ", newSub(subset))


}

func newSub(arr []string)   []string {
	subM:= make(map[string]struct{}) //инициализирую мапку с пустой структурой так как она ничего не весит
	subArr := make([]string, 0) //инициализирую слайс для собственного множества

	for _, val:= range arr{ //перебираю все значения слайса в цикле, чтобы добавить в мапку
		subM[val] = struct{}{}

	}

	for key := range subM{ //создаю новое множество
		subArr = append(subArr, key)
	}

	return subArr
}