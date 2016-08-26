package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/robbert229/jwt"
)

const mysecret = "MySecretKey"

//Index Handler for root
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//UserAdd Handle for create a new user
func UserAdd(w http.ResponseWriter, r *http.Request) {
	isValidToken := validateJwt(r)
	if isValidToken {
		var user User
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &user); err != nil {
			userid := SaveUser(&user)
			if userid != "" {
				w.WriteHeader(http.StatusCreated)
				fmt.Fprintf(w, "User Id: %q", userid)
			} else {
				w.WriteHeader(http.StatusExpectationFailed)
				fmt.Fprint(w, "Error")
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//UserInfo Handle for get the user information
func UserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are trying to get the user information")
}

//GoToken Handle for create and print a new token
func GoToken(w http.ResponseWriter, r *http.Request) {
	algorithm := jwt.HmacSha256(mysecret)
	claims := jwt.NewClaim()
	claims["isAdmin"] = true
	token, err := jwt.Encode(algorithm, claims)
	if err == nil {
		fmt.Fprint(w, token)
	}
}

func validateJwt(r *http.Request) bool {
	valid := false
	token := r.Header.Get("Authorization")
	if token != "" {
		token = strings.Replace(token, "Bearer ", "", -1)
		algorithm := jwt.HmacSha256(mysecret)
		if jwt.IsValid(algorithm, token) == nil {
			claims, err := jwt.Decode(algorithm, token)
			if err != nil {
				valid = false
			} else {
				valid = claims["isAdmin"].(bool)
			}
		}
	}
	return valid
}
