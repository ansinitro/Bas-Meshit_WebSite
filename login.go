package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func signInHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Fail ParseForm()")
	}
	var name string
	email := r.Form.Get("email")
	passwd := r.Form.Get("passwd")

	//check is user with same email exist
	stmt := "SELECT name, password FROM users WHERE email = $1"
	row := DB.QueryRow(stmt, email)
	var hash string
	err = row.Scan(&name, &hash)
	if err == sql.ErrNoRows {
		log.Printf("username doesn't exists, err:", err)
		tpl.ExecuteTemplate(w, "login.html", "Email not exist.")
		return
	}

	//hash from passwd
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	if err == nil {
		log.Print("successe login")
		createSession(w, r, name, email)
		tpl.ExecuteTemplate(w, "index.html", "You succesufully logged in")
		return
	}
	log.Print("incorrect password")
	tpl.ExecuteTemplate(w, "login.html", "Incorrect Password")
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Fail ParseForm()")
	}
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	passwd := r.Form.Get("passwd")

	//check is user with same email exist
	stmt := "SELECT name FROM users WHERE email = $1"
	row := DB.QueryRow(stmt, email)
	var value string
	err = row.Scan(&value)
	if err != sql.ErrNoRows {
		log.Printf("username already exists, err:", err)
		tpl.ExecuteTemplate(w, "login.html", "Username Already Exists")
		return
	}

	//hash from passwd
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("bcrypt err %v\n", err)
		tpl.ExecuteTemplate(w, "login.html", "There Occuring with Registration")
		return
	}
	var insertStatement *sql.Stmt
	insertStatement, err = DB.Prepare("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)")
	if err != nil {
		log.Printf("error preparing statement %v\n", err)
		tpl.ExecuteTemplate(w, "login.html", "There Occuring with Registration")
		return
	}
	defer insertStatement.Close()
	_, err = insertStatement.Exec(name, email, string(hash))
	if err != nil {
		log.Printf("Error inserting new user")
		tpl.ExecuteTemplate(w, "login.html", "Error inserting new user")
		return
	}
	tpl.ExecuteTemplate(w, "login.html", "User was created")
	fmt.Println("Everything good")
}
