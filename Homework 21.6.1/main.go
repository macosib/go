package main

import (
	"fmt"
	"math"
)

func getResult(x int16, y uint8, z float32) float32 {
	S := 2*float64(x) + math.Pow(float64(y), 2) - 3/float64(z)
	fmt.Println(S)
	return float32(S)
}

func main() {
	fmt.Println(getResult(2, 4, 3))
}
