package main

import (
	"log"
	"math/rand"
	"time"
)

// getArray - Функция принимает аргумент типа int
// и вовзращает массив случайных значений типа int в диапазоне от 0 до 100.
// Длина массива равна значения получаемого аргумента.
func getArray(number int) []int {
	array := make([]int, number)
	for i := 0; i < len(array); i++ {
		array[i] = getRandomNumber()
	}
	return array
}

// getRandomNumber - Функция не принимает никаких аргументов
// и вовзращает случайное число типа int в диапазоне от 0 до 100.
func getRandomNumber() int {
	return rand.Intn(100)
}

// choiceSorted - Функция принимает массив типа int (array)
// Сортирует значения (метод вставки) в массиве int.
func choiceSorted(array []int) {
	if len(array) == 0 {
		log.Println("Массив не должен быть пустым! ")
		return
	}
	log.Println("Исходный массив: ", array)
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	log.Println("Отсортированный массив: ", array)
}

// main - Основная функция программы, не принимает никаких аргументов.
// Устанавливает начальную точку отсчета для модуля rand и выводит в консоль результат выполнения функции choiceSorted.
func main() {
	rand.Seed(time.Now().UnixNano())
	choiceSorted(getArray(getRandomNumber()))
}
