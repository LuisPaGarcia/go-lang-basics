package main

import "fmt"

func plus(first int, second int) int {
	return first + second
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	result := plus(1, 2)
	fmt.Println(result)

	result2 := plusPlus(1, 2, 3)
	fmt.Println(result2)
}
