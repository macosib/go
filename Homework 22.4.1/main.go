package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция принимает аргумент типа int и вовзращает массив случайных значений типа int в диапазоне от 0 до 100.
// Длина массива равна значения получаемого аргумента.
func getArray(number int) []int {
	array := make([]int, number)
	for i := 0; i < len(array); i++ {
		array[i] = getRandomNumber()
	}
	return array
}

// Функция не принимает никаких аргументов и вовзращает случайное число типа int в диапазоне от 0 до 100.
func getRandomNumber() int {
	return rand.Intn(100)
}

// Функция принимает следующие аргументы: массив значений типа int (array), число типа int (number)
// и вовзращает число элементов в массиве, после переданного number.
// В случае если аргумент number отрицательный или number == array[len(array) - 1 ] возвращается 0.
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

// Основная функция программы, не принимает никаких аргументов.
// Устанавливает начальную точку отсчета для модуля rand и выводит в консоль результат выполнения функции countNumber.
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Результат:", countNumber(getArray(getRandomNumber()), getRandomNumber()))
}
