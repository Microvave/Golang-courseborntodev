package main

import "fmt"

func main() {

	var courseName []string
	var name string

	for {
		var menu int

		// แสดงเมนู
		fmt.Println("กรุณาเลือกเมนู:")
		fmt.Println("1 = เพิ่มชื่อ")
		fmt.Println("2 = แสดงรายชื่อทั้งหมด")
		fmt.Println("3 = ออกจากโปรแกรม")

		// รับค่าเมนู
		fmt.Print("เลือกเมนู: ")
		fmt.Scan(&menu)
		switch menu {
		case 1:
			fmt.Scan(&name)
			courseName = append(courseName, name)
			// fmt.Println(courseName)
		case 2:
			fmt.Println("-", courseName)
		case 3:
			fmt.Println("ลาก่อน")
			return
		default:
			fmt.Println("กรุณากรอกตัวเลข 1-3")
		}
	}
}
