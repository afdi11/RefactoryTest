package main

import (
	"RefactoryTest/Logic-Test/Nearest-Fibonacci/fibonacci"
	"fmt"
)

func main() {
	var arr = []int{15, 1, 3}
	fmt.Println("the nearest fibonacci from array ", arr, " yaitu : ", fibonacci.Nearest(arr))
}
