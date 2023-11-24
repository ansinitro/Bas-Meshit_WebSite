package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var store = sessions.NewCookieStore([]byte("bismillah"))

type User struct {
	Name, Surname, Email, Course, ErrMessage string
	Id, Phone, Age                           int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	err := OpenDB()
	if err != nil {
		log.Printf("error connecting postgresql database %v", err)
	}
	defer CloseDB()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", Auth(indexHandler))
	http.HandleFunc("/index", Auth(indexHandler))
	http.HandleFunc("/forgotPassword", forgotPasswordHandler)
	http.HandleFunc("/forgotPasswordChange", forgotPasswordHandler)
	http.HandleFunc("/contact-us", contactUsHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/registration", registrationHandler)
	http.HandleFunc("/verifyemail", verifyEmailHandler)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/signin", signInHandler)
	http.HandleFunc("/logout", deleteSessionHandler)
	http.HandleFunc("/admin", adminHandler)
	http.ListenAndServe("localhost:777", context.ClearHandler(http.DefaultServeMux))
}
