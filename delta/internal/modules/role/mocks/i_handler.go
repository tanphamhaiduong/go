// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	arguments "github.com/tanphamhaiduong/go/delta/internal/arguments"

	mock "github.com/stretchr/testify/mock"

	models "github.com/tanphamhaiduong/go/delta/internal/models"
)

// IHandler is an autogenerated mock type for the IHandler type
type IHandler struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, params
func (_m *IHandler) Count(ctx context.Context, params arguments.RoleCount) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RoleCount) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RoleCount) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, param
func (_m *IHandler) Delete(ctx context.Context, param arguments.RoleDelete) (int64, error) {
	ret := _m.Called(ctx, param)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RoleDelete) int64); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RoleDelete) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, param
func (_m *IHandler) GetByID(ctx context.Context, param arguments.RoleGetByID) (models.Role, error) {
	ret := _m.Called(ctx, param)

	var r0 models.Role
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RoleGetByID) models.Role); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(models.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RoleGetByID) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDs provides a mock function with given fields: ctx, param
func (_m *IHandler) GetByIDs(ctx context.Context, param arguments.RoleGetByIDs) ([]models.Role, error) {
	ret := _m.Called(ctx, param)

	var r0 []models.Role
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RoleGetByIDs) []models.Role); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RoleGetByIDs) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, params
func (_m *IHandler) Insert(ctx context.Context, params arguments.RoleInsert) (models.Role, error) {
	ret := _m.Called(ctx, params)

	var r0 models.Role
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RoleInsert) models.Role); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RoleInsert) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *IHandler) List(ctx context.Context, params arguments.RoleList) ([]models.Role, error) {
	ret := _m.Called(ctx, params)

	var r0 []models.Role
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RoleList) []models.Role); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RoleList) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *IHandler) Update(ctx context.Context, params arguments.RoleUpdate) (models.Role, error) {
	ret := _m.Called(ctx, params)

	var r0 models.Role
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RoleUpdate) models.Role); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RoleUpdate) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
