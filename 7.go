package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//Реализовать конкурентную запись данных в map.

func main() {
	//==============Mutex==================//
	mutexMap()
	//==============sync.Map==================//
	syncMap()
}

func mutexMap() {
	myMap := make(map[int]int)
	mu := &sync.Mutex{} //инициализирую мьютекс для блокировки горутин во время записи в карту

	wg := &sync.WaitGroup{} //инициализирую waitgroup чтобы отработала группа горутин

	for i := 1; i <= 5; i++ {
		wg.Add(1) //увеличиваю счетчик группы горутин в соответствии с их количеством
		go func(i int) { //функции передаю значение i чтобы проитерировать от 1 до 5
			mu.Lock()             //блокирую доступ к карте во время записи в горутине
			myMap[i] = rand.Int() //пишу случайное число
			mu.Unlock()           //разблокирую доступ к карте для других горутин
			wg.Done()             //уменьшаю счетчик группы горутин
		}(i)
	}
	wg.Wait() //блокируем основную горутину, ждем пока отработают остальные горутины
	fmt.Println("Конкурентная запись с помощью mutex:")

	for k := range myMap { //перебираю и вывожу значения карты
		fmt.Printf("key: %d, val: %d\n", k, myMap[k])
	}
}

func syncMap() {
	myMap := &sync.Map{} //инициализирую карту

	wg := &sync.WaitGroup{} //инициализирую waitgroup чтобы отработала группа горутин

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			myMap.Store(i, rand.Int()) //записываю данные в карту происходит lock и unlock
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Конкурентная запись с помощью sync.Map:")

	//перебираю и вывожу значения карты
	for i := 1; i <= 5; i++ {
		val, _ := myMap.Load(i)
		fmt.Printf("key: %d, val: %d\n", i, val)
	}
}
