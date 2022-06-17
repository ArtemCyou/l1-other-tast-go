package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	//инициализирую канал для записи
	ch := make(chan int)
	//создаю тикер, чтобы каждую секунду отправлять данные в канал
	ticker := time.NewTicker(time.Second)

	//использую метод WithTimeout, чтобы  функция cancel() применилась самостоятельно
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)

	//в функции читаю с канала и передаю контекст
	go read(ctx, ch)

	//последовательно пишем в канал
	for {
		select {
		case <-ticker.C:
			ch <- time.Now().Second()
		}
	}
}

func read(ctx context.Context, ch <-chan int) {
	for {
		select {
		case x := <-ch: //читаю из канала
			fmt.Println(x)
		case <-ctx.Done(): //жду сигнал
			fmt.Println(ctx.Err()) //вывожу ошибку объясняющую, почему завершилась программа
			os.Exit(0)             //завершаю программу
		}
	}
}
