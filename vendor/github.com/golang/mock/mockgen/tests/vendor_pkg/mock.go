// Code generated by MockGen. DO NOT EDIT.
// Source: a (interfaces: Ifc)

// Package vendor_pkg is a generated GoMock package.
package vendor_pkg

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIfc is a mock of Ifc interface
type MockIfc struct {
	ctrl     *gomock.Controller
	recorder *MockIfcMockRecorder
}

// MockIfcMockRecorder is the mock recorder for MockIfc
type MockIfcMockRecorder struct {
	mock *MockIfc
}

// NewMockIfc creates a new mock instance
func NewMockIfc(ctrl *gomock.Controller) *MockIfc {
	mock := &MockIfc{ctrl: ctrl}
	mock.recorder = &MockIfcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIfc) EXPECT() *MockIfcMockRecorder {
	return m.recorder
}

// A mocks base method
func (m *MockIfc) A(arg0 string) string {
	ret := m.ctrl.Call(m, "A", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// A indicates an expected call of A
func (mr *MockIfcMockRecorder) A(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "A", reflect.TypeOf((*MockIfc)(nil).A), arg0)
}

// B mocks base method
func (m *MockIfc) B(arg0 int) int {
	ret := m.ctrl.Call(m, "B", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// B indicates an expected call of B
func (mr *MockIfcMockRecorder) B(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "B", reflect.TypeOf((*MockIfc)(nil).B), arg0)
}

// C mocks base method
func (m *MockIfc) C(arg0 chan int) chan int {
	ret := m.ctrl.Call(m, "C", arg0)
	ret0, _ := ret[0].(chan int)
	return ret0
}

// C indicates an expected call of C
func (mr *MockIfcMockRecorder) C(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockIfc)(nil).C), arg0)
}

// D mocks base method
func (m *MockIfc) D(arg0 interface{}) {
	m.ctrl.Call(m, "D", arg0)
}

// D indicates an expected call of D
func (mr *MockIfcMockRecorder) D(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "D", reflect.TypeOf((*MockIfc)(nil).D), arg0)
}

// E mocks base method
func (m *MockIfc) E(arg0 map[string]interface{}) {
	m.ctrl.Call(m, "E", arg0)
}

// E indicates an expected call of E
func (mr *MockIfcMockRecorder) E(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "E", reflect.TypeOf((*MockIfc)(nil).E), arg0)
}

// F mocks base method
func (m *MockIfc) F(arg0 []float64) {
	m.ctrl.Call(m, "F", arg0)
}

// F indicates an expected call of F
func (mr *MockIfcMockRecorder) F(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "F", reflect.TypeOf((*MockIfc)(nil).F), arg0)
}
