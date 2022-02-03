package main

import (
	"github.com/gorilla/csrf"
	"net/http"
)

func CsrfMiddleware(authKey []byte) func(next http.Handler) http.Handler {
	return csrf.Protect(authKey)
}
