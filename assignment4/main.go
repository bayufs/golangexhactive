package main

import "fmt"

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

	for _, val := range persons {
		result := userSvc.Register(&User{Name: val})
		fmt.Println(result, "Berhasil didaftarkan")
	}

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
