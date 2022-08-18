package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

//getNumber - функция генирирует случайное число и передает его в канал chan int.
func getNumber() chan int {
	intChan := make(chan int)
	go func() {
		time.Sleep(1000)
		intChan <- rand.Intn(10)
	}()
	return intChan
}

//getNumber - функция принимает число из канала chan int, возводит во вторую степень и передает его в канал chan int.
func powNumber(c chan int) chan int {
	intChan := make(chan int)
	go func() {
		number := <-c
		intChan <- number * number
	}()
	return intChan
}

//getNumber - Основная функция программы. Выводит в консоль результат выполднения функции powNumber(getNumber())), пока
// не будет получен сигнал завершения.
func main() {
	var wg sync.WaitGroup
	signalChanel := make(chan os.Signal, 1)

	signal.Notify(signalChanel, os.Interrupt, os.Kill)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-signalChanel:
				fmt.Println("Выхожу из программы!")
				return
			default:
				fmt.Println("Результат:", <-powNumber(getNumber()))
			}
		}

	}()
	wg.Wait()
}
