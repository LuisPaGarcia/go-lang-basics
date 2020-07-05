package main

import "fmt"

func multiple(first string, last string) (string, string) {
	return first, last
}

func main() {
	f, l := multiple("luispa", "garcia")
	fmt.Println(f, l)
	f2, _ := multiple("luispa", "garcia")
	fmt.Println(f2)
}
