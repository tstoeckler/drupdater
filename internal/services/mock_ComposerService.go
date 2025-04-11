// Code generated by mockery v2.51.1. DO NOT EDIT.

package services

import mock "github.com/stretchr/testify/mock"

// MockComposerService is an autogenerated mock type for the ComposerService type
type MockComposerService struct {
	mock.Mock
}

type MockComposerService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockComposerService) EXPECT() *MockComposerService_Expecter {
	return &MockComposerService_Expecter{mock: &_m.Mock}
}

// CheckPatchApplies provides a mock function with given fields: packageName, packageVersion, patchPath
func (_m *MockComposerService) CheckPatchApplies(packageName string, packageVersion string, patchPath string) (bool, error) {
	ret := _m.Called(packageName, packageVersion, patchPath)

	if len(ret) == 0 {
		panic("no return value specified for CheckPatchApplies")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return rf(packageName, packageVersion, patchPath)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(packageName, packageVersion, patchPath)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(packageName, packageVersion, patchPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockComposerService_CheckPatchApplies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckPatchApplies'
type MockComposerService_CheckPatchApplies_Call struct {
	*mock.Call
}

// CheckPatchApplies is a helper method to define mock.On call
//   - packageName string
//   - packageVersion string
//   - patchPath string
func (_e *MockComposerService_Expecter) CheckPatchApplies(packageName interface{}, packageVersion interface{}, patchPath interface{}) *MockComposerService_CheckPatchApplies_Call {
	return &MockComposerService_CheckPatchApplies_Call{Call: _e.mock.On("CheckPatchApplies", packageName, packageVersion, patchPath)}
}

func (_c *MockComposerService_CheckPatchApplies_Call) Run(run func(packageName string, packageVersion string, patchPath string)) *MockComposerService_CheckPatchApplies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockComposerService_CheckPatchApplies_Call) Return(_a0 bool, _a1 error) *MockComposerService_CheckPatchApplies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockComposerService_CheckPatchApplies_Call) RunAndReturn(run func(string, string, string) (bool, error)) *MockComposerService_CheckPatchApplies_Call {
	_c.Call.Return(run)
	return _c
}

// GetComposerLockHash provides a mock function with given fields: dir
func (_m *MockComposerService) GetComposerLockHash(dir string) (string, error) {
	ret := _m.Called(dir)

	if len(ret) == 0 {
		panic("no return value specified for GetComposerLockHash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(dir)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(dir)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockComposerService_GetComposerLockHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComposerLockHash'
type MockComposerService_GetComposerLockHash_Call struct {
	*mock.Call
}

// GetComposerLockHash is a helper method to define mock.On call
//   - dir string
func (_e *MockComposerService_Expecter) GetComposerLockHash(dir interface{}) *MockComposerService_GetComposerLockHash_Call {
	return &MockComposerService_GetComposerLockHash_Call{Call: _e.mock.On("GetComposerLockHash", dir)}
}

func (_c *MockComposerService_GetComposerLockHash_Call) Run(run func(dir string)) *MockComposerService_GetComposerLockHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockComposerService_GetComposerLockHash_Call) Return(_a0 string, _a1 error) *MockComposerService_GetComposerLockHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockComposerService_GetComposerLockHash_Call) RunAndReturn(run func(string) (string, error)) *MockComposerService_GetComposerLockHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetComposerUpdates provides a mock function with given fields: dir, packagesToUpdate, minimalChanges
func (_m *MockComposerService) GetComposerUpdates(dir string, packagesToUpdate []string, minimalChanges bool) ([]PackageChange, error) {
	ret := _m.Called(dir, packagesToUpdate, minimalChanges)

	if len(ret) == 0 {
		panic("no return value specified for GetComposerUpdates")
	}

	var r0 []PackageChange
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string, bool) ([]PackageChange, error)); ok {
		return rf(dir, packagesToUpdate, minimalChanges)
	}
	if rf, ok := ret.Get(0).(func(string, []string, bool) []PackageChange); ok {
		r0 = rf(dir, packagesToUpdate, minimalChanges)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]PackageChange)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, bool) error); ok {
		r1 = rf(dir, packagesToUpdate, minimalChanges)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockComposerService_GetComposerUpdates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComposerUpdates'
type MockComposerService_GetComposerUpdates_Call struct {
	*mock.Call
}

// GetComposerUpdates is a helper method to define mock.On call
//   - dir string
//   - packagesToUpdate []string
//   - minimalChanges bool
func (_e *MockComposerService_Expecter) GetComposerUpdates(dir interface{}, packagesToUpdate interface{}, minimalChanges interface{}) *MockComposerService_GetComposerUpdates_Call {
	return &MockComposerService_GetComposerUpdates_Call{Call: _e.mock.On("GetComposerUpdates", dir, packagesToUpdate, minimalChanges)}
}

func (_c *MockComposerService_GetComposerUpdates_Call) Run(run func(dir string, packagesToUpdate []string, minimalChanges bool)) *MockComposerService_GetComposerUpdates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string), args[2].(bool))
	})
	return _c
}

