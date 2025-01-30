// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "Tasks/internal/model"
)

// CacheRepository is an autogenerated mock type for the CacheRepository type
type CacheRepository struct {
	mock.Mock
}

// DeleteTaskFromCache provides a mock function with given fields: ctx, taskID
func (_m *CacheRepository) DeleteTaskFromCache(ctx context.Context, taskID int) error {
	ret := _m.Called(ctx, taskID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTaskFromCache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, taskID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTaskFromCache provides a mock function with given fields: ctx, taskID
func (_m *CacheRepository) GetTaskFromCache(ctx context.Context, taskID int) (model.Task, error) {
	ret := _m.Called(ctx, taskID)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskFromCache")
	}

	var r0 model.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (model.Task, error)); ok {
		return rf(ctx, taskID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) model.Task); ok {
		r0 = rf(ctx, taskID)
	} else {
		r0 = ret.Get(0).(model.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertingCache provides a mock function with given fields: ctx, task
func (_m *CacheRepository) InsertingCache(ctx context.Context, task model.Task) error {
	ret := _m.Called(ctx, task)

	if len(ret) == 0 {
		panic("no return value specified for InsertingCache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Task) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTaskStatusInCache provides a mock function with given fields: ctx, taskID, status
func (_m *CacheRepository) UpdateTaskStatusInCache(ctx context.Context, taskID int, status string) error {
	ret := _m.Called(ctx, taskID, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTaskStatusInCache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, taskID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCacheRepository creates a new instance of CacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCacheRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CacheRepository {
	mock := &CacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
