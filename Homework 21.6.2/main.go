package main

import (
	"fmt"
	"time"
)

func decorator(A func(int, int) int, arg1, arg2 int) {
	defer fmt.Println("Завершение вызова функции:", time.Now().UTC().Format("2006-01-02 15:04:05"))
	defer fmt.Println("Результат вызова функции: ", A(arg1, arg2))
	fmt.Println("Старт вызова функции:", time.Now().UTC().Format("2006-01-02 15:04:05"))

}

func main() {
	decorator(func(arg1, arg2 int) int { return arg1 + arg2 }, 4, 5)
	decorator(func(arg1, arg2 int) int { return arg1 - arg2 }, 4, 5)
	decorator(func(arg1, arg2 int) int { return arg1 * arg2 }, 4, 5)
}
