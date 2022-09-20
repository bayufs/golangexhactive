package main

import "fmt"

type Person struct{ name string }

func main() {

	persons := []*Person{{name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}}
	printFriends := func(friends []*Person) {
		for index, p := range friends {
			fmt.Println(index, ":", p.name)
		}
	}

	printFriends(persons)

}
