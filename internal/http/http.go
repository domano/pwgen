// Package http provides the web handlers and functions for our service
package http

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
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
	Password(minLength, specialChars, numbers int) string
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

	pw, err := ph.password(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

func (ph *PasswordHandler) password(r *http.Request) (string, error) {
	// Get parameters from URL & validate them
	params := r.URL.Query()

	minLength, err := numberFromParams(params, paramMinLength)
	if err != nil {
		return "", errors.Wrap(err, "Could not read minLength parameter")
	}
	specialChars, err := numberFromParams(params, paramSpecialChars)
	if err != nil {
		return "", errors.Wrap(err, "Could not read special chars parameter")
	}
	numbers, err := numberFromParams(params, paramNumbers)
	if err != nil {
		return "", errors.Wrap(err, "Could not read numbers parameter")
	}

	pw := ph.Password(minLength, specialChars, numbers)

	return pw, nil
}

func numberFromParams(vals url.Values, name string) (int, error) {
	val := vals.Get(name)
	if val == "" {
		return 0, nil
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.Wrapf(err, "Query Parameter %s was no number, got %s instead", name, val)
	}
	return num, nil

}
