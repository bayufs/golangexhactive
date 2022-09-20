package main

import (
	"fmt"
	"strconv"
)

func main() {
	angka := "1"

	convert, err := strconv.Atoi(angka)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(convert)
	}
}
