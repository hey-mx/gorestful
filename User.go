package main

//User struct mode
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Users group of users
type Users []User
