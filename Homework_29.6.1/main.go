package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func handlerStop(str string) bool {
	if str == "стоп" {
		return true
	}
	return false
}

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

func getNumber() chan int {
	intChan := make(chan int)
	go func() {
		intChan <- inputNumber()
	}()
	return intChan
}

func powNumber(c chan int) chan int {
	intChan := make(chan int)
	go func() {
		number := <-c
		intChan <- number * number
	}()
	return intChan
}

func multNumber(c chan int) chan int {
	intChan := make(chan int)
	go func() {
		number := <-c
		intChan <- number * 2
	}()
	return intChan
}

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
