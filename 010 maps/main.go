package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["edad"] = 12
	m["estatura"] = 11
	m["posty"] = 0
	fmt.Println(m)

	v := m["edad"]
	fmt.Println(v)
	fmt.Println(len(m))

	delete(m, "edad")
	fmt.Println(m)
	// Existe en el map?
	_, estatura := m["estatura"] // true
	_, edad := m["edad"]         // false
	fmt.Println("prs:", estatura, edad)

	n := map[string]int{"item": 1, "case": 2}
	fmt.Println(n)
}
