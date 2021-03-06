package main

import (
	"fmt"
	"math"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

//инициализирую структуру с мьютексом и переменной для суммы квадратов чисел
type sum struct {
	sync.Mutex    // Синхронизирует доступ к данным путем явной блокировки
	sqSumNumber int
}

func main() {
	numbers := [...]int{2, 4, 6, 8, 10}

	//=====================WaitGroup========================

	//WaitGroup для контроля группы горутин
	wg := &sync.WaitGroup{}
	//счетчик количества горутин
	wg.Add(len(numbers))
	//инициализирую переменную со ссылкой на структуру с мьютексом
	st := &sum{}

	//в цикле перебираю массив чисел, передаю в функцию для возведения в квадрат и последующего вычисления их суммы
	for _, numb := range numbers {
		//запускаю анонимную функцию в горутине
		//так как го рутины выполняются в непредсказуемом порядке (0,4,4,2,3),
		//передадим анонимной функции параметры (n int)
		go func(n int) {
			//блокирую доступ к полям структуры для остальных горутин
			st.Lock()
			//вычисляю квадрат числа
			sqN := sqWG(n, wg)
			//записываю результат в поле заблокированной структуры
			st.sqSumNumber += sqN
			//разблокирую доступ для остальных горутин
			st.Unlock()
		}(numb)
	}
	//Wait ждет пока значение счетчика не станет 0
	wg.Wait()
	fmt.Println("\"WaitGroup\"\n Сумма квадратов чисел: ", st.sqSumNumber)


	//=====================Channel========================

	//синхронизируем горутины с помощью буферизированного канала ch
	ch := make(chan int, len(numbers))
	//переменная для записи результата
	var sumCh int

	//в цикле перебираю массив чисел и вычисляю их квадрат
	for _, numb := range numbers {
		//анонимную функцию запускаю в горутине, функции передаю параметр, чтобы горутины запускались по порядку
		go func(n int) {
			//вычисляю квадрат числа в отдельной функции
			sqCh(n, ch)
		}(numb)
	}

	//в цикле читаю с канала значения квадратов и суммирую их в переменной
	for i := 0; i < len(numbers); i++ {
		sumCh += <-ch
	}
	//закрываю канал
	defer close(ch)

	fmt.Println("\"Chanel\"\n Сумма квадратов чисел: ", sumCh)
}

func sqWG(numb int, wg *sync.WaitGroup) int {
	//сигнализирует, что горутина отработала
	defer wg.Done()
	//вычисляю квадрат числа используя math.Pow()
	sq := math.Pow(float64(numb), 2)
	return int(sq)
}

func sqCh(numb int, ch chan<- int) {
	//вычисляю квадрат числа используя math.Pow()
	sq := math.Pow(float64(numb), 2)
	//пишу в канал
	ch <- int(sq)

}
