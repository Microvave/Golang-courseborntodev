package main

import "fmt"

func average(average1 float64, average2 float64, average3 float64 ) float64 {
	return (average1 + average2 + average3)/3
}

func main() {
	var average1 float64
	var average2 float64
	var average3 float64

	fmt.Print("กรอกตัวเลข 1 : ")
	fmt.Scan(&average1)

	fmt.Print("กรอกตัวเลข 2 : ")
	fmt.Scan(&average2)

	fmt.Print("กรอกตัวเลข 3 : ")
	fmt.Scan(&average3)

	fmt.Println("ตัวเลขที่ 1 : ", average1)
	fmt.Println("ตัวเลขที่ 2 : ", average2)
	fmt.Println("ตัวเลขที่ 3 : ", average3)

	fmt.Println("ผลรวม = ", average(average1, average2, average3))
}
