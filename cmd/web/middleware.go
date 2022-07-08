package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit this page")
		next.ServeHTTP(w, r)
	})
}

// NOSurf prevents CSRF to all POST requests
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteStrictMode,
	})
	return crsfHandler
}

//SessionLoad loads and serve the session on any request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
