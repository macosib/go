package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func mergeArray(arrayFirst []int, arraySecond []int) []int {
	var i, j int

	size := len(arrayFirst) + len(arraySecond)
	result := make([]int, size)

	for k := 0; k < size; k++ {
		switch {
		case i < len(arrayFirst) && j < len(arraySecond):
			if arrayFirst[i] <= arraySecond[j] {
				result[k] = arrayFirst[i]
				i++
			} else {
				result[k] = arraySecond[j]
				j++
			}
		case i < len(arrayFirst):
			result[k] = arrayFirst[i]
			i++
		case j < len(arraySecond):
			result[k] = arraySecond[j]
			j++
		}

	}
	return result

}

func getTestArray(size int) []int {
	result := make([]int, size)
	for i := 0; i < len(result); i++ {
		result[i] = getRandomNumber()
	}
	sort.Ints(result)
	return result
}

func getRandomNumber() int {
	return rand.Intn(100)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(mergeArray(getTestArray(getRandomNumber())[:], getTestArray(getRandomNumber())[:]))

}
