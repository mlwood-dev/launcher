// Code generated by mockery v2.21.1. DO NOT EDIT.

package tuf

import (
	mock "github.com/stretchr/testify/mock"
	data "github.com/theupdateframework/go-tuf/data"
)

// mockLibrarian is an autogenerated mock type for the librarian type
type mockLibrarian struct {
	mock.Mock
}

// AddToLibrary provides a mock function with given fields: binary, targetFilename, targetMetadata
func (_m *mockLibrarian) AddToLibrary(binary autoupdatableBinary, targetFilename string, targetMetadata data.TargetFileMeta) error {
	ret := _m.Called(binary, targetFilename, targetMetadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(autoupdatableBinary, string, data.TargetFileMeta) error); ok {
		r0 = rf(binary, targetFilename, targetMetadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TidyLibrary provides a mock function with given fields:
func (_m *mockLibrarian) TidyLibrary() {
	_m.Called()
}

type mockConstructorTestingTnewMockLibrarian interface {
	mock.TestingT
	Cleanup(func())
}

// newMockLibrarian creates a new instance of mockLibrarian. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockLibrarian(t mockConstructorTestingTnewMockLibrarian) *mockLibrarian {
	mock := &mockLibrarian{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
