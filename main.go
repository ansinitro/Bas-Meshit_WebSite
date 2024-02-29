package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
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
	http.HandleFunc("/sendEverybody", sendEverybody)
	port := os.Getenv("PORT")
	if port == "" {
		port = "777" // Default port if PORT environment variable is not set
	}
	log.Printf("Server listening on port %s", port)
	err = http.ListenAndServe(":"+port, context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
