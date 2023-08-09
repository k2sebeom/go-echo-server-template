// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	schema "github.com/k2sebeom/go-echo-server-template/models/schema"
)

// MockIUserRepository is an autogenerated mock type for the IUserRepository type
type MockIUserRepository struct {
	mock.Mock
}

type MockIUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIUserRepository) EXPECT() *MockIUserRepository_Expecter {
	return &MockIUserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *MockIUserRepository) CreateUser(user *schema.User) (*schema.User, error) {
	ret := _m.Called(user)

	var r0 *schema.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*schema.User) (*schema.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*schema.User) *schema.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*schema.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockIUserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user *schema.User
func (_e *MockIUserRepository_Expecter) CreateUser(user interface{}) *MockIUserRepository_CreateUser_Call {
	return &MockIUserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *MockIUserRepository_CreateUser_Call) Run(run func(user *schema.User)) *MockIUserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*schema.User))
	})
	return _c
}

func (_c *MockIUserRepository_CreateUser_Call) Return(_a0 *schema.User, _a1 error) *MockIUserRepository_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_CreateUser_Call) RunAndReturn(run func(*schema.User) (*schema.User, error)) *MockIUserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *MockIUserRepository) GetUserByEmail(email string) (*schema.User, error) {
	ret := _m.Called(email)

	var r0 *schema.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*schema.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *schema.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type MockIUserRepository_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockIUserRepository_Expecter) GetUserByEmail(email interface{}) *MockIUserRepository_GetUserByEmail_Call {
	return &MockIUserRepository_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", email)}
}

func (_c *MockIUserRepository_GetUserByEmail_Call) Run(run func(email string)) *MockIUserRepository_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockIUserRepository_GetUserByEmail_Call) Return(_a0 *schema.User, _a1 error) *MockIUserRepository_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_GetUserByEmail_Call) RunAndReturn(run func(string) (*schema.User, error)) *MockIUserRepository_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: id
func (_m *MockIUserRepository) GetUserById(id uint) (*schema.User, error) {
	ret := _m.Called(id)

	var r0 *schema.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*schema.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *schema.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type MockIUserRepository_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - id uint
func (_e *MockIUserRepository_Expecter) GetUserById(id interface{}) *MockIUserRepository_GetUserById_Call {
	return &MockIUserRepository_GetUserById_Call{Call: _e.mock.On("GetUserById", id)}
}

func (_c *MockIUserRepository_GetUserById_Call) Run(run func(id uint)) *MockIUserRepository_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint))
	})
	return _c
}

func (_c *MockIUserRepository_GetUserById_Call) Return(_a0 *schema.User, _a1 error) *MockIUserRepository_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_GetUserById_Call) RunAndReturn(run func(uint) (*schema.User, error)) *MockIUserRepository_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserBySocialHash provides a mock function with given fields: hash
func (_m *MockIUserRepository) GetUserBySocialHash(hash string) (*schema.User, error) {
	ret := _m.Called(hash)

	var r0 *schema.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*schema.User, error)); ok {
		return rf(hash)
	}
	if rf, ok := ret.Get(0).(func(string) *schema.User); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_GetUserBySocialHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserBySocialHash'
type MockIUserRepository_GetUserBySocialHash_Call struct {
	*mock.Call
}

// GetUserBySocialHash is a helper method to define mock.On call
//   - hash string
func (_e *MockIUserRepository_Expecter) GetUserBySocialHash(hash interface{}) *MockIUserRepository_GetUserBySocialHash_Call {
	return &MockIUserRepository_GetUserBySocialHash_Call{Call: _e.mock.On("GetUserBySocialHash", hash)}
}

func (_c *MockIUserRepository_GetUserBySocialHash_Call) Run(run func(hash string)) *MockIUserRepository_GetUserBySocialHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockIUserRepository_GetUserBySocialHash_Call) Return(_a0 *schema.User, _a1 error) *MockIUserRepository_GetUserBySocialHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_GetUserBySocialHash_Call) RunAndReturn(run func(string) (*schema.User, error)) *MockIUserRepository_GetUserBySocialHash_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: user
func (_m *MockIUserRepository) UpdateUser(user schema.User) (*schema.User, error) {
	ret := _m.Called(user)

	var r0 *schema.User
	var r1 error
	if rf, ok := ret.Get(0).(func(schema.User) (*schema.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(schema.User) *schema.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schema.User)
		}
	}

	if rf, ok := ret.Get(1).(func(schema.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIUserRepository_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type MockIUserRepository_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - user schema.User
func (_e *MockIUserRepository_Expecter) UpdateUser(user interface{}) *MockIUserRepository_UpdateUser_Call {
	return &MockIUserRepository_UpdateUser_Call{Call: _e.mock.On("UpdateUser", user)}
}

func (_c *MockIUserRepository_UpdateUser_Call) Run(run func(user schema.User)) *MockIUserRepository_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(schema.User))
	})
	return _c
}

func (_c *MockIUserRepository_UpdateUser_Call) Return(_a0 *schema.User, _a1 error) *MockIUserRepository_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIUserRepository_UpdateUser_Call) RunAndReturn(run func(schema.User) (*schema.User, error)) *MockIUserRepository_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIUserRepository creates a new instance of MockIUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIUserRepository {
	mock := &MockIUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}