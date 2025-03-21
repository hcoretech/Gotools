package main

import (
	"fmt"
)

type user struct {
	id       int
	name     string
	username string
	password string
	isLogin  bool
}

func main() {

	userLogin(user{
		name:     "Henry ",
		username: "Hcore",
		password: "tony@2015164218",
	})

	// time := 14

	// if time <= 12 {
	// 	fmt.Println("good morning")
	// 	time = 13
	// }
	// if time <= 18 {
	// 	fmt.Println("good afternoon")
	// 	time = 19
	// }
	// if time <= 24 {
	// 	fmt.Println("good evening")
	// 	time = 0
	// }

}

func userLogin(value user) bool {

	username := "Hcore"
	password := "tony@2015164218"

	if value.name == "" {
		fmt.Println("no user name input ")
		return false
	}
	if value.password == "" {
		fmt.Println("no password input")
		return false

	}
	if value.username != username {
		fmt.Println("incorrect username")
		return false
	}
	if value.password != password {
		fmt.Println("wrong password")
	}
	fmt.Println("sucesful login")
	fmt.Println("welcome", value.name)

	return true

}
