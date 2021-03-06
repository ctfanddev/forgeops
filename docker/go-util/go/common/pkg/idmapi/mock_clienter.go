// Code generated by mockery v1.0.0. DO NOT EDIT.

package idmapi

import mock "github.com/stretchr/testify/mock"
import user "github.com/ForgeCloud/saas/tree/master/services/go/common/pkg/models/user"

// MockClienter is an autogenerated mock type for the MockClienter type
type MockClienter struct {
	mock.Mock
}

// CreateTeamMember provides a mock function with given fields: email, password, firstName, lastName
func (_m *MockClienter) CreateTeamMember(email string, password string, firstName string, lastName string) (string, error) {
	ret := _m.Called(email, password, firstName, lastName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string) string); ok {
		r0 = rf(email, password, firstName, lastName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(email, password, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTeamMember provides a mock function with given fields: id
func (_m *MockClienter) DeleteTeamMember(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAppUser provides a mock function with given fields: id
func (_m *MockClienter) GetAppUser(id string) (*user.User, error) {
	ret := _m.Called(id)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamMember provides a mock function with given fields: id
func (_m *MockClienter) GetTeamMember(id string) (*user.User, error) {
	ret := _m.Called(id)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTeamMember provides a mock function with given fields: id, email, password, firstName, lastName
func (_m *MockClienter) UpdateTeamMember(id string, email string, password string, firstName string, lastName string) error {
	ret := _m.Called(id, email, password, firstName, lastName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) error); ok {
		r0 = rf(id, email, password, firstName, lastName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
