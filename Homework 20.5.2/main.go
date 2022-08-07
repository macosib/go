package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func getTestMatrix(row, col int, flagZero bool) [][]int {
	matrix := make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
		for j := 0; j < len(matrix[i]); j++ {
			if !flagZero {
				matrix[i][j] = getRandomNumber()
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func getMultMatrix(matrix1 [][]int, matrix2 [][]int) ([][]int, error) {
	rowRes := len(matrix1)
	colRes := len(matrix2[0])
	col := len(matrix1[0])
	row := len(matrix2)

	fmt.Println("Исходная матрица:", matrix1)
	fmt.Println("Исходная матрица:", matrix2)

	result := getTestMatrix(rowRes, colRes, true)

	if col != row {
		return result, errors.New("Эти матрицы нельзя умножить!")
	}
	for i := 0; i < rowRes; i++ {
		for j := 0; j < colRes; j++ {
			sum := 0
			for k := 0; k < col; k++ {
				sum += matrix1[i][k] * matrix2[k][j]
			}
			result[i][j] = sum
		}
	}
	return result, nil
}

func getRandomNumber() int {
	return rand.Intn(10)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	response, err := getMultMatrix(getTestMatrix(3, 5, false), getTestMatrix(5, 4, false))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Результат умножения:", response)
}
