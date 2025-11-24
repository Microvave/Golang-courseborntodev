package main

import "fmt"

var input string
var colorcode string

func main() {
	fmt.Scan(&input)
	switch input {
	case "blue":
		fmt.Println("one")
	case "green":
		fmt.Println("two")
	case "pink":
		fmt.Println("two")
	case "yellow":
		fmt.Println("two")
	default:
		fmt.Println("no color")
	}
}
