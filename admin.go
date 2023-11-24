package main

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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
		log.Printf("username doesn't exists, err:", err)
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
