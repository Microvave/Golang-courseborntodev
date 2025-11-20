package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"java", "python"}
	fmt.Println(courseName)
	courseName = append(courseName, "c", "c#")
	fmt.Println(courseName)

	coursweb := courseName[2:4]
	fmt.Println(coursweb)

	coursweb = courseName[:2]
	fmt.Println(coursweb)
}
