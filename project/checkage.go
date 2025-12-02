package main

import "fmt"

var age int

func main() {
	fmt.Println("กรอกอายุของคุณ")
	fmt.Scanf("%d", &age)
	fmt.Println(" age = ", age)
	if age >= 18 {
		fmt.Println("เป็นผู้ใหญ่")
	} else {
		fmt.Println("ไม่เป็นผู้ใหญ่")
	} 
}
