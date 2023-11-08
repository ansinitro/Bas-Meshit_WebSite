package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func contactUsHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact-us.html", nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		tpl.ExecuteTemplate(w, "login.html", nil)
	}
	tpl.ExecuteTemplate(w, "profile.html", user)
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "registration.html", nil)
}

func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "forgot-password.html", nil)
}
