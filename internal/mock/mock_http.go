// Automatically generated by MockGen. DO NOT EDIT!
// Source: http.go

package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Passworder interface
type MockPassworder struct {
	ctrl     *gomock.Controller
	recorder *_MockPassworderRecorder
}

// Recorder for MockPassworder (not exported)
type _MockPassworderRecorder struct {
	mock *MockPassworder
}

func NewMockPassworder(ctrl *gomock.Controller) *MockPassworder {
	mock := &MockPassworder{ctrl: ctrl}
	mock.recorder = &_MockPassworderRecorder{mock}
	return mock
}

func (_m *MockPassworder) EXPECT() *_MockPassworderRecorder {
	return _m.recorder
}

func (_m *MockPassworder) Password(minLength int, specialChars int, numbers int) string {
	ret := _m.ctrl.Call(_m, "Password", minLength, specialChars, numbers)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockPassworderRecorder) Password(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Password", arg0, arg1, arg2)
}