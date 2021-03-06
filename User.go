package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
)

//User struct mode
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Users group of users
type Users []User

//SaveUser method for write user information to disk
func SaveUser(userInfo *User) string {
	userid := base64.StdEncoding.EncodeToString([]byte(userInfo.Name))
	fi, err := os.Create("/tmp/" + userid)
	defer fi.Close()
	if err != nil {
		println(err)
	} else {
		if jsonStr, err := json.Marshal(userInfo); err == nil {
			_, err := fi.Write([]byte(jsonStr))
			if err != nil {
				userid = ""
			}
		}
	}
	return userid
}

//GetUserToJSONStr method for get the user information from the disk
func GetUserToJSONStr(userID string) string {
	userInfo := ""
	jsonStr, err := ioutil.ReadFile("/tmp/" + userID)
	if err == nil {
		userInfo = string(jsonStr)
	}
	return userInfo
}
