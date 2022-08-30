package main

import (
	"fmt"
	"sync"
)

func main() {
	var db []*User
	userSvc := Newuser(db)

	persons :=
		[]string{
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

	var wg sync.WaitGroup
	wg.Add(len(persons))
	for _, val := range persons {

		go func(name string) {

			result := userSvc.Register(&User{Name: name})
			fmt.Println(result, "Berhasil didaftarkan")
			wg.Done()

		}(val)
	}

	wg.Wait()

	resGet := userSvc.GetUser()
	fmt.Println("-----------Hasil get user-------------")
	for _, v := range resGet {
		fmt.Println(v.Name)
	}

}

type userInterface interface {
	Register(u *User) string
	GetUser() []*User
}

type User struct {
	Name string
}

type dbUser struct {
	db []*User
}

func Newuser(db []*User) userInterface {
	return &dbUser{
		db: db,
	}
}

func (du *dbUser) Register(u *User) string {

	du.db = append(du.db, u)

	return u.Name

}

func (du *dbUser) GetUser() []*User {
	return du.db
}
