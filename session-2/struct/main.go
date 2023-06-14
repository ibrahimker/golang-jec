package main

import "fmt"

type User struct {
	name string
	age  int
}

var users []User
var users2 []*User

func main() {
	u := User{
		name: "ronald",
		age:  18,
	}
	fmt.Println(u)

	var u2 User
	u2.name = "Dwi"
	u2.age = 19
	fmt.Println(u2)

	u3 := &User{
		name: "anggi",
		age:  19,
	}
	fmt.Println(u3)

	u4 := struct {
		user User
	}{}
	u4.user = u2

	users2 := []*User{&u, &u2, u3, nil}
	fmt.Println(users2)

	// Latihan 2
	fmt.Println("===========Start Latihan 2=======")
	Register(u)
	Register(u2)
	Register(*u3)
	for k, v := range Get() {
		fmt.Println("User ke ", k, " adalah ", v)
	}
}

func Register(user User) {
	users = append(users, user)
}

func Get() []User {
	return users
}
