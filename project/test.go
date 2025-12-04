package main

import "fmt"

func main() {

	var courseName []string // slice เก็บรายชื่อ
	var name string         // ตัวแปรรับชื่อแต่ละครั้ง

	for { // ลูปเมนูไม่รู้จบ
		var menu int

		fmt.Println("กรุณาเลือกเมนู:")
		fmt.Println("1 = เพิ่มชื่อ")
		fmt.Println("2 = แสดงรายชื่อทั้งหมด")
		fmt.Println("3 = ออกจากโปรแกรม")

		fmt.Print("เลือกเมนู: ")
		fmt.Scan(&menu)

		switch menu {

		case 1:
			fmt.Print("กรอกชื่อ: ")
			fmt.Scan(&name)
			courseName = append(courseName, name)
			fmt.Println("เพิ่มชื่อเรียบร้อย")

		case 2:
			if len(courseName) == 0 {
				fmt.Println("ยังไม่มีข้อมูล")
			} else {
				fmt.Println("รายชื่อทั้งหมด:")
				for i := 0; i < len(courseName); i++ {
					fmt.Println("-", courseName[i])
				}
			}

		case 3:
			fmt.Println("ออกจากโปรแกรม...")
			return // จบโปรแกรมทันที

		default:
			fmt.Println("กรุณาเลือกหมายเลข 1 - 3 เท่านั้น")
		}

		fmt.Println() // เว้นบรรทัดเพื่อความสวยงาม
	}
}
