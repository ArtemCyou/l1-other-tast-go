package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	//реализация функции sleep с помощью таймера
	sleepTime(1)

	fmt.Println("-----------")

	//реализация функции sleep с помощью context
	ctxSleep(2)

}

func sleepTime(secondSleep time.Duration) {
	fmt.Printf("Реализация ф-н sleep с помощью таймера, ждем: %dсек\n", secondSleep)

	//инициализирую таймер
	timer := time.NewTimer(secondSleep * time.Second)
	<-timer.C //жду сигнала из канала

	fmt.Println("Время вышло \"timer\"")
}

func ctxSleep(secondSleep time.Duration) {
	fmt.Printf("Реализация ф-н sleep с помощью context, ждем: %dсек\n", secondSleep)

	//создаю контекст и время жизни
	ctx, _ := context.WithTimeout(context.Background(), secondSleep*time.Second)
	<-ctx.Done() //жду сигнала из канала

	fmt.Println("Время вышло \"context\"")
}
