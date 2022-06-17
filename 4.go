package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.

func main() {

	ch := make(chan int, 1)        //инициализирую канал для главного потока
	doneWork := make(chan bool) //канал для отслеживания завершения воркеров
	var numb int                //переменная для кол-во воркеров
	defer close(ch) //закрываем каналы хоть в этом нет необходимости
	defer close(doneWork)

	// создаю контекст для работы
	ctx, cancel := context.WithCancel(context.Background())

	//запускаю цикл для ввода количества воркеров при старте
	for {
		fmt.Print("Введите количество воркеров: ")
		_, err := fmt.Scan(&numb) //сохраняю вол-во воркеров в переменную
		if err != nil {
			fmt.Println("Алло: ожидаю целое число")
		} else {
			break
		}
	}

	//запускаю заданное кол-во воркеров
	for i := 0; i < numb; i++ {
		go worker(ctx, ch, doneWork, i) //передаю контекст, канал главного потока, воркера и номер воркера
	}

	//функция graceful shutdown в отдельной горутине
	go shutdown(cancel)

	ticker := time.NewTicker(time.Second) //создаю тикер, чтобы каждую секунду отправлять данные в канал

	exit := false //переменная для выхода из цикла
	//в цикле записываю данные в главный поток и слушаю контекст
	for {
		select { //мультиплексирование для работы с несколькими каналами
		case <-ticker.C:
			ch <- time.Now().Second() //каждую секунду отправляю данные
		case <-ctx.Done():
			exit = true //выходим из цикла если получаем сигнал от контекста
			break
		}
		if exit {
			break
		}
	}

	// читаем из канала, для завершения работы запущенных воркеров
	for i := 0; i < numb; i++ {
		//для красоты в терминале, выводим сообщения поочерёдно
		func(i int) {
			fmt.Printf("Воркер #%d - завершен.\n", i)
			<-doneWork
		}(i)
	}

	log.Println("Идеальное завершение программы.")
 }

func worker(ctx context.Context, ch <-chan int, doneWork chan<- bool, index int) {
	exit := false // инициализирую переменную для выхода из цикла

	for { //в бесконечном цикле слушаем каналы
		select { //мультиплексирование для работы с несколькими каналами
		case x := <-ch:
			fmt.Printf("Воркер #%d: %d\n", index, x)
		case <-ctx.Done():
			exit = true //пишем true чтобы сработал if для выхода из цикла
			doneWork <- true
			break //выходим из select
		}
		if exit {
			break //выходим из цикла
		}
	}

}

func shutdown(cancel context.CancelFunc) {
	chCtx := make(chan os.Signal, 1)       //создаём канал для сигналов системы
	signal.Notify(chCtx, os.Interrupt)     //ждем сигнал ОС что наше приложение завершилось
	log.Printf("system call:%+v", <-chCtx) // выводим полученный сигнал в stdout

	cancel() // отменяем контекст
}
