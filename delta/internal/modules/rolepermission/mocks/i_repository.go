// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import arguments "github.com/tanphamhaiduong/go/delta/internal/arguments"
import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/tanphamhaiduong/go/delta/internal/models"

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, params
func (_m *IRepository) Count(ctx context.Context, params arguments.RolePermissionCountArgs) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RolePermissionCountArgs) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RolePermissionCountArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, params
func (_m *IRepository) Delete(ctx context.Context, params arguments.RolePermissionDeleteArgs) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RolePermissionDeleteArgs) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RolePermissionDeleteArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, params
func (_m *IRepository) GetByID(ctx context.Context, params arguments.RolePermissionGetByIDArgs) (models.RolePermission, error) {
	ret := _m.Called(ctx, params)

	var r0 models.RolePermission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RolePermissionGetByIDArgs) models.RolePermission); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.RolePermission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RolePermissionGetByIDArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDs provides a mock function with given fields: ctx, params
func (_m *IRepository) GetByIDs(ctx context.Context, params arguments.RolePermissionGetByIDsArgs) ([]models.RolePermission, error) {
	ret := _m.Called(ctx, params)

	var r0 []models.RolePermission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RolePermissionGetByIDsArgs) []models.RolePermission); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.RolePermission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RolePermissionGetByIDsArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, params
func (_m *IRepository) Insert(ctx context.Context, params arguments.RolePermissionInsertArgs) (models.RolePermission, error) {
	ret := _m.Called(ctx, params)

	var r0 models.RolePermission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RolePermissionInsertArgs) models.RolePermission); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.RolePermission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RolePermissionInsertArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *IRepository) List(ctx context.Context, params arguments.RolePermissionListArgs) ([]models.RolePermission, error) {
	ret := _m.Called(ctx, params)

	var r0 []models.RolePermission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RolePermissionListArgs) []models.RolePermission); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.RolePermission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RolePermissionListArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *IRepository) Update(ctx context.Context, params arguments.RolePermissionUpdateArgs) (models.RolePermission, error) {
	ret := _m.Called(ctx, params)

	var r0 models.RolePermission
	if rf, ok := ret.Get(0).(func(context.Context, arguments.RolePermissionUpdateArgs) models.RolePermission); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.RolePermission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.RolePermissionUpdateArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
