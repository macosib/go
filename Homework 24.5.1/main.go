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

func choiceSorted(array []int) {
	fmt.Println("Исходный массив: ", array)
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	fmt.Println("Отсортированный массив: ", array)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	choiceSorted(getArray(getRandomNumber()))
}
