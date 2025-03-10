// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nika-gromova/o-architecture-patterns/game/macro_command/move (interfaces: MovingWithFuelObj)
//
// Generated by this command:
//
//	mockgen -destination mf_obj_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/macro_command/move MovingWithFuelObj
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	base "github.com/nika-gromova/o-architecture-patterns/game/base"
	gomock "go.uber.org/mock/gomock"
)

// MockMovingWithFuelObj is a mock of MovingWithFuelObj interface.
type MockMovingWithFuelObj struct {
	ctrl     *gomock.Controller
	recorder *MockMovingWithFuelObjMockRecorder
	isgomock struct{}
}

// MockMovingWithFuelObjMockRecorder is the mock recorder for MockMovingWithFuelObj.
type MockMovingWithFuelObjMockRecorder struct {
	mock *MockMovingWithFuelObj
}

// NewMockMovingWithFuelObj creates a new mock instance.
func NewMockMovingWithFuelObj(ctrl *gomock.Controller) *MockMovingWithFuelObj {
	mock := &MockMovingWithFuelObj{ctrl: ctrl}
	mock.recorder = &MockMovingWithFuelObjMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovingWithFuelObj) EXPECT() *MockMovingWithFuelObjMockRecorder {
	return m.recorder
}

// GetFuel mocks base method.
func (m *MockMovingWithFuelObj) GetFuel() (base.FuelInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFuel")
	ret0, _ := ret[0].(base.FuelInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetFuel indicates an expected call of GetFuel.
func (mr *MockMovingWithFuelObjMockRecorder) GetFuel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFuel", reflect.TypeOf((*MockMovingWithFuelObj)(nil).GetFuel))
}

// GetLocation mocks base method.
func (m *MockMovingWithFuelObj) GetLocation() (base.Vector, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation")
	ret0, _ := ret[0].(base.Vector)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetLocation indicates an expected call of GetLocation.
func (mr *MockMovingWithFuelObjMockRecorder) GetLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*MockMovingWithFuelObj)(nil).GetLocation))
}

// GetVelocity mocks base method.
func (m *MockMovingWithFuelObj) GetVelocity() (base.Vector, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVelocity")
	ret0, _ := ret[0].(base.Vector)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetVelocity indicates an expected call of GetVelocity.
func (mr *MockMovingWithFuelObjMockRecorder) GetVelocity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVelocity", reflect.TypeOf((*MockMovingWithFuelObj)(nil).GetVelocity))
}

// SetFuel mocks base method.
func (m *MockMovingWithFuelObj) SetFuel(f base.FuelInfo) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFuel", f)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetFuel indicates an expected call of SetFuel.
func (mr *MockMovingWithFuelObjMockRecorder) SetFuel(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFuel", reflect.TypeOf((*MockMovingWithFuelObj)(nil).SetFuel), f)
}

// SetLocation mocks base method.
func (m *MockMovingWithFuelObj) SetLocation(arg0 base.Vector) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLocation", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetLocation indicates an expected call of SetLocation.
func (mr *MockMovingWithFuelObjMockRecorder) SetLocation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLocation", reflect.TypeOf((*MockMovingWithFuelObj)(nil).SetLocation), arg0)
}
