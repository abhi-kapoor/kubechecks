// Code generated by mockery v2.37.1. DO NOT EDIT.

package github_client

import (
	context "context"

	github "github.com/google/go-github/v62/github"

	mock "github.com/stretchr/testify/mock"
)

// MockRepositoriesServices is an autogenerated mock type for the RepositoriesServices type
type MockRepositoriesServices struct {
	mock.Mock
}

type MockRepositoriesServices_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepositoriesServices) EXPECT() *MockRepositoriesServices_Expecter {
	return &MockRepositoriesServices_Expecter{mock: &_m.Mock}
}

// CreateHook provides a mock function with given fields: ctx, owner, repo, hook
func (_m *MockRepositoriesServices) CreateHook(ctx context.Context, owner string, repo string, hook *github.Hook) (*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, hook)

	var r0 *github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.Hook) (*github.Hook, *github.Response, error)); ok {
		return rf(ctx, owner, repo, hook)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.Hook) *github.Hook); ok {
		r0 = rf(ctx, owner, repo, hook)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.Hook) *github.Response); ok {
		r1 = rf(ctx, owner, repo, hook)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.Hook) error); ok {
		r2 = rf(ctx, owner, repo, hook)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepositoriesServices_CreateHook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateHook'
type MockRepositoriesServices_CreateHook_Call struct {
	*mock.Call
}

// CreateHook is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - hook *github.Hook
func (_e *MockRepositoriesServices_Expecter) CreateHook(ctx interface{}, owner interface{}, repo interface{}, hook interface{}) *MockRepositoriesServices_CreateHook_Call {
	return &MockRepositoriesServices_CreateHook_Call{Call: _e.mock.On("CreateHook", ctx, owner, repo, hook)}
}

func (_c *MockRepositoriesServices_CreateHook_Call) Run(run func(ctx context.Context, owner string, repo string, hook *github.Hook)) *MockRepositoriesServices_CreateHook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*github.Hook))
	})
	return _c
}

func (_c *MockRepositoriesServices_CreateHook_Call) Return(_a0 *github.Hook, _a1 *github.Response, _a2 error) *MockRepositoriesServices_CreateHook_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepositoriesServices_CreateHook_Call) RunAndReturn(run func(context.Context, string, string, *github.Hook) (*github.Hook, *github.Response, error)) *MockRepositoriesServices_CreateHook_Call {
	_c.Call.Return(run)
	return _c
}

// CreateStatus provides a mock function with given fields: ctx, owner, repo, ref, status
func (_m *MockRepositoriesServices) CreateStatus(ctx context.Context, owner string, repo string, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, ref, status)

	var r0 *github.RepoStatus
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.RepoStatus) (*github.RepoStatus, *github.Response, error)); ok {
		return rf(ctx, owner, repo, ref, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.RepoStatus) *github.RepoStatus); ok {
		r0 = rf(ctx, owner, repo, ref, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepoStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *github.RepoStatus) *github.Response); ok {
		r1 = rf(ctx, owner, repo, ref, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, *github.RepoStatus) error); ok {
		r2 = rf(ctx, owner, repo, ref, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepositoriesServices_CreateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateStatus'
type MockRepositoriesServices_CreateStatus_Call struct {
	*mock.Call
}

// CreateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - ref string
//   - status *github.RepoStatus
func (_e *MockRepositoriesServices_Expecter) CreateStatus(ctx interface{}, owner interface{}, repo interface{}, ref interface{}, status interface{}) *MockRepositoriesServices_CreateStatus_Call {
	return &MockRepositoriesServices_CreateStatus_Call{Call: _e.mock.On("CreateStatus", ctx, owner, repo, ref, status)}
}

func (_c *MockRepositoriesServices_CreateStatus_Call) Run(run func(ctx context.Context, owner string, repo string, ref string, status *github.RepoStatus)) *MockRepositoriesServices_CreateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(*github.RepoStatus))
	})
	return _c
}

