package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

func createSession(w http.ResponseWriter, r *http.Request, name, email string) {
	session, err := store.Get(r, "biscuits")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	r.ParseForm()
	if name != "" && email != "" {
		session.Values["name"] = name
		session.Values["email"] = email
	}
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//tpl.ExecuteTemplate(w, "index.html", User{name, email})
}

func deleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "biscuits")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func getUser(r *http.Request) (*User, error) {
	session, _ := store.Get(r, "biscuits")
	if session.Values["name"] == nil {
		return nil, errors.New("nilUser")
	}
	fmt.Println("Name: ", session.Values["name"].(string))
	fmt.Println("Email: ", session.Values["email"].(string))
	return &User{Name: session.Values["name"].(string),
		Email: session.Values["email"].(string),
	}, nil
}
