package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

var arr = []int{1, 2, 3, 4, 5}

func main() {
	wg := &sync.WaitGroup{} //инициализирую WaitGroup для работы горутины
	wg.Add(1) //даю отработать конвейеру в горутине

	fmt.Println("Значение массива до работы конвейера: ",arr)
	go conveer(wg) //запускаю конвейер в горутине

	wg.Wait() //жду окончание работы
}

func conveer(wg *sync.WaitGroup) {
	convChan := make(chan int)     //инициализирую канал для записи из массива
	convChanDobl := make(chan int) //инициализирую канал для записи результата х*2
	var result = make([]int,0)

	go func() {                   //в горутине в канал пишутся числа (x) из массива arr
		for _, val := range arr { //перебираем массив и пишем в канал
			convChan <- val
		}
		close(convChan) //закрываем канал
	}()

	go func() { //в горутине читаем из канала convChan, и пишем в канал convChanDobl результат х*2
		for val := range convChan {
			convChanDobl <- val * 2
		}
		close(convChanDobl) //закрываем канал

	}()

	//благодаря WaitGroup код ниже исполнится
	for val := range convChanDobl { //читаю канал в цикле и создаю новое значение массива
		result = append(result, val)
	}

	fmt.Println("Значение массива после работы конвейера: ", result)
	wg.Done() //уменьшаю счетчик
}
