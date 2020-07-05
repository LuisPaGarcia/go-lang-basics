package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("hey", s)

	s[0] = "a"
	s[1] = "a"
	s[2] = "a"
	fmt.Println("set", s)

	s = append(s, "dd")
	s = append(s, "ad")
	fmt.Println("set", s)

}
