package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf creates CSRF protection token
func NoSurf(next http.Handler) http.Handler {
	csrfToken := nosurf.New(next)

	csrfToken.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfToken
}

// LoadSession loads and saves session
func LoadSession(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
