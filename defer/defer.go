package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result : ", result)
}
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func deferloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("i : ", i)
	}
}
func main() {

	// fmt.Println("Welcome")
	// defer fmt.Println("End")
	// defer add(10, 20)
	// add(30, 40)
	// add(50, 60)
	loop()
	deferloop()
}
