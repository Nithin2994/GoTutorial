package main

import (
	"fmt"
	"funcs"
	"strconv"
)

//UsersList ...
var UsersList []*User

//User ...
type User struct {
	name     string // ""
	id       int    // 0
	password string // ""
}

//PrintUser ..
func (u *User) PrintUser() {
	fmt.Println("----")
	fmt.Println("Name : ", u.name)
	fmt.Println("Id : ", strconv.Itoa(u.id))
	fmt.Println("Password : ", u.password)
	fmt.Println("----")
}

//Login ...
func (u *User) Login(username string, password string) bool {
	if u.name == username && u.password == password {
		return true
	}
	return false
}

func main() {
	s := "Nithin"
	tags := make([]string, 0)
	tags = append(tags, "Java")
	tags = append(tags, "Go")

	//funcs.Print(s, 1, tags)
	funcs.AddTag(&tags, "Python")

	formatStr, id := funcs.FormatedPrintStatement(s, 1, tags)

	fmt.Println(formatStr, id) // Nithin with Id : 1 10

	user1 := CreateUser("Nithin", "123123", 1) //address : 1231231

	UpdateUserID(user1)

	//fmt.Println(user1)
	user1.PrintUser()

	if user1.Login("Nit", "123") {
		fmt.Println("valid User")
	} else {
		fmt.Println("in-valid User")
	}

}

//CreateUser ...
func CreateUser(name string, pass string, id int) *User {
	//var user1 User
	user1 := new(User)
	user1.name = name
	user1.password = pass
	user1.id = id
	if UsersList == nil {
		UsersList = make([]*User, 0)
	}
	UsersList = append(UsersList, user1)
	return user1
}

//UpdateUserID ..
func UpdateUserID(u *User) { //address : 1231231
	u.id = 10
}
