// Package http provides the web handlers and functions for our service
package http

import "net/http"

type PasswordHandler struct {
	Passworder
}

// Constants for the available query params
const paramMinLength = "minLength"
const paramSpecialChars = "specialChars"
const paramNumbers = "numbers"

// Passworder provides us with a Password function to generate passwords
type Passworder interface {
	Password(minLength, specialChars, numbers int) string
}

func (ph *PasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// As we only provide passwords currently nothing except GET is supported
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
