package main

import (
	"flag"
	"log"
	"strings"
)

// Функция принимает два аргумента типа string.
// Проверяет вхождение второго аргумента в первый и возращает результат типа bool.
func checkSubstrInStr(str, substr string) bool {
	if strings.Index(str, substr) > -1 {
		return true
	}
	return false
}

//  Функция считывает значения флагов с ввода в консоли, если флаги не используются принмает значение по умолчанию.
//  Возвращает два значения типа string
func getData() (textInput string, textSearch string) {
	flag.StringVar(&textInput, "str", "default value", "input some text")
	flag.StringVar(&textSearch, "substr", "default value", "input some subtext")
	flag.Parse()
	return textInput, textSearch
}

// Главная функция, не принимает никаких значений.
// Получает значения типи string из функции getData и передает их в функцию checkSubstrInStr.
func main() {
	str, substr := getData()
	log.Printf("Исходная строка: %v", str)
	log.Printf("Искомая строка: %v", substr)
	log.Println("Результат:", checkSubstrInStr(str, substr))
}
