package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
)

func token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	_, err := fmt.Fprint(w, r.FormValue("say"))
	if err != nil {
		return
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	_, err := w.Write([]byte("The time is: " + tm))
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()

	csrf.Secure(false) // dev only
	csrfKey := []byte("32-byte-long-auth-key")

	//csrfMiddleware := csrf.Protect(csrfKey)
	csrfMiddleware := CsrfMiddleware(csrfKey)

	mux.Handle("/api/token", csrfMiddleware(http.HandlerFunc(token)))

	// OPTIONS, POST
	mux.Handle("/api/post", csrfMiddleware(http.HandlerFunc(post)))

	// NOTE: HandleFunc
	mux.Handle("/time", http.HandlerFunc(timeHandler))

	fmt.Println("http://localhost:8000")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
