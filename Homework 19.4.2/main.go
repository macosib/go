package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sortBabelArray(arrayFirst []int) []int {
	fmt.Println("Исходный массив:", arrayFirst)
	flag := false

OuterLoop:
	for i := 0; i < len(arrayFirst)-1; i++ {
		flag = false
		for j := 1; j < len(arrayFirst); j++ {
			if arrayFirst[j] < arrayFirst[j-1] {
				arrayFirst[j], arrayFirst[j-1] = arrayFirst[j-1], arrayFirst[j]
				flag = true
			}
		}
		if flag == false {
			break OuterLoop
		}
	}
	fmt.Println("Отсортированный массив:", arrayFirst)
	return arrayFirst

}

func getTestArray(size int) []int {
	result := make([]int, size)
	for i := 0; i < len(result); i++ {
		result[i] = getRandomNumber()
	}
	return result
}

func getRandomNumber() int {
	return rand.Intn(100)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	sortBabelArray(getTestArray(getRandomNumber()))

}
