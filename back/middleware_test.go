package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/csrf"
	"github.com/stretchr/testify/assert"
)

func TestCsrfMiddleware2(t *testing.T) {
	// given
	var (
		token   string
		cookie  string
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get token from middleware
			token = csrf.Token(r)
			t.Log("now visit other handlers...")
		})
		m = CsrfMiddleware([]byte(`some-key-here`))(handler)
	)
	// required trick by csrf package
	csrf.Secure(false)

	// GET
	getReq := httptest.NewRequest("GET", "/csrf", nil)
	getRr := httptest.NewRecorder()

	// then
	m.ServeHTTP(
		getRr,
		getReq,
	)
	// get cookie from response
	cookie = getRr.Header().Get("Set-Cookie")

	// POST
	postReq := httptest.NewRequest("POST", "/csrf", nil)
	postRr := httptest.NewRecorder()

	// WARN: must set both cookie and token, or you won't get 200 !!
	postReq.Header.Set("Cookie", cookie)
	postReq.Header.Set("X-CSRF-token", token)

	t.Log(cookie)
	t.Log(token)

	// then
	m.ServeHTTP(
		postRr,
		postReq,
	)

	// assert
	assert.Equal(t, 200, postRr.Code)
}
