package main

import "fmt"

type emoployee struct {
	emoployeeID   string
	emoployeeName string
	phone         string
}

func main() {

	// emoployee1 := emoployee{
	// 	emoployeeID:   " 101",
	// 	emoployeeName: "wave",
	// 	phone:         "111",
	// }
	// fmt.Println("emoployee1 = ", emoployee1)

	// emoployeeList := [3]emoployee{}
	// emoployeeList[0] = emoployee{
	// 	emoployeeID:   "1",
	// 	emoployeeName: "wave",
	// 	phone:         "11",
	// }
	// emoployeeList[1] = emoployee{
	// 	emoployeeID:   "2",
	// 	emoployeeName: "wave2",
	// 	phone:         "22",
	// }
	// emoployeeList[2] = emoployee{
	// 	emoployeeID:   "3",
	// 	emoployeeName: "wave",
	// 	phone:         "33",
	// }

	emoployeeList := []emoployee{}
	emoployee1 := emoployee{
		emoployeeID:   "101",
		emoployeeName: "wave",
		phone:         "11",
	}
	emoployee2 := emoployee{
		emoployeeID:   "102",
		emoployeeName: "vave",
		phone:         "22",
	}

	emoployeeList = append(emoployeeList, emoployee1)
	emoployeeList = append(emoployeeList, emoployee2)
	fmt.Println("emoployee = ", emoployeeList)
}
