package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//checkArgs - Функция считывает аргументы с командной строки и передает их в виде []string.
// Если аргументов нет, то возвращается ошибка. Если аргументов больше 3, то возвращается только []string, где len == 3.
func checkArgs() ([]string, error) {
	if len(os.Args[1:]) == 0 {
		return nil, errors.New("Необходимо ввести хотя бы одно имя файла")
	} else if len(os.Args[1:]) > 3 {
		return os.Args[1:4], nil
	}
	return os.Args[1:], nil
}

//getData - Функция принимает на вход имя файла в формате string, считывает данные и возвращает их в виде []byte.
func getData(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)
	checkError(err)
	return content
}

//joinContent - Функция принимает на вход два массива формата []byte, преобразует их в string и возращает string.
func joinContent(fileFirst []byte, fileSecond []byte) string {
	return strings.Join([]string{string(fileFirst), string(fileSecond)}, "")
}

// getContent - Функция принимает на вход массив []string, где указаны имена файлов.
// Функция читает содержимое файлов и в зависимости от длины []string выводит данные в консоль
// или записывает в отдельный файл.
func getContent(fileNames []string) {
	var content string
	switch {
	case len(fileNames) == 1:
		log.Println(string(getData(fileNames[0])))
	case len(fileNames) == 2:
		content = joinContent(getData(fileNames[0]), getData(fileNames[1]))
		log.Println(content)
	case len(fileNames) == 3:
		content = joinContent(getData(fileNames[0]), getData(fileNames[1]))
		err := ioutil.WriteFile(fileNames[2], []byte(content), 0666)
		checkError(err)
	}
}

//checkError - Функция на вход принимает состояние ошибки, если ошибка есть, то программа завершается.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//main - Основная функция программы. Считывает имена файлов с командной строки и выводит содержимое этих файлов
// в консоль или записывает в отдельный файл.
func main() {
	files, err := checkArgs()
	checkError(err)
	getContent(files)
}
