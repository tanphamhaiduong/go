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
func (_m *IHandler) Count(ctx context.Context, params arguments.PermissionCount) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.PermissionCount) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.PermissionCount) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, param
func (_m *IHandler) Delete(ctx context.Context, param arguments.PermissionDelete) (int64, error) {
	ret := _m.Called(ctx, param)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.PermissionDelete) int64); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.PermissionDelete) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, param
func (_m *IHandler) GetByID(ctx context.Context, param arguments.PermissionGetByID) (models.Permission, error) {
	ret := _m.Called(ctx, param)

	var r0 models.Permission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.PermissionGetByID) models.Permission); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(models.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.PermissionGetByID) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDs provides a mock function with given fields: ctx, param
func (_m *IHandler) GetByIDs(ctx context.Context, param arguments.PermissionGetByIDs) ([]models.Permission, error) {
	ret := _m.Called(ctx, param)

	var r0 []models.Permission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.PermissionGetByIDs) []models.Permission); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Permission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.PermissionGetByIDs) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, params
func (_m *IHandler) Insert(ctx context.Context, params arguments.PermissionInsert) (models.Permission, error) {
	ret := _m.Called(ctx, params)

	var r0 models.Permission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.PermissionInsert) models.Permission); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.PermissionInsert) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *IHandler) List(ctx context.Context, params arguments.PermissionList) ([]models.Permission, error) {
	ret := _m.Called(ctx, params)

	var r0 []models.Permission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.PermissionList) []models.Permission); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Permission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.PermissionList) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *IHandler) Update(ctx context.Context, params arguments.PermissionUpdate) (models.Permission, error) {
	ret := _m.Called(ctx, params)

	var r0 models.Permission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.PermissionUpdate) models.Permission); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.PermissionUpdate) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
