package main

import (
	"log"
	"math/rand"
	"time"
)

// sortBabelArray - Функция принимает любое количество аргументов типа int (array)
// Сортирует значения (пузырьковый метод) в массиве int по убыванию.
func sortBabelArray(array ...int) {
	log.Println("Исходный массив:", array)

OuterLoop:
	for i := 0; i < len(array)-1; i++ {
		flag := false
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
	log.Println("Отсортированный массив:", array)
}

// getTestArray - Функция принимает аргумент типа int
// и вовзращает массив случайных значений типа int в диапазоне от 0 до 100.
// Длина массива равна значения получаемого аргумента.
func getTestArray(size int) []int {
	result := make([]int, size)
	for i := 0; i < len(result); i++ {
		result[i] = getRandomNumber()
	}
	return result
}

// getRandomNumber - Функция не принимает никаких аргументов
// и вовзращает случайное число типа int в диапазоне от 0 до 100.
func getRandomNumber() int {
	return rand.Intn(100)
}

// main - Основная функция программы, не принимает никаких аргументов.
// Устанавливает начальную точку отсчета для модуля rand и выводит в консоль результат выполнения анонимной функции,
// которая принимает любое количество аргументов типа int
// и выводит в консоль результат отсортированный массив в обратном порядке.
func main() {
	rand.Seed(time.Now().UnixNano())
	func(args ...int) { sortBabelArray(args...) }(getTestArray(getRandomNumber())...)
}
