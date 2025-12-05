package main

import (
	"fmt"
)

var name string
var phone string
var searchName string
var phoneBook = make(map[string]string)

func main() {

	for {
		var menu int

		// แสดงเมนู
		fmt.Println("กรุณาเลือกเมนู:")
		fmt.Println("1 = เพิ่มข้อมูลเบอร์โทรศัพท์")
		fmt.Println("2 = ค้นหาข้อมูลเบอร์โทรศัพท์ตามชื่อ")
		fmt.Println("3 = แสดงรายชื่อทั้งหมด")
		fmt.Println("4 = ออกจากโปรแกรม")

		// รับค่าเมนู
		fmt.Print("เลือกเมนู: ")
		fmt.Scan(&menu)
		switch menu {
		case 1:

			fmt.Println("กรอกชื่อ")
			fmt.Scan(&name)
			fmt.Println("กรอกเบอร์โทรศัพท์")
			fmt.Scan(&phone)
			phoneBook[name] = phone
			fmt.Println("บันทึกข้อมูลเรียบร้อย")
		case 2:
			fmt.Println("กรอกชื่อที่ต้องการค้นหา")
			fmt.Scan(&searchName)
			phone, ok := phoneBook[searchName]
			if ok {
				fmt.Println("เบอร์โทรศัพท์ของ", searchName, "คือ", phone)
			} else {
				fmt.Println("ไม่พบข้อมูลชื่อ", searchName)
			}
		case 3:
			for key, value := range phoneBook {
				fmt.Println("- ชื่อ:", key, "เบอร์โทรศัพท์:", value)
			}

		case 4:
			fmt.Println("จบการทำงาน ..... ")
			return
		default:
			fmt.Println("กรุณากรอกตัวเลข 1-4")
		}
	}

}
