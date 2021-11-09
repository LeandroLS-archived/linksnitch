package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Olá")
	dat, err := os.ReadFile("test.md")
	check(err)
	fmt.Print(string(dat))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
