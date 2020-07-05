package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10

	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	case 10:
		fmt.Println("diez")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend bitch")
	default:
		fmt.Println("weekday")
	}

	whatType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Booleano")
		case int:
			fmt.Println("Entero")
		default:
			fmt.Printf("Ningun tipo controlado: %T\n", t)
		}
	}
	whatType(123)
	whatType("kj")

}
