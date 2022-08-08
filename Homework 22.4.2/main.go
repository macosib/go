package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

func getArray(number int) []int {
	array := make([]int, number)
	for i := 0; i < len(array); i++ {
		array[i] = getRandomNumber()
	}
	sort.Ints(array)
	return array
}

func getRandomNumber() int {
	return rand.Intn(12)
}

func getFirstIndex(array []int, number int) (index int) {
	min := 0
	max := len(array) - 1
	index = -1
	for max >= min {
		middle := (min + max) / 2
		if array[middle] == number {
			index = middle
			break
		} else if middle > number {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	for index > 0 && number == array[index-1] {
		index--
	}
	log.Printf("Исходный массив: %v", array)
	log.Printf("Индекс первого вхождения числа %v равен: %v", number, index)
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	getFirstIndex(getArray(getRandomNumber()), getRandomNumber())
}
