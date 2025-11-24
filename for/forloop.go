package main

import "fmt"

const count = 15

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input = ", input)
		if input == "exit" {
			break
		}
	}
}
