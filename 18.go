package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	count int64
}

func main() {
	//инициализирую несколько структур для использования в конкурентной среде
	x := &counter{}
	y := &counter{}
	mu := &sync.Mutex{}

	fmt.Println("Значение счетчика да запуска горутин: ", x.count)

	//запускаю тысячу горутин с использования атомик
	for i := 0; i < 1000; i++ {
		go atomCount(x)
	}

	//запускаю тысячу горутин с использованием mutex
	for i := 0; i < 1000; i++ {
		go mutCount(y, mu)
	}

	//time.Sleep(time.Second)
	fmt.Printf("--------\nЗначение счетчика после отработки 1000 горутин\n")
	fmt.Println("С использованием atomic: ", x.count)
	fmt.Println("С использованием mutex: ", y.count)
}

//Мьютекс достаточно тяжеловесная конструкция и для данного примера лучше подходит атомик.
func atomCount(con *counter) {
	atomic.AddInt64(&con.count, 1)
}

func mutCount(con *counter, mu *sync.Mutex) {
	mu.Lock()   //блокирую доступ к структуру для других го рутин
	con.count++ //увеличиваю значение счетчика на 1
	mu.Unlock() //разблокирую структуру для других го рутин
}
