package http

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingHandlerFunc_Access_Log(t *testing.T) {
	// given some writer to test our log output
	rc :=httptest.NewRecorder()
	logBuffer := bytes.NewBufferString("")
	logrus.SetOutput(logBuffer)

	// and a request
	req := httptest.NewRequest("GET", "https://www.test.de/test", nil)

	// when
	LoggingHandlerFunc(nil)(rc, req)

	//then
	body := logBuffer.String()
	assert.Contains(t, body, "path=/test")
	assert.Contains(t, body, "response_code=200")
	assert.Contains(t, body, "type=access")
	assert.Contains(t, body, "host=www.test.de")
}

func TestLoggingHandlerFunc_Access_Log_With_Next(t *testing.T) {
	// given some writer to test our log output
	rc :=httptest.NewRecorder()
	logBuffer := bytes.NewBufferString("")
	logrus.SetOutput(logBuffer)

	// and a test handler to check if it was called as next
	var called bool
	testHandlerFunc := func (_ http.ResponseWriter, _ *http.Request) {
		called = true
	}
	// and a request
	req := httptest.NewRequest("GET", "https://www.test.de/test", nil)

	// when we create and call our logging handler function
	LoggingHandlerFunc(http.HandlerFunc(testHandlerFunc))(rc, req)

	//then
	body := logBuffer.String()
	assert.Contains(t, body, "path=/test")
	assert.Contains(t, body, "response_code=200")
	assert.Contains(t, body, "type=access")
	assert.Contains(t, body, "host=www.test.de")
	assert.True(t, called)
}
