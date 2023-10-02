package main

import (
	"fmt"
	"go_auth_basic/helpers"
	"net/http"
	"strings"
)

func main() {
	username, email, fullName, cpwd, pwd := "", "", "", "", ""
	router := http.NewServeMux()
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//for key, value := range r.Form {
		//	fmt.Printf("%s= %s\n", key, value)
		//}
		username = r.FormValue("username")
		email = r.FormValue("email")
		fullName = r.FormValue("full_name")
		pwd = r.FormValue("password")
		cpwd = r.FormValue("confirm_password")
		usernameCheck := helpers.IsEmpty(username)
		emailCheck := helpers.IsEmpty(email)
		fullNameCheck := helpers.IsEmpty(fullName)
		passwordCheck := helpers.IsEmpty(pwd)
		passwordConformCheck := helpers.IsEmpty(cpwd)
		validateEmail := strings.Contains(email, "@")

		if usernameCheck || emailCheck || fullNameCheck || passwordCheck || passwordConformCheck || !validateEmail {
			fmt.Fprintf(w, "Maglumatlar doly derejede doldurylmady!!!\n")
			if !validateEmail {
				fmt.Fprintf(w, "email  maglumatlarynda  '@' belgi ulanylmady!!!")
			}
			return
		}
		if pwd == cpwd {
			fmt.Fprintf(w, "Register Successfully!!!")
			return
		} else {
			fmt.Fprintf(w, "Password-lar biri-birine gabat gelenok!!!")
			return
		}
	})

	// login func
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		emailCheck := helpers.IsEmpty(email)
		passwordCheck := helpers.IsEmpty(pwd)
		Email := "polat@gmail.com"
		Password := "12345"
		if emailCheck || passwordCheck {
			fmt.Fprintf(w, "Maglumatlar doldurylmady")
			return
		}
		if Email == email && Password == pwd {
			fmt.Fprintf(w, "Login successfully")
			return
		} else {
			fmt.Fprintf(w, "Giris maglumatlary nadogry!!!")
			return
		}

	})

	http.ListenAndServe(":8080", router)
}
