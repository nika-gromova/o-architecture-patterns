// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nika-gromova/o-architecture-patterns/game/change_velocity (interfaces: RotatingAndMovingObject)
//
// Generated by this command:
//
//	mockgen -destination mr_obj_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/change_velocity RotatingAndMovingObject
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	base "github.com/nika-gromova/o-architecture-patterns/game/base"
	gomock "go.uber.org/mock/gomock"
)

// MockRotatingAndMovingObject is a mock of RotatingAndMovingObject interface.
type MockRotatingAndMovingObject struct {
	ctrl     *gomock.Controller
	recorder *MockRotatingAndMovingObjectMockRecorder
	isgomock struct{}
}

// MockRotatingAndMovingObjectMockRecorder is the mock recorder for MockRotatingAndMovingObject.
type MockRotatingAndMovingObjectMockRecorder struct {
	mock *MockRotatingAndMovingObject
}

// NewMockRotatingAndMovingObject creates a new mock instance.
func NewMockRotatingAndMovingObject(ctrl *gomock.Controller) *MockRotatingAndMovingObject {
	mock := &MockRotatingAndMovingObject{ctrl: ctrl}
	mock.recorder = &MockRotatingAndMovingObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRotatingAndMovingObject) EXPECT() *MockRotatingAndMovingObjectMockRecorder {
	return m.recorder
}

// GetVelocityVector mocks base method.
func (m *MockRotatingAndMovingObject) GetVelocityVector() (base.Vector, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVelocityVector")
	ret0, _ := ret[0].(base.Vector)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetVelocityVector indicates an expected call of GetVelocityVector.
func (mr *MockRotatingAndMovingObjectMockRecorder) GetVelocityVector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVelocityVector", reflect.TypeOf((*MockRotatingAndMovingObject)(nil).GetVelocityVector))
}

// SetVelocityVector mocks base method.
func (m *MockRotatingAndMovingObject) SetVelocityVector(vector base.Vector) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVelocityVector", vector)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetVelocityVector indicates an expected call of SetVelocityVector.
func (mr *MockRotatingAndMovingObjectMockRecorder) SetVelocityVector(vector any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVelocityVector", reflect.TypeOf((*MockRotatingAndMovingObject)(nil).SetVelocityVector), vector)
}
