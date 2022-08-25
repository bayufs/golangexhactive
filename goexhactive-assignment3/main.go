package main

import "fmt"

type Person struct{ name string }

func main() {

	persons := []*Person{{name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}, {name: "andi"}, {name: "budi"}}
	printFriends := func(friends []*Person) {
		for _, p := range friends {
			fmt.Println(p.name)
		}
	}

	printFriends(persons)

}
