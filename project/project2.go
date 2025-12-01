package main

import "fmt"

func add(number1 int, number2 int) int {
	return number1 + number2
}

func main() {
	var number1 int
	var number2 int

	fmt.Print("กรอกตัวเลข 1 : ")
	fmt.Scan(&number1)

	fmt.Print("กรอกตัวเลข 2 : ")
	fmt.Scan(&number2)

	fmt.Println("number1 : ", number1)
	fmt.Println("number2 : ", number2)

	fmt.Println("ผลการบวก = ", add(number1, number2))
}
