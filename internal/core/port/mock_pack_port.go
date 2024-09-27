// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/port/pack_interface.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/port/pack_interface.go -destination=internal/core/port/mock_pack_port.go -package=port
//
// Package port is a generated GoMock package.
package port

import (
	domain "repartners/internal/core/domain"
	reflect "reflect"

	logrus "github.com/sirupsen/logrus"
	gomock "go.uber.org/mock/gomock"
)

// MockPackageService is a mock of PackageService interface.
type MockPackageService struct {
	ctrl     *gomock.Controller
	recorder *MockPackageServiceMockRecorder
}

// MockPackageServiceMockRecorder is the mock recorder for MockPackageService.
type MockPackageServiceMockRecorder struct {
	mock *MockPackageService
}

// NewMockPackageService creates a new mock instance.
func NewMockPackageService(ctrl *gomock.Controller) *MockPackageService {
	mock := &MockPackageService{ctrl: ctrl}
	mock.recorder = &MockPackageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageService) EXPECT() *MockPackageServiceMockRecorder {
	return m.recorder
}

// CalculatePackages mocks base method.
func (m *MockPackageService) CalculatePackages(items int, packagesCofig []int, logger *logrus.Entry) (domain.PackList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePackages", items, packagesCofig, logger)
	ret0, _ := ret[0].(domain.PackList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculatePackages indicates an expected call of CalculatePackages.
func (mr *MockPackageServiceMockRecorder) CalculatePackages(items, packagesCofig, logger any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePackages", reflect.TypeOf((*MockPackageService)(nil).CalculatePackages), items, packagesCofig, logger)
}
