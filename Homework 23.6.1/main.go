package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// getArray - Функция принимает аргумент типа int и вовзращает массив случайных значений типа int в диапазоне от 0 до 100.
// Длина массива равна значения получаемого аргумента.
func getArray(number int) []int {
	array := make([]int, number)
	for i := 0; i < len(array); i++ {
		array[i] = getRandomNumber()
	}
	return array
}

// getRandomNumber - Функция не принимает никаких аргументов и вовзращает случайное число типа int в диапазоне от 0 до 100.
func getRandomNumber() int {
	return rand.Intn(100)
}

// sortedEvenOdd - Функция принимает массив значений типа int (array)
// Возращает два массива со значениями типа int, один содержит четные значения, другой нечетные.
func sortedEvenOdd(array []int) ([]int, []int) {
	evenArray := make([]int, 0, len(array)/2)
	oddArray := make([]int, 0, len(array)/2)
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			evenArray = append(evenArray, array[i])
		} else {
			oddArray = append(oddArray, array[i])
		}
	}
	getInfoArray(array)
	getInfoArray(evenArray)
	getInfoArray(oddArray)
	return evenArray, oddArray
}

// getInfoArray - Функция принимает массив значений типа int (array).
// Выводит в консоль расширенную информацию (данные, длина, емкость) массива.
func getInfoArray(array []int) {
	log.Printf("data: %+v\n", array)
	log.Printf("len: %+v\n", len(array))
	log.Printf("cap: %+v\n", cap(array))
	fmt.Println()
}

// main - Основная функция программы, не принимает никаких аргументов.
// Устанавливает начальную точку отсчета для модуля rand и выводит в консоль результат выполнения функции sortedEvenOdd.
func main() {
	rand.Seed(time.Now().UnixNano())
	sortedEvenOdd(getArray(getRandomNumber()))
}
