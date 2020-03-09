package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const a = 500000

	const d = 3e20 / a
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))
}
