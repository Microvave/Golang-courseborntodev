package main

import "fmt"

var productName [4]string

func main() {
	productName[0] = "macbook"
	productName[1] = "macbook1"
	productName[2] = "macbook2"
	productName[3] = "macbook3"

	price := [4]float32{1, 2, 3, 4}

	fmt.Println(productName)
	fmt.Println(price)
}
