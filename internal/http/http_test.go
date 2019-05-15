package http

import (
	"github.com/domano/pwgen/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPasswordHandler_ServeHTTP(t *testing.T) {
	testCases := []struct {
		desc string

		//given
		method            string
		queryParams       map[string]string
		returnedPasswords []string

		// expect
		expectedResponse      int
		expectedBody          string
		expectedContentLength int
	}{
		{
			desc:                  "GET, no params",
			method:                http.MethodGet,
			queryParams:           nil,
			returnedPasswords:     []string{""},
			expectedResponse:      http.StatusOK,
			expectedBody:          "[\"\"]",
			expectedContentLength: 4,
		},
		{
			desc:                  "HEAD, no params",
			method:                http.MethodHead,
			queryParams:           nil,
			returnedPasswords:     []string{""},
			expectedResponse:      http.StatusOK,
			expectedBody:          "",
			expectedContentLength: 4,
		},
		{
			desc:              "POST, no params",
			method:            http.MethodPost,
			queryParams:       nil,
			returnedPasswords: nil,
			expectedResponse:  http.StatusMethodNotAllowed,
			expectedBody:      "",
		},
		{
			desc:              "PUT, no params",
			method:            http.MethodPut,
			queryParams:       nil,
			returnedPasswords: nil,
			expectedResponse:  http.StatusMethodNotAllowed,
			expectedBody:      "",
		},
		{
			desc:              "DELETE, no params",
			method:            http.MethodDelete,
			queryParams:       nil,
			returnedPasswords: nil,
			expectedResponse:  http.StatusMethodNotAllowed,
			expectedBody:      "",
		},
		{
			desc:              "PATCH, no params",
			method:            http.MethodPatch,
			queryParams:       nil,
			returnedPasswords: nil,
			expectedResponse:  http.StatusMethodNotAllowed,
			expectedBody:      "",
		},
		{
			desc:                  "GET, minlength 1",
			method:                http.MethodGet,
			queryParams:           map[string]string{paramMinLength: "1"},
			returnedPasswords:     []string{"o"},
			expectedResponse:      http.StatusOK,
			expectedBody:          "[\"o\"]",
			expectedContentLength: 5,
		},
		{
			desc:                  "GET, numbers 1",
			method:                http.MethodGet,
			queryParams:           map[string]string{paramNumbers: "1"},
			returnedPasswords:     []string{"1"},
			expectedResponse:      http.StatusOK,
			expectedBody:          "[\"1\"]",
			expectedContentLength: 5,
		},
		{
			desc:                  "GET, specialchars 1",
			method:                http.MethodGet,
			queryParams:           map[string]string{paramSpecialChars: "1"},
			returnedPasswords:     []string{"!"},
			expectedResponse:      http.StatusOK,
			expectedBody:          "[\"!\"]",
			expectedContentLength: 5,
		},
		{
			desc:                  "GET, minLength 3, specialchars 1, numbers 1",
			method:                http.MethodGet,
			queryParams:           map[string]string{paramMinLength: "1", paramNumbers: "1", paramSpecialChars: "1"},
			returnedPasswords:     []string{"a1!"},
			expectedResponse:      http.StatusOK,
			expectedBody:          "[\"a1!\"]",
			expectedContentLength: 7,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// given a mock controller
			ctrl := gomock.NewController(t)

			// and a mocked password generator
			mockPassworder := mock.NewMockPassworder(ctrl)

			// and our handler
			ph := &PasswordHandler{mockPassworder}

			// and a recorder for our response
			rc := httptest.NewRecorder()

			// and a test request
			req, _ := http.NewRequest(tC.method, "", nil)
			for k, v := range tC.queryParams {
				req.URL.Query().Set(k, v)
			}

			// expect calls to the password generator
			passwordCall := mockPassworder.EXPECT().Password(gomock.Any(), gomock.Any(), gomock.Any())
			for _, pw := range tC.returnedPasswords {
				passwordCall.Return(pw)
			}
			passwordCall.Times(len(tC.returnedPasswords))

			// when our endpoint is called
			ph.ServeHTTP(rc, req)

			// then
			assert.Equal(t, tC.expectedResponse, rc.Code)
			assert.Contains(t, rc.Body.String(), tC.expectedBody)
			contentLength, _ := strconv.Atoi(rc.Header().Get("Content-Length"))
			assert.Equal(t, tC.expectedContentLength, contentLength)
			assert.Equal(t, len(tC.expectedBody), rc.Body.Len())
		})
	}
}