func (_c *MockComposerService_GetComposerUpdates_Call) Return(_a0 []PackageChange, _a1 error) *MockComposerService_GetComposerUpdates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockComposerService_GetComposerUpdates_Call) RunAndReturn(run func(string, []string, bool) ([]PackageChange, error)) *MockComposerService_GetComposerUpdates_Call {
	_c.Call.Return(run)
	return _c
}

// GetInstalledPlugins provides a mock function with given fields: dir
func (_m *MockComposerService) GetInstalledPlugins(dir string) (map[string]interface{}, error) {
	ret := _m.Called(dir)

	if len(ret) == 0 {
		panic("no return value specified for GetInstalledPlugins")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (map[string]interface{}, error)); ok {
		return rf(dir)
	}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(dir)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockComposerService_GetInstalledPlugins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInstalledPlugins'
type MockComposerService_GetInstalledPlugins_Call struct {
	*mock.Call
}

// GetInstalledPlugins is a helper method to define mock.On call
//   - dir string
func (_e *MockComposerService_Expecter) GetInstalledPlugins(dir interface{}) *MockComposerService_GetInstalledPlugins_Call {
	return &MockComposerService_GetInstalledPlugins_Call{Call: _e.mock.On("GetInstalledPlugins", dir)}
}

func (_c *MockComposerService_GetInstalledPlugins_Call) Run(run func(dir string)) *MockComposerService_GetInstalledPlugins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockComposerService_GetInstalledPlugins_Call) Return(_a0 map[string]interface{}, _a1 error) *MockComposerService_GetInstalledPlugins_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockComposerService_GetInstalledPlugins_Call) RunAndReturn(run func(string) (map[string]interface{}, error)) *MockComposerService_GetInstalledPlugins_Call {
	_c.Call.Return(run)
	return _c
}

// RunComposerAudit provides a mock function with given fields: dir
func (_m *MockComposerService) RunComposerAudit(dir string) (ComposerAudit, error) {
	ret := _m.Called(dir)

	if len(ret) == 0 {
		panic("no return value specified for RunComposerAudit")
	}

	var r0 ComposerAudit
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ComposerAudit, error)); ok {
		return rf(dir)
	}
	if rf, ok := ret.Get(0).(func(string) ComposerAudit); ok {
		r0 = rf(dir)
	} else {
		r0 = ret.Get(0).(ComposerAudit)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockComposerService_RunComposerAudit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunComposerAudit'
type MockComposerService_RunComposerAudit_Call struct {
	*mock.Call
}

// RunComposerAudit is a helper method to define mock.On call
//   - dir string
func (_e *MockComposerService_Expecter) RunComposerAudit(dir interface{}) *MockComposerService_RunComposerAudit_Call {
	return &MockComposerService_RunComposerAudit_Call{Call: _e.mock.On("RunComposerAudit", dir)}
}

func (_c *MockComposerService_RunComposerAudit_Call) Run(run func(dir string)) *MockComposerService_RunComposerAudit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockComposerService_RunComposerAudit_Call) Return(_a0 ComposerAudit, _a1 error) *MockComposerService_RunComposerAudit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockComposerService_RunComposerAudit_Call) RunAndReturn(run func(string) (ComposerAudit, error)) *MockComposerService_RunComposerAudit_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockComposerService creates a new instance of MockComposerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockComposerService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockComposerService {
	mock := &MockComposerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
