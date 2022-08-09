package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

// Функция принимает аргумент типа int и вовзращает массив случайных значений типа int в диапазоне от 0 до 100.
// Длина массива равна значения получаемого аргумента.
func getArray(number int) []int {
	array := make([]int, number)
	for i := 0; i < len(array); i++ {
		array[i] = getRandomNumber()
	}
	sort.Ints(array)
	return array
}

// Функция не принимает никаких аргументов и вовзращает случайное число типа int в диапазоне от 0 до 100.
func getRandomNumber() int {
	return rand.Intn(100)
}

// Функция принимает следующие аргументы: массив значений типа int (array), число типа int (number)
// и вовзращает индекс типа int первого вхождения числа number в массив array.
// В случае если аргумент number входит в массив array, будет возращено значение -1
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

// Основная функция программы, не принимает никаких аргументов.
// Устанавливает начальную точку отсчета для модуля rand и выводит в консоль результат выполнения функции getFirstIndex.
func main() {
	rand.Seed(time.Now().UnixNano())
	getFirstIndex(getArray(getRandomNumber()), getRandomNumber())
}
