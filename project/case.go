package main

import "fmt"

var input string


func main() {


	 var menu int

    // แสดงเมนู
    fmt.Println("กรุณาเลือกเมนู:")
    fmt.Println("1 = ทักทาย")
    fmt.Println("2 = บวกเลข")
    fmt.Println("3 = ออกจากโปรแกรม")

    // รับค่าเมนู
    fmt.Print("เลือกเมนู: ")
    fmt.Scan(&menu)

	switch menu {
	case 1:
		fmt.Println("สวัสดี")
	case 2:
		var a, b int
        fmt.Print("กรอกตัวเลขที่ 1: ")
        fmt.Scan(&a)
        fmt.Print("กรอกตัวเลขที่ 2: ")
        fmt.Scan(&b)
        fmt.Println("ผลรวม =", a+b)
	case 3:
		fmt.Println("ลาก่อน")
	default:
		fmt.Println("กรุณากรอกตัวเลข 1-3")
	}
}
