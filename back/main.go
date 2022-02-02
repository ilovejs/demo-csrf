package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// debug
	csrf.Secure(false)

	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))

	api := r.PathPrefix("/api").
		Subrouter()

	api.Use(csrfMiddleware)

	api.HandleFunc("/token", token).Methods(http.MethodGet)
	api.HandleFunc("/post", post).Methods(http.MethodOptions) // プリフライトを受け付ける
	api.HandleFunc("/post", post).Methods(http.MethodPost)

	fmt.Println("http://localhost:8000")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}

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
