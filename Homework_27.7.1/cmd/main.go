package main

import (
	"Homework_27.7.1/pkg/parser"
	"Homework_27.7.1/pkg/storage"
	"Homework_27.7.1/pkg/student"
	"log"
)

//main - Основная функция программы. Считывет с консоли информацию о студенте и добавляет его в хранилище,
// если информация введена в соответствии с шаблоном. При получении сигнала EOF, выводит имена всех студенотов.
// и завершает работу
func main() {
	studentsRepo := storage.CreateStudentsRepo()
	for {
		name, age, grade, emptyData, err := parser.InputNewStudent()
		if !emptyData {
			break
		} else if err != nil {
			continue
		}
		status := studentsRepo.Put(student.CreateNewStudent(name, age, grade))
		if status != nil {
			log.Println(err)
			continue
		}
	}
	storage.ShowStudents(&studentsRepo)
}
