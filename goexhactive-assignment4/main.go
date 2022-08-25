package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct{ name string }

func main() {

	args := os.Args[1]
	intVar, _ := strconv.Atoi(args)
	getInput(intVar)

}

func getInput(arg int) {
	var name = []string{
		"bayu",
		"fajar",
		"sidik",
		"nugraha",
		"muthus",
		"muhamad",
		"hilmi",
		"eka",
		"bintang",
		"mail"}

	fmt.Println(name[arg])
}
