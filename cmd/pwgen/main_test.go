package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Test_run(t *testing.T) {
	// given a request for passwords
	req, _ := http.NewRequest("GET", "https://localhost:8443/passwords?minLength=10", nil)

	// and a valid test config
	cfg = config{
		"../../cert.pem",
		"../../key.unencrypted.pem",
		8443,
		5 * time.Second,
	}

	// and our started app
	go run()

	// and a Transport which accepts self signed certs
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// when we send a password request after a second
	<-time.After(time.Second)
	resp, err := http.DefaultClient.Do(req)

	// then we should get passwords in our response
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	var passwords []string
	err = json.NewDecoder(resp.Body).Decode(&passwords)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(passwords))
}
