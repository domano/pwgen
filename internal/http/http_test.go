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
		desc        string
		method      string
		queryParams map[string]string
		returnedPasswords []string
		expectedResponse int
		expectedBody string
	}{
		{
			desc:        "GET, no params",
			method:      http.MethodGet,
			queryParams: nil,
			returnedPasswords:[]string{"one"},
			expectedResponse: http.StatusOK,
			expectedBody:"['one']",
		},
		{
			desc:        "HEAD, no params",
			method:      http.MethodGet,
			queryParams: nil,
			returnedPasswords:[]string{"one"},
			expectedResponse: http.StatusOK,
			expectedBody:"",
		},
		{
			desc:        "POST, no params",
			method:      http.MethodPost,
			queryParams: nil,
			returnedPasswords:nil,
			expectedResponse: http.StatusMethodNotAllowed,
			expectedBody:"",
		},
		{
			desc:        "PUT, no params",
			method:      http.MethodPut,
			queryParams: nil,
			returnedPasswords:nil,
			expectedResponse: http.StatusMethodNotAllowed,
			expectedBody:"",
		},
		{
			desc:        "DELETE, no params",
			method:      http.MethodDelete,
			queryParams: nil,
			returnedPasswords:nil,
			expectedResponse: http.StatusMethodNotAllowed,
			expectedBody:"",
		},
		{
			desc:        "PATCH, no params",
			method:      http.MethodPatch,
			queryParams: nil,
			returnedPasswords:nil,
			expectedResponse: http.StatusMethodNotAllowed,
			expectedBody:"",
		},
		{
			desc:        "GET, minlength 1",
			method:      http.MethodGet,
			queryParams: map[string]string{paramMinLength:"1"},
			returnedPasswords:[]string{"o"},
			expectedResponse: http.StatusOK,
			expectedBody:"['o']",
		},
		{
			desc:        "GET, numbers 1",
			method:      http.MethodGet,
			queryParams: map[string]string{paramNumbers:"1"},
			returnedPasswords:[]string{"['1']"},
			expectedResponse: http.StatusOK,
			expectedBody:"['1']",
		},
		{
			desc:        "GET, specialchars 1",
			method:      http.MethodGet,
			queryParams: map[string]string{paramSpecialChars:"1"},
			returnedPasswords:[]string{"['!']"},
			expectedResponse: http.StatusOK,
			expectedBody:"['!']",
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
			passwordCall := mockPassworder.EXPECT().Password(gomock.Any(),gomock.Any(),gomock.Any())
			for _,pw := range tC.returnedPasswords {
				passwordCall.Return(pw)
			}
			passwordCall.Times(len(tC.returnedPasswords))

			// when our endpoint is called
			ph.ServeHTTP(rc, req)

			// then
			assert.Equal(t,tC.expectedResponse, rc.Code)
			assert.Equal(t,tC.expectedBody,rc.Body.String())
			contentLength, err := strconv.Atoi(rc.Header().Get("Content-Length"))
			assert.NoError(t, err)
			assert.Equal(t,len(tC.expectedBody), contentLength)
		})
	}
}
