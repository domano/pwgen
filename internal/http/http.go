// Package http provides the web handlers and functions for our service
package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PasswordHandler struct {
	Passworder
}

// Constants for the available query params
const paramMinLength = "minLength"
const paramSpecialChars = "specialChars"
const paramNumbers = "numbers"

// Passworder provides us with a Password function to generate passwords
type Passworder interface {
	Password(minLength, specialChars, numbers string) string
}

func NewPasswordHandler(p Passworder) *PasswordHandler {
	return &PasswordHandler{p}
}

func (ph *PasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// As we only provide passwords currently nothing except GET is supported
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pw := ph.password(r)

	// Write password as json response, implicit 200 if write succeeds
	body, err := json.Marshal([]string{pw})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(body)))

	// No Body for HEAD requests
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ph *PasswordHandler) password(r *http.Request) string {
	// Get parameters from URL
	params := r.URL.Query()
	minLength := params.Get(paramMinLength)
	specialChars := params.Get(paramSpecialChars)
	numbers := params.Get(paramNumbers)
	pw := ph.Password(minLength, specialChars, numbers)
	return pw
}
