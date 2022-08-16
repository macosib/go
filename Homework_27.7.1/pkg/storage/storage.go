package storage

import (
	"Homework_27.7.1/pkg/student"
	"errors"
	"fmt"
)

type StudentsRepo map[string]*student.Student

// CreateStudentsRepo - Функция создает новый map[string]*student.Student для записи в нее новых студентов.
func CreateStudentsRepo() StudentsRepo {
	return make(map[string]*student.Student)
}

// Get - Метод StudentsRepo, который принимает имя студента в виде ключа.
// Возвращает ссылку на объект структуры student.Student и состояние ошибки.
// Если студента нет в экземпляре StudentsRepo, то возращает nil и ошибку.
func (sr StudentsRepo) Get(student string) (*student.Student, error) {
	stud, status := sr[student]
	if !status {
		return nil, errors.New("Такого студента нет!")
	}
	return stud, nil
}

// Put - Метод StudentsRepo, который принимает ссылку на объект структуры student student.Student
// и по имени параметру student.Name сохраняет в экземляр StudentsRepo.
// Если ключ student.Name уже есть в экземляр StudentsRepo, то возращает ошибку.
// Если запись прошла успешно, возвращает nil.
func (sr StudentsRepo) Put(student *student.Student) error {
	_, status := sr[student.Name]
	if status {
		return errors.New("Такой студент уже существует!")
	}
	sr[student.Name] = student
	return nil
}

// ShowStudents - Принимает на вход ссылку на экземпляр *StudentsRepo
// Выводит в консоль всех студентов из экземпляра *StudentsRepo
func ShowStudents(storage *StudentsRepo) {
	fmt.Println("\nСтуденты из хранилища:")
	for _, student := range *storage {
		fmt.Println(student.Name, student.Age, student.Grade)
	}
}
