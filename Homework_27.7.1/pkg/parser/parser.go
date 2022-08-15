package parser

import (
	"Homework_27.7.1/pkg/validator"
	"bufio"
	"fmt"
	"log"
	"os"
)

func InputNewStudent() (string, int, int, bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Добавьте нового студента - Имя Возраст Оценка: ")
	student, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return "", 0, 0, false, err
	}
	name, age, grade, errData := parseInputData(student)
	if errData != nil {
		log.Println(errData)
		return "", 0, 0, true, errData
	}
	return name, age, grade, true, errData
}

func parseInputData(data string) (string, int, int, error) {
	name, age, grade, err := validator.ValidInputData(data)
	if err != nil {
		return "", 0, 0, err
	}
	return name, age, grade, err
}