func (_c *MockRepositoriesServices_CreateStatus_Call) Return(_a0 *github.RepoStatus, _a1 *github.Response, _a2 error) *MockRepositoriesServices_CreateStatus_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepositoriesServices_CreateStatus_Call) RunAndReturn(run func(context.Context, string, string, string, *github.RepoStatus) (*github.RepoStatus, *github.Response, error)) *MockRepositoriesServices_CreateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, owner, repo
func (_m *MockRepositoriesServices) Get(ctx context.Context, owner string, repo string) (*github.Repository, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 *github.Repository
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*github.Repository, *github.Response, error)); ok {
		return rf(ctx, owner, repo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *github.Repository); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, owner, repo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepositoriesServices_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRepositoriesServices_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
func (_e *MockRepositoriesServices_Expecter) Get(ctx interface{}, owner interface{}, repo interface{}) *MockRepositoriesServices_Get_Call {
	return &MockRepositoriesServices_Get_Call{Call: _e.mock.On("Get", ctx, owner, repo)}
}

func (_c *MockRepositoriesServices_Get_Call) Run(run func(ctx context.Context, owner string, repo string)) *MockRepositoriesServices_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockRepositoriesServices_Get_Call) Return(_a0 *github.Repository, _a1 *github.Response, _a2 error) *MockRepositoriesServices_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepositoriesServices_Get_Call) RunAndReturn(run func(context.Context, string, string) (*github.Repository, *github.Response, error)) *MockRepositoriesServices_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetContents provides a mock function with given fields: ctx, owner, repo, path, opts
func (_m *MockRepositoriesServices) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, path, opts)

	var r0 *github.RepositoryContent
	var r1 []*github.RepositoryContent
	var r2 *github.Response
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.RepositoryContentGetOptions) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error)); ok {
		return rf(ctx, owner, repo, path, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.RepositoryContentGetOptions) *github.RepositoryContent); ok {
		r0 = rf(ctx, owner, repo, path, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.RepositoryContent)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *github.RepositoryContentGetOptions) []*github.RepositoryContent); ok {
		r1 = rf(ctx, owner, repo, path, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*github.RepositoryContent)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, *github.RepositoryContentGetOptions) *github.Response); ok {
		r2 = rf(ctx, owner, repo, path, opts)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*github.Response)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, string, string, string, *github.RepositoryContentGetOptions) error); ok {
		r3 = rf(ctx, owner, repo, path, opts)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockRepositoriesServices_GetContents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContents'
type MockRepositoriesServices_GetContents_Call struct {
	*mock.Call
}

// GetContents is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - path string
//   - opts *github.RepositoryContentGetOptions
func (_e *MockRepositoriesServices_Expecter) GetContents(ctx interface{}, owner interface{}, repo interface{}, path interface{}, opts interface{}) *MockRepositoriesServices_GetContents_Call {
	return &MockRepositoriesServices_GetContents_Call{Call: _e.mock.On("GetContents", ctx, owner, repo, path, opts)}
}

func (_c *MockRepositoriesServices_GetContents_Call) Run(run func(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions)) *MockRepositoriesServices_GetContents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(*github.RepositoryContentGetOptions))
	})
	return _c
}

func (_c *MockRepositoriesServices_GetContents_Call) Return(fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error) *MockRepositoriesServices_GetContents_Call {
	_c.Call.Return(fileContent, directoryContent, resp, err)
	return _c
}

func (_c *MockRepositoriesServices_GetContents_Call) RunAndReturn(run func(context.Context, string, string, string, *github.RepositoryContentGetOptions) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error)) *MockRepositoriesServices_GetContents_Call {
	_c.Call.Return(run)
	return _c
}

// ListHooks provides a mock function with given fields: ctx, owner, repo, opts
func (_m *MockRepositoriesServices) ListHooks(ctx context.Context, owner string, repo string, opts *github.ListOptions) ([]*github.Hook, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo, opts)

	var r0 []*github.Hook
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.ListOptions) ([]*github.Hook, *github.Response, error)); ok {
		return rf(ctx, owner, repo, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.ListOptions) []*github.Hook); ok {
		r0 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.Hook)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.ListOptions) *github.Response); ok {
		r1 = rf(ctx, owner, repo, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.ListOptions) error); ok {
		r2 = rf(ctx, owner, repo, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepositoriesServices_ListHooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListHooks'
type MockRepositoriesServices_ListHooks_Call struct {
	*mock.Call
}

// ListHooks is a helper method to define mock.On call
//   - ctx context.Context
//   - owner string
//   - repo string
//   - opts *github.ListOptions
func (_e *MockRepositoriesServices_Expecter) ListHooks(ctx interface{}, owner interface{}, repo interface{}, opts interface{}) *MockRepositoriesServices_ListHooks_Call {
	return &MockRepositoriesServices_ListHooks_Call{Call: _e.mock.On("ListHooks", ctx, owner, repo, opts)}
}

func (_c *MockRepositoriesServices_ListHooks_Call) Run(run func(ctx context.Context, owner string, repo string, opts *github.ListOptions)) *MockRepositoriesServices_ListHooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*github.ListOptions))
	})
	return _c
}

func (_c *MockRepositoriesServices_ListHooks_Call) Return(_a0 []*github.Hook, _a1 *github.Response, _a2 error) *MockRepositoriesServices_ListHooks_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepositoriesServices_ListHooks_Call) RunAndReturn(run func(context.Context, string, string, *github.ListOptions) ([]*github.Hook, *github.Response, error)) *MockRepositoriesServices_ListHooks_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepositoriesServices creates a new instance of MockRepositoriesServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepositoriesServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepositoriesServices {
	mock := &MockRepositoriesServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
