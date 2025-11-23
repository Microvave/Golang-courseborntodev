package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)

	//add
	product["Macbook"] = 40000
	product["Macbook1"] = 40001
	product["Macbook2"] = 40002
	fmt.Println("product =", product)

	//delete
	delete(product, "Macbook1")
	fmt.Println("product =", product)

	//update
	product["Macbook2"] = 20002
	fmt.Println("product =", product)

	value1 := product["Macbook"]
	fmt.Println("value1 =", value1)

	courseName := map[string]string{"101": "java", "102": "c"}
	fmt.Println("courseName =", courseName)
}
