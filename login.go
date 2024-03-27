package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	stmt := "SELECT name, hash FROM users WHERE email = $1"
	row := DB.QueryRow(stmt, email)
	var hash string
	err = row.Scan(&name, &hash)
	if err == sql.ErrNoRows {
		log.Printf("username doesn't exists, %v:", err)
		w.WriteHeader(http.StatusNotFound)
		tpl.ExecuteTemplate(w, "login.html", "Email not exist.")
		return
	}

	//hash from passwd
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	if err == nil {
		log.Print("successe login")
		createSession(w, r, name, email)
		//tpl.ExecuteTemplate(w, "index.html", "You succesufully logged in")
		w.WriteHeader(http.StatusOK)
		http.Redirect(w, r, "/index", 302)
		return
	}
	log.Print("incorrect password")
	w.WriteHeader(http.StatusUnauthorized)
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

	var tx *sql.Tx
	tx, err = DB.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err: ", err)
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		return
	}
	defer tx.Rollback()

	//check is user with same email exist
	stmt := "SELECT name FROM users WHERE email = $1"
	row := tx.QueryRow(stmt, email)
	var value string
	err = row.Scan(&value)
	if err != sql.ErrNoRows {
		log.Printf("username already exists, %v:", err)
		tpl.ExecuteTemplate(w, "login.html", "Username Already Exists")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}

	//hash from passwd
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("bcrypt err %v\n", err)
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}
	var insertStatement *sql.Stmt
	insertStatement, err = tx.Prepare("INSERT INTO users (name, email, hash, created_at, is_active) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Printf("error preparing statement %v\n", err)
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}
	defer insertStatement.Close()

	result, err := insertStatement.Exec(name, email, string(hash), time.Now(), false)
	rowsAff, err := result.RowsAffected()
	if err != nil || rowsAff != 1 {
		log.Printf("Error inserting new user")
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}

	// ver code
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(999999-100000) + 100000
	fmt.Println("random num: ", rn)
	var insertEmailVerStatement *sql.Stmt
	insertEmailVerStatement, err = tx.Prepare("INSERT INTO email_ver (username, email, ver_code) VALUES ($1, $2, $3)")
	if err != nil {
		fmt.Println("error preparing stmt: ", err)
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}
	defer insertEmailVerStatement.Close()
	result, err = insertEmailVerStatement.Exec(name, email, rn)
	rowsAff, _ = result.RowsAffected()
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting into email_ver: ", err)
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}

	err = emailVerCode(rn, email)
	if err != nil {
		fmt.Println("error email_ver code: ", err)
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error: ", err)
		tpl.ExecuteTemplate(w, "login.html", "There was issue registering, please try again.")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error commiting changes, rollbackErr: ", rollbackErr)
		}
		return
	}

	var m User
	m.Email = email
	tpl.ExecuteTemplate(w, "verify-email.html", m)
	fmt.Println("Everything good")
}

func verifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "verify-email.html", nil)
		return
	}
	r.ParseForm()
	email := r.FormValue("email")
	verCode := r.FormValue("vercode")
	fmt.Println(email, verCode)
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err: ", err)
		tpl.ExecuteTemplate(w, "verify-email.html", "Sorry, there was an issue verifying email, please try again.")
		return
	}
	defer tx.Rollback()
	stmt := "SELECT ver_code FROM email_ver WHERE email = $1"
	row := tx.QueryRow(stmt, email)
	var dbCode string
	err = row.Scan(&dbCode)
	if err != nil {
		fmt.Println("error scanning verCode, err: ", err)
		var m User
		m.Email, m.ErrMessage = email, "Sorry, there was an issue verifying email, please try again."
		tpl.ExecuteTemplate(w, "verify-email.html", m)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
		}
		return
	}

	fmt.Println(verCode, dbCode)
	if verCode == dbCode {
		stmt := "UPDATE users SET is_active = true WHERE email = $1"
		updateIsActiveStmt, err := tx.Prepare(stmt)
		if err != nil {
			fmt.Println("error preparing updateIsActiveStmt, err: ", err)
			var m User
			m.Email, m.ErrMessage = email, "Sorry, there was an issue verifying email, please try again."
			tpl.ExecuteTemplate(w, "verify-email.html", m)
			return
		}
		defer updateIsActiveStmt.Close()
		var result sql.Result
		result, err = updateIsActiveStmt.Exec(email)
		rowAff, err := result.RowsAffected()
		if err != nil || rowAff != 1 {
			fmt.Println("error inserting new user, err: ", err)
			var m User
			m.Email, m.ErrMessage = email, "Sorry, there was an issue verifying email, please try again."
			tpl.ExecuteTemplate(w, "verify-email.html", m)
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				fmt.Println("occured error with rollback, rollbackErr: ", rollbackErr)
			}
			return
		}
		tpl.ExecuteTemplate(w, "login.html", "email verified, go ahead and login!")
		tx.Commit()
		return
	}
	var m User
	m.Email, m.ErrMessage = email, "Sorry, there was an issue verifying email, please try again."
	if rollBackErr := tx.Rollback(); rollBackErr != nil {
		fmt.Println("occured error with rollback, rollbackErr: ", err)
	}
	tpl.ExecuteTemplate(w, "verify-email.html", m)
}

func emailVerCode(rn int, toEmail string) error {
	from := "ansish2005@gmail.com"
	password := "zqin vdaw xnxi luuk"
	to := []string{toEmail}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	subject := "Subject: Email Verification Code\r\n\r\n"
	verCode := strconv.Itoa(rn)
	body := "verification code: " + verCode
	fmt.Println("body: ", body)
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(address, auth, from, to, message)

	return err
}

// IsValidEmail checks if the given email address is valid.
func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
