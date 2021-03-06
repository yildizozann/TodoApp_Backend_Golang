// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/todo/handler.go

// Package mock_todo is a generated GoMock package.
package mock_todo

import (
	context "context"
	reflect "reflect"
	todo "todo_app/internal/todo"

	gomock "github.com/golang/mock/gomock"
)

// MockHandlerService is a mock of HandlerService interface.
type MockHandlerService struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerServiceMockRecorder
}

// MockHandlerServiceMockRecorder is the mock recorder for MockHandlerService.
type MockHandlerServiceMockRecorder struct {
	mock *MockHandlerService
}

// NewMockHandlerService creates a new mock instance.
func NewMockHandlerService(ctrl *gomock.Controller) *MockHandlerService {
	mock := &MockHandlerService{ctrl: ctrl}
	mock.recorder = &MockHandlerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandlerService) EXPECT() *MockHandlerServiceMockRecorder {
	return m.recorder
}

// AddNewToDo mocks base method.
func (m *MockHandlerService) AddNewToDo(ctx context.Context, request *todo.AddToDoRequest) (*todo.AddToDoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewToDo", ctx, request)
	ret0, _ := ret[0].(*todo.AddToDoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewToDo indicates an expected call of AddNewToDo.
func (mr *MockHandlerServiceMockRecorder) AddNewToDo(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewToDo", reflect.TypeOf((*MockHandlerService)(nil).AddNewToDo), ctx, request)
}

// GetAll mocks base method.
func (m *MockHandlerService) GetAll() *todo.GetAllResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*todo.GetAllResponse)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockHandlerServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockHandlerService)(nil).GetAll))
}
