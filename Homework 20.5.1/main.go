package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 3

func getTestMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = getRandomNumber()
		}
	}
	fmt.Println("Исходная матрица:", matrix)
	return matrix
}

func getDeterminantThird(matrix [][]int) int {
	var determinant int
	determinant =
		matrix[0][0]*matrix[1][1]*matrix[2][2] +
			matrix[0][2]*matrix[1][0]*matrix[2][1] +
			matrix[0][1]*matrix[1][2]*matrix[2][0] -
			matrix[0][2]*matrix[1][1]*matrix[2][0] -
			matrix[0][1]*matrix[1][0]*matrix[2][2] -
			matrix[0][0]*matrix[1][2]*matrix[2][1]
	fmt.Println("Определьтель матрицы равен:", determinant)
	return determinant
}

func getRandomNumber() int {
	return rand.Intn(10)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	getDeterminantThird(getTestMatrix(size))
}
