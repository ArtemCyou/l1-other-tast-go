package main

import (
	"context"
	"fmt"

	_ "runtime"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func main() {
	//=============Остановка горутины=================
	//contextStop()
	//channelStop()
	timerStop()
	//=============Блокировка горутины=================
	//mutexBlock()
	//waitGroupBlock()

}

func channelStop() {
	stop := make(chan struct{}) //инициализирую канал со структурой, так как она ничего не весит

	go func() { //запускаю горутины
		for {   //в цикле слушаю два канала
			select {
			case <-time.Tick(500 * time.Millisecond): //каждую секунду вывожу сообщение
				fmt.Println("Канал: Горутина активна")
			case <-stop: //жду сигнал о завершении от другой горутины
				return
			}
		}
	}()

	go func() { //запускаю горутину в которой через 3 секунды даю сигнал в канал stop
		time.Sleep(2 * time.Second)
		stop <- struct{}{}
		close(stop)
	}()

	time.Sleep(3 * time.Second) //даю горутинам время на отработку
	fmt.Println("Канал: Горутина остановлена")
	fmt.Println("Выход программы")
}

func contextStop() {
	//создаю контекст для контроля горутины с заданным таймаутом
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	goDead := make(chan bool) //инициализирую канал для сигнала остановки горутины

	go func() { //запускаю горутину
		for {   //в цикле жду сообщения из двух каналов
			select {
			case <-time.Tick(time.Second): //каждую секунду вывожу сообщение
				fmt.Println("Контекст: горутина активна")
			case <-ctx.Done(): //получаю сообщение от контекста, выхожу из цикла
				fmt.Println("Контекст: горутина остановлена")
				goDead <- true //сигнализирую в канал об остановке горутины
				return

			}
		}
	}()

	<-goDead //ожидаем сигнал об остановке горутины
	fmt.Println("Выход программы")
}

func timerStop() {
	//создаю таймер для контроля горутины с заданным временем
	timer := time.NewTimer(4 * time.Second)
	timeDeath := make(chan struct{}) //канал для сигнала об остановки горутины

	go func() {
		for { //в цикле жду сообщения из двух каналов
			select {
			case <-time.Tick(time.Second):
				fmt.Println("Таймер: Горутина активна")
			case <-timer.C: //как только в таймере закончится время получим сигнал
				fmt.Println("Таймер: Горутина остановлена")
				timeDeath <- struct{}{} //пишем в канал
				return
			}
		}
	}()

	<-timeDeath //ожидаем сигнал об остановке горутины
	fmt.Println("Выход программы")
}

func waitGroupBlock() {
	wg := &sync.WaitGroup{}

	go func() { //запускаю горутину
		for {
			fmt.Println("WaitGroup: Горутина активна")
			time.Sleep(1 * time.Second) //вывожу сообщение в цикле раз в секунду
			wg.Wait()                   //блокирует горутину если счетчик больше 0
		}
	}()

	time.Sleep(3 * time.Second) //ждем 3 секунды

	wg.Add(1) //инкрементируем счетчик WaitGroup
	fmt.Println("WaitGroup: Горутина заблокирована")

	time.Sleep(time.Second)
	fmt.Println("Выход приложения")

}

func mutexBlock() {
	mu := &sync.Mutex{} //инициализирую мьютекс

	go func() { //запускаю горутину
		for {
			mu.Lock() //блокирую горутину, вывожу сообщение
			fmt.Println("Mutex: Горутина активна")
			mu.Unlock()                 //разблокирую горутину
			time.Sleep(1 * time.Second) //вывожу сообщение в цикле раз в секунду
		}
	}()

	time.Sleep(3 * time.Second) //выделяю горутине 3 секунды

	mu.Lock() //блокирую горутину
	fmt.Println("Mutex: Горутина заблокирована")

	time.Sleep(time.Second)
	fmt.Println("Выход приложения")
}
