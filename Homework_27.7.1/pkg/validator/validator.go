package validator

import (
	"errors"
	"strconv"
	"strings"
)

func ValidInputData(data string) (string, int, int, error) {
	result := strings.Split(strings.Trim(data, " \n"), " ")

	if len(result) != 3 {
		return "", 0, 0, errors.New("Введено слишком мало или слишком много данных!")
	}

	name := strings.ContainsAny(result[0], " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~.0123456789")
	age, ageErr := strconv.Atoi(result[1])
	grade, errGrade := strconv.Atoi(result[2])

	if ageErr != nil || name || errGrade != nil {
		return "", 0, 0, errors.New("Ошибка при вводе данных! Проверьте правильность ввода!")
	}

	return result[0], age, grade, nil
}
