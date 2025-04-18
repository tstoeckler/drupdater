// Code generated by mockery v2.53.3. DO NOT EDIT.

package codehosting

import mock "github.com/stretchr/testify/mock"

// MockPlatform is an autogenerated mock type for the Platform type
type MockPlatform struct {
	mock.Mock
}

type MockPlatform_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlatform) EXPECT() *MockPlatform_Expecter {
	return &MockPlatform_Expecter{mock: &_m.Mock}
}

// CreateMergeRequest provides a mock function with given fields: title, description, sourceBranch, targetBranch
func (_m *MockPlatform) CreateMergeRequest(title string, description string, sourceBranch string, targetBranch string) (MergeRequest, error) {
	ret := _m.Called(title, description, sourceBranch, targetBranch)

	if len(ret) == 0 {
		panic("no return value specified for CreateMergeRequest")
	}

	var r0 MergeRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) (MergeRequest, error)); ok {
		return rf(title, description, sourceBranch, targetBranch)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string) MergeRequest); ok {
		r0 = rf(title, description, sourceBranch, targetBranch)
	} else {
		r0 = ret.Get(0).(MergeRequest)
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(title, description, sourceBranch, targetBranch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlatform_CreateMergeRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMergeRequest'
type MockPlatform_CreateMergeRequest_Call struct {
	*mock.Call
}

// CreateMergeRequest is a helper method to define mock.On call
//   - title string
//   - description string
//   - sourceBranch string
//   - targetBranch string
func (_e *MockPlatform_Expecter) CreateMergeRequest(title interface{}, description interface{}, sourceBranch interface{}, targetBranch interface{}) *MockPlatform_CreateMergeRequest_Call {
	return &MockPlatform_CreateMergeRequest_Call{Call: _e.mock.On("CreateMergeRequest", title, description, sourceBranch, targetBranch)}
}

func (_c *MockPlatform_CreateMergeRequest_Call) Run(run func(title string, description string, sourceBranch string, targetBranch string)) *MockPlatform_CreateMergeRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockPlatform_CreateMergeRequest_Call) Return(_a0 MergeRequest, _a1 error) *MockPlatform_CreateMergeRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlatform_CreateMergeRequest_Call) RunAndReturn(run func(string, string, string, string) (MergeRequest, error)) *MockPlatform_CreateMergeRequest_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadComposerFiles provides a mock function with given fields: branch
func (_m *MockPlatform) DownloadComposerFiles(branch string) string {
	ret := _m.Called(branch)

	if len(ret) == 0 {
		panic("no return value specified for DownloadComposerFiles")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(branch)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlatform_DownloadComposerFiles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadComposerFiles'
type MockPlatform_DownloadComposerFiles_Call struct {
	*mock.Call
}

// DownloadComposerFiles is a helper method to define mock.On call
//   - branch string
func (_e *MockPlatform_Expecter) DownloadComposerFiles(branch interface{}) *MockPlatform_DownloadComposerFiles_Call {
	return &MockPlatform_DownloadComposerFiles_Call{Call: _e.mock.On("DownloadComposerFiles", branch)}
}

func (_c *MockPlatform_DownloadComposerFiles_Call) Run(run func(branch string)) *MockPlatform_DownloadComposerFiles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlatform_DownloadComposerFiles_Call) Return(_a0 string) *MockPlatform_DownloadComposerFiles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlatform_DownloadComposerFiles_Call) RunAndReturn(run func(string) string) *MockPlatform_DownloadComposerFiles_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlatform creates a new instance of MockPlatform. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlatform(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlatform {
	mock := &MockPlatform{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
