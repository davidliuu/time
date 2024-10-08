/*
Copyright (c) Facebook, Inc. and its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: phc/phc.go

// Package phc is a generated GoMock package.
package phc

import (
	os "os"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockDeviceController is a mock of DeviceController interface.
type MockDeviceController struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceControllerMockRecorder
}

// MockDeviceControllerMockRecorder is the mock recorder for MockDeviceController.
type MockDeviceControllerMockRecorder struct {
	mock *MockDeviceController
}

// NewMockDeviceController creates a new mock instance.
func NewMockDeviceController(ctrl *gomock.Controller) *MockDeviceController {
	mock := &MockDeviceController{ctrl: ctrl}
	mock.recorder = &MockDeviceControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceController) EXPECT() *MockDeviceControllerMockRecorder {
	return m.recorder
}

// File mocks base method.
func (m *MockDeviceController) File() *os.File {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "File")
	ret0, _ := ret[0].(*os.File)
	return ret0
}

// File indicates an expected call of File.
func (mr *MockDeviceControllerMockRecorder) File() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "File", reflect.TypeOf((*MockDeviceController)(nil).File))
}

// Time mocks base method.
func (m *MockDeviceController) Time() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Time indicates an expected call of Time.
func (mr *MockDeviceControllerMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockDeviceController)(nil).Time))
}

// setPTPPerout mocks base method.
func (m *MockDeviceController) setPTPPerout(req PTPPeroutRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setPTPPerout", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// setPTPPerout indicates an expected call of setPTPPerout.
func (mr *MockDeviceControllerMockRecorder) setPTPPerout(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setPTPPerout", reflect.TypeOf((*MockDeviceController)(nil).setPTPPerout), req)
}

// setPinFunc mocks base method.
func (m *MockDeviceController) setPinFunc(index uint, pf PinFunc, ch uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setPinFunc", index, pf, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// setPinFunc indicates an expected call of setPinFunc.
func (mr *MockDeviceControllerMockRecorder) setPinFunc(index, pf, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setPinFunc", reflect.TypeOf((*MockDeviceController)(nil).setPinFunc), index, pf, ch)
}
