package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "nombre", "b": "apellido"}

	for k, v := range kvs {
		fmt.Println(k, v)
	}

	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}

	for i, c := range "GO" {
		fmt.Println(i, c)
	}

}
