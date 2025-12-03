package main

import "fmt"

func main() {

    for {
        var menu int
        
        // แสดงเมนู
        fmt.Println("=== เมนูหลัก ===")
        fmt.Println("1) ตัวเลือกที่หนึ่ง")
        fmt.Println("2) ตัวเลือกที่สอง")
        fmt.Println("3) ออกจากโปรแกรม")

        
        fmt.Print("เลือกเมนู: ")
        fmt.Scan(&menu)

        switch menu {
        case 1:
            fmt.Println("เลือกอีกครั้ง")

        case 2:
            fmt.Println("เลือกใหม่")

        case 3:
            fmt.Println("กำลังออกจากโปรแกรม...")
            return  

        default:
            fmt.Println("กรุณาเลือก 1 - 3")
        }

        fmt.Println() 
    }
}
