// Automatically generated by MockGen. DO NOT EDIT!
// Source: wwqdrh/handbook/tools/profile/mockgen (interfaces: GetSetter)

package mockgen

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of GetSetter interface
type MockGetSetter struct {
	ctrl     *gomock.Controller
	recorder *_MockGetSetterRecorder
}

// Recorder for MockGetSetter (not exported)
type _MockGetSetterRecorder struct {
	mock *MockGetSetter
}

func NewMockGetSetter(ctrl *gomock.Controller) *MockGetSetter {
	mock := &MockGetSetter{ctrl: ctrl}
	mock.recorder = &_MockGetSetterRecorder{mock}
	return mock
}

func (_m *MockGetSetter) EXPECT() *_MockGetSetterRecorder {
	return _m.recorder
}

func (_m *MockGetSetter) Get(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockGetSetterRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockGetSetter) Set(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Set", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockGetSetterRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Set", arg0, arg1)
}
