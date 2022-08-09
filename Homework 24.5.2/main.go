package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sortBabelArray(array ...int) []int {
	fmt.Println("Исходный массив:", array)
	flag := false

OuterLoop:
	for i := 0; i < len(array)-1; i++ {
		flag = false
		for j := 1; j < len(array); j++ {
			if array[j] > array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
				flag = true
			}
		}
		if flag == false {
			break OuterLoop
		}
	}
	fmt.Println("Отсортированный массив:", array)
	return array

}

func getTestArray(size int) []int {
	result := make([]int, size)
	for i := 0; i < len(result); i++ {
		result[i] = getRandomNumber()
	}
	return result
}

func getRandomNumber() int {
	return rand.Intn(10)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	massSort := func(args ...int) []int { return sortBabelArray(args...) }
	massSort(getTestArray(getRandomNumber())...)
	massSort()
}
