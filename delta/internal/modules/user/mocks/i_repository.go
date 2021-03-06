// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	arguments "github.com/tanphamhaiduong/go/delta/internal/arguments"

	mock "github.com/stretchr/testify/mock"

	models "github.com/tanphamhaiduong/go/delta/internal/models"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, params
func (_m *IRepository) Count(ctx context.Context, params arguments.UserCount) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.UserCount) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.UserCount) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, param
func (_m *IRepository) Delete(ctx context.Context, param arguments.UserDelete) (int64, error) {
	ret := _m.Called(ctx, param)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.UserDelete) int64); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.UserDelete) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, param
func (_m *IRepository) GetByID(ctx context.Context, param arguments.UserGetByID) (models.User, error) {
	ret := _m.Called(ctx, param)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, arguments.UserGetByID) models.User); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.UserGetByID) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDs provides a mock function with given fields: ctx, param
func (_m *IRepository) GetByIDs(ctx context.Context, param arguments.UserGetByIDs) ([]models.User, error) {
	ret := _m.Called(ctx, param)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, arguments.UserGetByIDs) []models.User); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.UserGetByIDs) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *IRepository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	ret := _m.Called(ctx, username)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, params
func (_m *IRepository) Insert(ctx context.Context, params arguments.UserInsert) (models.User, error) {
	ret := _m.Called(ctx, params)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, arguments.UserInsert) models.User); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.UserInsert) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *IRepository) List(ctx context.Context, params arguments.UserList) ([]models.User, error) {
	ret := _m.Called(ctx, params)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, arguments.UserList) []models.User); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.UserList) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *IRepository) Update(ctx context.Context, params arguments.UserUpdate) (models.User, error) {
	ret := _m.Called(ctx, params)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, arguments.UserUpdate) models.User); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.UserUpdate) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
