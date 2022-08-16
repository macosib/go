package student

type Student struct {
	Name  string
	Age   int
	Grade int
}

//CreateNewStudent - Функция на вход принимает name формат string, age и grade формат int.
// Возвращает ссылку на новый объект структуры Student.
func CreateNewStudent(name string, age, grade int) *Student {
	return &Student{name, age, grade}
}
