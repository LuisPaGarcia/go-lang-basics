package main

import "fmt"

func main(){
	
	if 7 % 2 == 0 {
		fmt.Println("that")
	} else {
		fmt.Println("this")
	}

	var ocho, cuatro int = 8, 4

	if ocho % cuatro == 0 {
			fmt.Println("8 es divisible entre 4")
	}


	if num := 9; num < 0 {
			fmt.Println(num, "es negativo")
	} else if num < 10 {
			fmt.Println(num, "tiene un digito")
	} else {
			fmt.Println(num, "tiene multiples digitos")
	}

	if (1 > 2) || (3 > 4) {
		fmt.Println("imposible")
	} else {
		fmt.Println("correcto")
	}

}