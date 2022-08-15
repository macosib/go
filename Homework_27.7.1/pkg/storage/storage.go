package storage

import (
	"Homework_27.7.1/pkg/student"
	"errors"
	"fmt"
)

type StudentsRepo map[string]*student.Student

func CreateStudentsRepo() StudentsRepo {
	return make(map[string]*student.Student)
}

func (sr StudentsRepo) Get(student string) (*student.Student, error) {
	stud, status := sr[student]
	if !status {
		return nil, errors.New("Такого студента нет!")
	}
	return stud, nil
}

func (sr StudentsRepo) Put(student *student.Student) error {
	_, status := sr[student.Name]
	if status {
		return errors.New("Такой студент уже существует!")
	}
	sr[student.Name] = student
	return nil
}

func ShowStudents(storage *StudentsRepo) {
	fmt.Println("\nСтуденты из хранилища:")
	for _, student := range *storage {
		fmt.Println(student.Name, student.Age, student.Grade)
	}
}
