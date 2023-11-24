package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		return
	}
	tpl.ExecuteTemplate(w, "profile.html", user)
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "registration.html", nil)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Printf("Fail ParseForm()")
	}
	name := r.Form.Get("name")
	surname := r.Form.Get("surname")
	phone := r.Form.Get("phone-number")
	email := r.Form.Get("email")
	course := r.Form.Get("course")
	age := r.Form.Get("age")

	phoneInt, err := strconv.Atoi(strings.Trim(phone, " "))
	if err != nil {
		tpl.ExecuteTemplate(w, "registration.html", "Input without spaces phone")
	}
	ageInt, err := strconv.Atoi(strings.Trim(age, " "))
	if err != nil {
		tpl.ExecuteTemplate(w, "registration.html", "Input without spaces age")
	}

	fmt.Printf("%v, %v, %v, %T\n", phone, phoneInt, err, phoneInt)
	fmt.Printf("%v, %v, %T\n", ageInt, err, ageInt)

	var insertStatement *sql.Stmt
	insertStatement, err = DB.Prepare("INSERT INTO course (name, surname, phone, email, course, age) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Printf("error preparing statement %v\n", err)
		tpl.ExecuteTemplate(w, "registration.html", "There Occuring with Registration")
		return
	}
	defer insertStatement.Close()
	_, err = insertStatement.Exec(strings.Trim(name, " "), strings.Trim(surname, " "), phoneInt, strings.Trim(email, " "),
		strings.Trim(course, " "), ageInt)
	if err != nil {
		log.Printf("Error inserting new user")
		tpl.ExecuteTemplate(w, "registration.html", "Error inserting new user")
		return
	}

	tpl.ExecuteTemplate(w, "success.html", "Сіз курсқа жазылдыңыз!")
}
