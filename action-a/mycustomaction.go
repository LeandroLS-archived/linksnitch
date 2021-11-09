package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Olá")
	file := CreateFile()
	fmt.Printf(file.Name())
}

func CreateFile() *os.File {

	file, err := os.Create("aaaa.txt")
	if err != nil {
		fmt.Println(err)
	}
	return file
}
