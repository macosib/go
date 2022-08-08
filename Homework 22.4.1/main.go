package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getArray(number int) []int {
	array := make([]int, number)
	for i := 0; i < len(array); i++ {
		array[i] = getRandomNumber()
	}
	return array
}

func getRandomNumber() int {
	return rand.Intn(100)
}

func countNumber(array []int, number int) int {
	index := len(array) - 1
	for i := 0; i < len(array); i++ {
		if number == array[i] {
			index = i
			break
		}
	}
	fmt.Println("Исходный массив:", array)
	fmt.Println("Число для поиска:", number)
	if index != len(array)-1 {
		return len(array[index+1:])
	} else {
		return 0
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Результат:", countNumber(getArray(getRandomNumber()), getRandomNumber()))
}
