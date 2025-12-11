package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("Microwave\n Hi")
	err := os.WriteFile("file/data.txt", data1, 0644)

	check(err)

	f, err := os.Create("employNames")
	check(err)
	defer f.Close()

	data2 := []byte("John\n Doe\n Smith")
	err = os.WriteFile("file/employNames.txt", data2, 0644)
}
