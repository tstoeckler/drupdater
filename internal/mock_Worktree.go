// Code generated by mockery v2.53.3. DO NOT EDIT.

package internal

import (
	git "github.com/go-git/go-git/v5"
	mock "github.com/stretchr/testify/mock"

	plumbing "github.com/go-git/go-git/v5/plumbing"
)

// MockWorktree is an autogenerated mock type for the Worktree type
type MockWorktree struct {
	mock.Mock
}

type MockWorktree_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorktree) EXPECT() *MockWorktree_Expecter {
	return &MockWorktree_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: path
func (_m *MockWorktree) Add(path string) (plumbing.Hash, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 plumbing.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (plumbing.Hash, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) plumbing.Hash); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plumbing.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorktree_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockWorktree_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - path string
func (_e *MockWorktree_Expecter) Add(path interface{}) *MockWorktree_Add_Call {
	return &MockWorktree_Add_Call{Call: _e.mock.On("Add", path)}
}

func (_c *MockWorktree_Add_Call) Run(run func(path string)) *MockWorktree_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockWorktree_Add_Call) Return(_a0 plumbing.Hash, _a1 error) *MockWorktree_Add_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorktree_Add_Call) RunAndReturn(run func(string) (plumbing.Hash, error)) *MockWorktree_Add_Call {
	_c.Call.Return(run)
	return _c
}

// AddGlob provides a mock function with given fields: pattern
func (_m *MockWorktree) AddGlob(pattern string) error {
	ret := _m.Called(pattern)

	if len(ret) == 0 {
		panic("no return value specified for AddGlob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(pattern)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorktree_AddGlob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGlob'
type MockWorktree_AddGlob_Call struct {
	*mock.Call
}

// AddGlob is a helper method to define mock.On call
//   - pattern string
func (_e *MockWorktree_Expecter) AddGlob(pattern interface{}) *MockWorktree_AddGlob_Call {
	return &MockWorktree_AddGlob_Call{Call: _e.mock.On("AddGlob", pattern)}
}

func (_c *MockWorktree_AddGlob_Call) Run(run func(pattern string)) *MockWorktree_AddGlob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockWorktree_AddGlob_Call) Return(_a0 error) *MockWorktree_AddGlob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorktree_AddGlob_Call) RunAndReturn(run func(string) error) *MockWorktree_AddGlob_Call {
	_c.Call.Return(run)
	return _c
}

// Checkout provides a mock function with given fields: opts
func (_m *MockWorktree) Checkout(opts *git.CheckoutOptions) error {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for Checkout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*git.CheckoutOptions) error); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorktree_Checkout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Checkout'
type MockWorktree_Checkout_Call struct {
	*mock.Call
}

// Checkout is a helper method to define mock.On call
//   - opts *git.CheckoutOptions
func (_e *MockWorktree_Expecter) Checkout(opts interface{}) *MockWorktree_Checkout_Call {
	return &MockWorktree_Checkout_Call{Call: _e.mock.On("Checkout", opts)}
}

func (_c *MockWorktree_Checkout_Call) Run(run func(opts *git.CheckoutOptions)) *MockWorktree_Checkout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*git.CheckoutOptions))
	})
	return _c
}

func (_c *MockWorktree_Checkout_Call) Return(_a0 error) *MockWorktree_Checkout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorktree_Checkout_Call) RunAndReturn(run func(*git.CheckoutOptions) error) *MockWorktree_Checkout_Call {
	_c.Call.Return(run)
	return _c
}

// Commit provides a mock function with given fields: msg, opts
func (_m *MockWorktree) Commit(msg string, opts *git.CommitOptions) (plumbing.Hash, error) {
	ret := _m.Called(msg, opts)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 plumbing.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *git.CommitOptions) (plumbing.Hash, error)); ok {
		return rf(msg, opts)
	}
	if rf, ok := ret.Get(0).(func(string, *git.CommitOptions) plumbing.Hash); ok {
		r0 = rf(msg, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plumbing.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *git.CommitOptions) error); ok {
		r1 = rf(msg, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorktree_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type MockWorktree_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
//   - msg string
//   - opts *git.CommitOptions
func (_e *MockWorktree_Expecter) Commit(msg interface{}, opts interface{}) *MockWorktree_Commit_Call {
	return &MockWorktree_Commit_Call{Call: _e.mock.On("Commit", msg, opts)}
}

func (_c *MockWorktree_Commit_Call) Run(run func(msg string, opts *git.CommitOptions)) *MockWorktree_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*git.CommitOptions))
	})
	return _c
}

func (_c *MockWorktree_Commit_Call) Return(_a0 plumbing.Hash, _a1 error) *MockWorktree_Commit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorktree_Commit_Call) RunAndReturn(run func(string, *git.CommitOptions) (plumbing.Hash, error)) *MockWorktree_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: path
func (_m *MockWorktree) Remove(path string) (plumbing.Hash, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 plumbing.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (plumbing.Hash, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) plumbing.Hash); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plumbing.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorktree_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MockWorktree_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - path string
func (_e *MockWorktree_Expecter) Remove(path interface{}) *MockWorktree_Remove_Call {
	return &MockWorktree_Remove_Call{Call: _e.mock.On("Remove", path)}
}

func (_c *MockWorktree_Remove_Call) Run(run func(path string)) *MockWorktree_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockWorktree_Remove_Call) Return(_a0 plumbing.Hash, _a1 error) *MockWorktree_Remove_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorktree_Remove_Call) RunAndReturn(run func(string) (plumbing.Hash, error)) *MockWorktree_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with no fields
func (_m *MockWorktree) Status() (git.Status, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 git.Status
	var r1 error
	if rf, ok := ret.Get(0).(func() (git.Status, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() git.Status); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(git.Status)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorktree_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type MockWorktree_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
func (_e *MockWorktree_Expecter) Status() *MockWorktree_Status_Call {
	return &MockWorktree_Status_Call{Call: _e.mock.On("Status")}
}

func (_c *MockWorktree_Status_Call) Run(run func()) *MockWorktree_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorktree_Status_Call) Return(_a0 git.Status, _a1 error) *MockWorktree_Status_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorktree_Status_Call) RunAndReturn(run func() (git.Status, error)) *MockWorktree_Status_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorktree creates a new instance of MockWorktree. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorktree(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorktree {
	mock := &MockWorktree{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
