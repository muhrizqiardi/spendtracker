// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/user.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/muhrizqiardi/spendtracker/internal/database/model"
	dto "github.com/muhrizqiardi/spendtracker/internal/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// DeleteOneByID mocks base method.
func (m *MockUserService) DeleteOneByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOneByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOneByID indicates an expected call of DeleteOneByID.
func (mr *MockUserServiceMockRecorder) DeleteOneByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOneByID", reflect.TypeOf((*MockUserService)(nil).DeleteOneByID), id)
}

// GetOneByEmail mocks base method.
func (m *MockUserService) GetOneByEmail(email string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneByEmail", email)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneByEmail indicates an expected call of GetOneByEmail.
func (mr *MockUserServiceMockRecorder) GetOneByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneByEmail", reflect.TypeOf((*MockUserService)(nil).GetOneByEmail), email)
}

// GetOneByID mocks base method.
func (m *MockUserService) GetOneByID(id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneByID", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneByID indicates an expected call of GetOneByID.
func (mr *MockUserServiceMockRecorder) GetOneByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneByID", reflect.TypeOf((*MockUserService)(nil).GetOneByID), id)
}

// Register mocks base method.
func (m *MockUserService) Register(payload dto.RegisterUserDTO) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", payload)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceMockRecorder) Register(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserService)(nil).Register), payload)
}

// UpdateOneByID mocks base method.
func (m *MockUserService) UpdateOneByID(id int, payload dto.UpdateUserDTO) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneByID", id, payload)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOneByID indicates an expected call of UpdateOneByID.
func (mr *MockUserServiceMockRecorder) UpdateOneByID(id, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneByID", reflect.TypeOf((*MockUserService)(nil).UpdateOneByID), id, payload)
}