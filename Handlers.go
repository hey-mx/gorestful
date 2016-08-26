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
			users := Users{
				user,
			}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(users); err != nil {
				println(err)
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//UserInfo Handle for get the user information
func UserInfo(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userid"])
	userInfo := Users[userID]*/
	fmt.Fprint(w, "You are trying to get the user information")
}

func validateJwt(r *http.Request) bool {
	valid := false
	token := r.Header.Get("Authorization")
	if token != "" {
		token = strings.Replace(token, "Bearer ", "", -1)
		algorithm := jwt.HmacSha256("MySecretKey")
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
