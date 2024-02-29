package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"golang.org/x/crypto/bcrypt"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "admin-login.html", nil)
		return
	}
	r.ParseForm()
	passwd := r.Form.Get("passwd")

	stmt := "SELECT hash FROM users WHERE email = $1"
	row := DB.QueryRow(stmt, "admin@admin.com")
	var hash string
	err := row.Scan(&hash)
	if err == sql.ErrNoRows {
		log.Printf("username doesn't exists, %v:", err)
		tpl.ExecuteTemplate(w, "login.html", "Email not exist.")
		return
	}

	//hash from passwd
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	if err == nil {
		log.Print("successe login")

		rows, err := DB.Query("SELECT * FROM course")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := make([]*User, 0, 10)
		// Iterate through the query results
		for rows.Next() {
			var id, phone, age int
			var name, surname, email, course string
			if err := rows.Scan(&id, &name, &surname, &phone, &email, &course, &age); err != nil {
				log.Fatal(err)
			}
			users = append(users, &User{Id: id, Name: name, Surname: surname, Phone: phone, Email: email, Course: course, Age: age})
		}

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		tpl.ExecuteTemplate(w, "admin.html", users)
		return
	}
	log.Print("incorrect password")
	tpl.ExecuteTemplate(w, "admin-login.html", "Incorrect Password")
}

func sendEverybody(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		subject := r.FormValue("subject")
		message := r.FormValue("message")

		fmt.Print(subject, message+"\r\n\r\n")
		stmt := "SELECT email FROM users WHERE email != 'admin@gmail.com';"
		rows, err := DB.Query(stmt)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		// Iterate over the rows
		for rows.Next() {
			var email string
			if err := rows.Scan(&email); err != nil {
				log.Fatal(err)
			}
			from := "ansish2005@gmail.com"
			password := "zqin vdaw xnxi luuk"
			to := []string{email}
			host := "smtp.gmail.com"
			port := "587"
			address := host + ":" + port
			subject := subject + "\r\n\r\n"
			body := message
			fmt.Println("body: ", body)
			message := []byte(subject + body)

			auth := smtp.PlainAuth("", from, password, host)
			err := smtp.SendMail(address, auth, from, to, message)
			if err != nil {
				w.Write([]byte("Something went wrong"))
			}
		}

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			// handle error
			log.Fatal(err)
		}
		http.Redirect(w, r, "/admin", http.StatusFound)
	}
}
