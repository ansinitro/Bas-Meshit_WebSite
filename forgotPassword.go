package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
)

type TempData struct {
	Username   string
	Email      string
	AuthInfo   string
	ErrMessage string
	Message    string
}

func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "forgot-password.html", nil)
		return
	}
	fmt.Println(r.Method)
	fmt.Println(r.Body)
	r.ParseForm()
	email := r.FormValue("email")
	var td TempData
	td.ErrMessage = "Sorry, this was an issue recovering account, please try again"
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err: ", err)
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			fmt.Println("failed rollback, rollBack: ", rollBackErr)
		}
		tpl.ExecuteTemplate(w, "forgot-password.html", td.ErrMessage)
		return
	}
	defer tx.Rollback()
	var name string
	fmt.Println(email)
	row := DB.QueryRow("SELECT email, name FROM users WHERE email = $1", email)
	err = row.Scan(&email, &name)
	if err != nil {
		fmt.Println("email not found in db")
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			fmt.Println("failed rollback, rollBack: ", rollBackErr)
		}
		tpl.ExecuteTemplate(w, "forgot-password.html", td.ErrMessage)
		return
	}
	now := time.Now()
	timeout := now.Add(time.Minute * 45)
	rand.Seed(time.Now().UnixNano())

	alphaNumRunes := []rune("qwerQWERtyuiTYUIopOPasdfASDFghjkGHJKlLzxcvZXCVbnmBNM1234567890")
	emailVerRandRune := make([]rune, 64)
	for i := 0; i < 64; i++ {
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes))]
	}
	emailVerPassword := string(emailVerRandRune)
	emailVerPasswordHash, err := bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcript err: ", err)
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			fmt.Println("failed rollback, rollBack: ", rollBackErr)
		}
		tpl.ExecuteTemplate(w, "forgot-password.html", td.ErrMessage)
		return
	}
	var updateEmailVerStmt *sql.Stmt
	updateEmailVerStmt, err = tx.Prepare("UPDATE email_ver_hash SET ver_hash = $1, timeout = $2 WHERE email = $3")
	if err != nil {
		fmt.Println("error preparing statement, err: ", err)
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			fmt.Println("failed rollback, rollBack: ", rollBackErr)
		}
		tpl.ExecuteTemplate(w, "forgot-password.html", td.ErrMessage)
		return
	}
	defer updateEmailVerStmt.Close()
	emailVerPasswordHashStr := string(emailVerPasswordHash)
	result, err := updateEmailVerStmt.Exec(emailVerPasswordHashStr, timeout, email)
	rowAff, _ := result.RowsAffected()
	if err != nil || rowAff != 1 {
		fmt.Println("error inserting new user, err: ", err)
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			fmt.Println("failed rollback, rollBack: ", rollBackErr)
		}
		tpl.ExecuteTemplate(w, "forgot-password.html", td.ErrMessage)
		return
	}

	from := "ansish2005@gmail.com"
	password := "zqin vdaw xnxi luuk"
	to := []string{email}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	subject := "Subject: Email Verification Code\r\n\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://www.mysite.com/forgotPasswordChange?u=" + name + "&evpw=" + emailVerPassword + "\">Change Password</a></body>"
	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, host)
	err = smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("error sending reset password email, err: ", err)
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			fmt.Println("failed rollback, rollBack: ", rollBackErr)
		}
		tpl.ExecuteTemplate(w, "forgot-password.html", td.ErrMessage)
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error commiting changes, commitErr: ", commitErr)
		tpl.ExecuteTemplate(w, "forgot-password.html", td.ErrMessage)
	}
	tpl.ExecuteTemplate(w, "forgot-password-mail.html", nil)
}
