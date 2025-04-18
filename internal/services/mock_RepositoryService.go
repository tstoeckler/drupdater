// Code generated by mockery v2.53.3. DO NOT EDIT.

package services

import (
	internal "drupdater/internal"

	mock "github.com/stretchr/testify/mock"

	object "github.com/go-git/go-git/v5/plumbing/object"
)

// MockRepositoryService is an autogenerated mock type for the RepositoryService type
type MockRepositoryService struct {
	mock.Mock
}

type MockRepositoryService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepositoryService) EXPECT() *MockRepositoryService_Expecter {
	return &MockRepositoryService_Expecter{mock: &_m.Mock}
}

// BranchExists provides a mock function with given fields: repository, branch
func (_m *MockRepositoryService) BranchExists(repository internal.Repository, branch string) (bool, error) {
	ret := _m.Called(repository, branch)

	if len(ret) == 0 {
		panic("no return value specified for BranchExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(internal.Repository, string) (bool, error)); ok {
		return rf(repository, branch)
	}
	if rf, ok := ret.Get(0).(func(internal.Repository, string) bool); ok {
		r0 = rf(repository, branch)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(internal.Repository, string) error); ok {
		r1 = rf(repository, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryService_BranchExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BranchExists'
type MockRepositoryService_BranchExists_Call struct {
	*mock.Call
}

// BranchExists is a helper method to define mock.On call
//   - repository internal.Repository
//   - branch string
func (_e *MockRepositoryService_Expecter) BranchExists(repository interface{}, branch interface{}) *MockRepositoryService_BranchExists_Call {
	return &MockRepositoryService_BranchExists_Call{Call: _e.mock.On("BranchExists", repository, branch)}
}

func (_c *MockRepositoryService_BranchExists_Call) Run(run func(repository internal.Repository, branch string)) *MockRepositoryService_BranchExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(internal.Repository), args[1].(string))
	})
	return _c
}

func (_c *MockRepositoryService_BranchExists_Call) Return(_a0 bool, _a1 error) *MockRepositoryService_BranchExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryService_BranchExists_Call) RunAndReturn(run func(internal.Repository, string) (bool, error)) *MockRepositoryService_BranchExists_Call {
	_c.Call.Return(run)
	return _c
}

// CloneRepository provides a mock function with given fields: repository, branch, token
func (_m *MockRepositoryService) CloneRepository(repository string, branch string, token string) (internal.Repository, internal.Worktree, string, error) {
	ret := _m.Called(repository, branch, token)

	if len(ret) == 0 {
		panic("no return value specified for CloneRepository")
	}

	var r0 internal.Repository
	var r1 internal.Worktree
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(string, string, string) (internal.Repository, internal.Worktree, string, error)); ok {
		return rf(repository, branch, token)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) internal.Repository); ok {
		r0 = rf(repository, branch, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(internal.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) internal.Worktree); ok {
		r1 = rf(repository, branch, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(internal.Worktree)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string, string) string); ok {
		r2 = rf(repository, branch, token)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(string, string, string) error); ok {
		r3 = rf(repository, branch, token)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockRepositoryService_CloneRepository_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloneRepository'
type MockRepositoryService_CloneRepository_Call struct {
	*mock.Call
}

// CloneRepository is a helper method to define mock.On call
//   - repository string
//   - branch string
//   - token string
func (_e *MockRepositoryService_Expecter) CloneRepository(repository interface{}, branch interface{}, token interface{}) *MockRepositoryService_CloneRepository_Call {
	return &MockRepositoryService_CloneRepository_Call{Call: _e.mock.On("CloneRepository", repository, branch, token)}
}

func (_c *MockRepositoryService_CloneRepository_Call) Run(run func(repository string, branch string, token string)) *MockRepositoryService_CloneRepository_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockRepositoryService_CloneRepository_Call) Return(_a0 internal.Repository, _a1 internal.Worktree, _a2 string, _a3 error) *MockRepositoryService_CloneRepository_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

func (_c *MockRepositoryService_CloneRepository_Call) RunAndReturn(run func(string, string, string) (internal.Repository, internal.Worktree, string, error)) *MockRepositoryService_CloneRepository_Call {
	_c.Call.Return(run)
	return _c
}

// GetHeadCommit provides a mock function with given fields: repository
func (_m *MockRepositoryService) GetHeadCommit(repository internal.Repository) (*object.Commit, error) {
	ret := _m.Called(repository)

	if len(ret) == 0 {
		panic("no return value specified for GetHeadCommit")
	}

	var r0 *object.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(internal.Repository) (*object.Commit, error)); ok {
		return rf(repository)
	}
	if rf, ok := ret.Get(0).(func(internal.Repository) *object.Commit); ok {
		r0 = rf(repository)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*object.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(internal.Repository) error); ok {
		r1 = rf(repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryService_GetHeadCommit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeadCommit'
type MockRepositoryService_GetHeadCommit_Call struct {
	*mock.Call
}

// GetHeadCommit is a helper method to define mock.On call
//   - repository internal.Repository
func (_e *MockRepositoryService_Expecter) GetHeadCommit(repository interface{}) *MockRepositoryService_GetHeadCommit_Call {
	return &MockRepositoryService_GetHeadCommit_Call{Call: _e.mock.On("GetHeadCommit", repository)}
}

func (_c *MockRepositoryService_GetHeadCommit_Call) Run(run func(repository internal.Repository)) *MockRepositoryService_GetHeadCommit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(internal.Repository))
	})
	return _c
}

func (_c *MockRepositoryService_GetHeadCommit_Call) Return(_a0 *object.Commit, _a1 error) *MockRepositoryService_GetHeadCommit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryService_GetHeadCommit_Call) RunAndReturn(run func(internal.Repository) (*object.Commit, error)) *MockRepositoryService_GetHeadCommit_Call {
	_c.Call.Return(run)
	return _c
}

// IsSomethingStagedInPath provides a mock function with given fields: worktree, dir
func (_m *MockRepositoryService) IsSomethingStagedInPath(worktree internal.Worktree, dir string) bool {
	ret := _m.Called(worktree, dir)

	if len(ret) == 0 {
		panic("no return value specified for IsSomethingStagedInPath")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(internal.Worktree, string) bool); ok {
		r0 = rf(worktree, dir)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockRepositoryService_IsSomethingStagedInPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSomethingStagedInPath'
type MockRepositoryService_IsSomethingStagedInPath_Call struct {
	*mock.Call
}

// IsSomethingStagedInPath is a helper method to define mock.On call
//   - worktree internal.Worktree
//   - dir string
func (_e *MockRepositoryService_Expecter) IsSomethingStagedInPath(worktree interface{}, dir interface{}) *MockRepositoryService_IsSomethingStagedInPath_Call {
	return &MockRepositoryService_IsSomethingStagedInPath_Call{Call: _e.mock.On("IsSomethingStagedInPath", worktree, dir)}
}

func (_c *MockRepositoryService_IsSomethingStagedInPath_Call) Run(run func(worktree internal.Worktree, dir string)) *MockRepositoryService_IsSomethingStagedInPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(internal.Worktree), args[1].(string))
	})
	return _c
}

func (_c *MockRepositoryService_IsSomethingStagedInPath_Call) Return(_a0 bool) *MockRepositoryService_IsSomethingStagedInPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepositoryService_IsSomethingStagedInPath_Call) RunAndReturn(run func(internal.Worktree, string) bool) *MockRepositoryService_IsSomethingStagedInPath_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepositoryService creates a new instance of MockRepositoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepositoryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepositoryService {
	mock := &MockRepositoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
