package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func checkArgs() ([]string, error) {
	if len(os.Args[1:]) == 0 {
		return nil, errors.New("Необходимо ввести хотя бы одно имя файла")
	} else if len(os.Args[1:]) > 3 {
		return os.Args[1:4], nil
	}
	return os.Args[1:], nil
}

func getData(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)
	checkError(err)
	return content
}

func joinContent(fileFirst []byte, fileSecond []byte) string {
	return strings.Join([]string{string(fileFirst), string(fileSecond)}, "")
}

func getConc(fileNames []string) {
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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	files, err := checkArgs()
	checkError(err)
	getConc(files)
}
