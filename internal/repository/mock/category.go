// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/category.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/muhrizqiardi/spendtracker/internal/database/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCategoryRepository) Delete(id uint) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", id)
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryRepository)(nil).Delete), id)
}

// GetMany mocks base method.
func (m *MockCategoryRepository) GetMany(userID uint, limit, offset int) ([]model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", userID, limit, offset)
	ret0, _ := ret[0].([]model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockCategoryRepositoryMockRecorder) GetMany(userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockCategoryRepository)(nil).GetMany), userID, limit, offset)
}

// GetOneByID mocks base method.
func (m *MockCategoryRepository) GetOneByID(id uint) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneByID", id)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneByID indicates an expected call of GetOneByID.
func (mr *MockCategoryRepositoryMockRecorder) GetOneByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneByID", reflect.TypeOf((*MockCategoryRepository)(nil).GetOneByID), id)
}

// GetOneByName mocks base method.
func (m *MockCategoryRepository) GetOneByName(name string) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneByName", name)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneByName indicates an expected call of GetOneByName.
func (mr *MockCategoryRepositoryMockRecorder) GetOneByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneByName", reflect.TypeOf((*MockCategoryRepository)(nil).GetOneByName), name)
}

// Insert mocks base method.
func (m *MockCategoryRepository) Insert(userID uint, name string) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", userID, name)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockCategoryRepositoryMockRecorder) Insert(userID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCategoryRepository)(nil).Insert), userID, name)
}