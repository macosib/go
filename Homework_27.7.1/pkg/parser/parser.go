package parser

import (
	"Homework_27.7.1/pkg/validator"
	"bufio"
	"fmt"
	"log"
	"os"
)

//InputNewStudent - Считывет с консоли информацию о студенте, валидирует данные и возвращает в формате:
// name string, age int, grade int, bool, err error. Если данные не считаны, "", 0, 0,  false и ошибку.
// Если данные считаны, но не прошли валидацию возвращает "", 0, 0, true и ошибку для перезапуска функции.
func InputNewStudent() (string, int, int, bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Добавьте нового студента - Имя Возраст Оценка: ")
	student, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return "", 0, 0, false, err
	}
	name, age, grade, err := validator.ValidInputData(student)
	if err != nil {
		log.Println(err)
		return "", 0, 0, true, err
	}
	return name, age, grade, true, nil
}
