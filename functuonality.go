package main

import (
	"net/http"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "biscuits")
		if session.Values["name"] == nil {
			http.Redirect(w, r, "/login", 302)
		}
		HandlerFunc.ServeHTTP(w, r)
	}
}
