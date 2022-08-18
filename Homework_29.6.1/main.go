package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// handlerStop - функция принмает строку в формате string и возвращает результат проверки остановки в формате bool
func handlerStop(str string) bool {
	if str == "стоп" {
		return true
	}
	return false
}

//inputNumber - функция считывает число из консоли.
func inputNumber() int {
	for {
		var input string
		log.Println("Введите число: ")
		fmt.Scan(&input)
		if handlerStop(input) {
			os.Exit(1)
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Println("Необходимо ввести число!!! ")
			continue
		}
		return number
	}
}
//getNumber - функция возвращает введенное число из консоли в канал.
func getNumber() chan int {
	intChan := make(chan int)
	go func() {
		intChan <- inputNumber()
	}()
	return intChan
}

//powNumber - функция принимает число из канала, возводит во вторую степень и возвращает результат в канал.
func powNumber(c chan int) chan int {
	intChan := make(chan int)
	go func() {
		number := <-c
		intChan <- number * number
	}()
	return intChan
}
//multNumber - функция принимает число из канала, умножает на 2 и возвращает результат в канал.
func multNumber(c chan int) chan int {
	intChan := make(chan int)
	go func() {
		number := <-c
		intChan <- number * 2
	}()
	return intChan
}

//main Основная функция программы. Считывает число с консоли выводит результат выполнения функций.
func main() {
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		number := getNumber()
		numberPow := powNumber(number)
		numbermult := multNumber(numberPow)
		fmt.Println("Результат:", <-numbermult)
		wg.Done()
	}
	wg.Wait()
}
