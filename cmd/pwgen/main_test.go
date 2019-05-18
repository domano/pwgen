package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
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
	go run(nil)

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

func Test_parseConfig(t *testing.T) {
	// when
	err := parseConfig()

	// then
	assert.NoError(t, err)
	assert.Equal(t,cfg.GracePeriod,5*time.Second)
}

func Test_parseConfig_withError(t *testing.T) {
	os.Setenv("GRACE_PERIOD", "awkljdnalksd")
	// when
	err := parseConfig()

	// then
	assert.Error(t, err)
}

func Test_run_withError(t *testing.T) {
	// when we start with an empty config for the key and cert file
	err := run(nil)

	// then
	assert.Error(t, err)
}

func Test_run_graceful_Shutdown(t *testing.T) {
	// when
	parseConfig()
	stopChan := make(chan os.Signal, 1)
	stopChan <- os.Interrupt
	err := run(stopChan)

	// then
	assert.NoError(t, err)
}