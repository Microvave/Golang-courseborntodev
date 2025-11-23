package main

import "fmt"

var score int

func main() {
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println(" score = ", score)
	if score >= 80 {
		fmt.Println("A grade")
	} else if score >= 70 {
		fmt.Println("B grade")
	} else if score >= 60 {
		fmt.Println("C grade")
	} else if score >= 50 {
		fmt.Println("D grade")
	} else {
		fmt.Println("F grade")
	}
}
