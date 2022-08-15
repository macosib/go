package student

type Student struct {
	Name  string
	Age   int
	Grade int
}

func CreateNewStudent(name string, age, grade int) *Student {
	return &Student{name, age, grade}
}
