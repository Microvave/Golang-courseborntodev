package main

import "fmt"

var username string
var password int

func main() {

	username = "admin"
	password = 1234

	isLogin := false
	for i := 1; i <= 3; i++ {

		var inputUser string
		var inputPass int
		fmt.Print("กรุณากรอกชื่อผู้ใช้: ")
		fmt.Scan(&inputUser)

		fmt.Print("กรุณากรอกรหัสผ่าน: ")
		fmt.Scan(&inputPass)

		if inputUser == username && inputPass == password {
			fmt.Println("ยินดีต้อนรับเข้าสู่ระบบ")
			isLogin = true
			break
		} else {
			fmt.Println("ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง")
		}

	}
	if !isLogin {
		fmt.Println("ล็อกระบบ")
	}

}
