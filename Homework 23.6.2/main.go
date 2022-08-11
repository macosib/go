package main

import (
	"errors"
	"log"
	"strings"
)

// parseTest  - Функция, которая на вход принимает массив типа string и массив типа rune.
// Возвращает двумерный массив типа int и состояние ошибки типа error.
// Проверяет вхождение каждого символа типа rune из chars в каждую строку массива sentences,
// и возвращает массив [i][j]int, где i индекс текущей строки в sentences,
// а j индекс вхождения руны в текущую строку с индексом i.
func parseTest(sentences []string, chars []rune) ([][]int, error) {
	if len(sentences) == 0 || len(chars) == 0 {
		return nil, errors.New("Массивы не должны быть пустыми!")
	}
	result := make([][]int, 0, 0)
	for _, char := range sentences {
		temp := make([]int, 0, 0)
		for _, run := range chars {
			indexLast := strings.LastIndexAny(strings.ToUpper(char), strings.ToUpper(string(run)))
			if indexLast > -1 {
				log.Printf("Rune: %v last position - %v in %v", string(run), indexLast, char)
				temp = append(temp, indexLast)
			}
		}
		if len(temp) > 0 {
			result = append(result, temp)
		}

	}
	return result, nil
}

// checkError - Фукция принимает на вход состояние ошибки error.
// Проверяет наличие ошибки, если ошибка на равна nil, то программа завершается.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

// main - Основная функция программы, выводит в консоль результат выполнения функции parseTest
func main() {
	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}
	res, err := parseTest(sentences[:], chars[:])
	checkError(err)
	log.Printf("Результат равен: %v", res)

}
