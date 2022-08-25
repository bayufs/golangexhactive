package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct{ name, alamat, pekerjaan, alasan string }

func main() {

	args := os.Args[1]
	intVar, _ := strconv.Atoi(args)
	getInput(intVar)

}

func getInput(arg int) {
	persons := []*Person{
		{name: "bayu", alamat: "Jakbar", pekerjaan: "Elite global", alasan: "Pengen bisa"},
		{name: "hilmi", alamat: "Jaktim", pekerjaan: "Dewan pengawas programmer", alasan: "Belum nanya"},
	}

	fmt.Println("Nama : ", persons[arg].name)
	fmt.Println("Alamat : ", persons[arg].alamat)
	fmt.Println("Pekerjaan : ", persons[arg].pekerjaan)
	fmt.Println("Alasan : ", persons[arg].alasan)
}
