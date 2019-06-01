// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import arguments "github.com/tanphamhaiduong/go/delta/server/arguments"
import context "context"

import mock "github.com/stretchr/testify/mock"
import models "github.com/tanphamhaiduong/go/delta/server/models"

// IHandler is an autogenerated mock type for the IHandler type
type IHandler struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, params
func (_m *IHandler) Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.FeatureCountArgs) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.FeatureCountArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, params
func (_m *IHandler) Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, arguments.FeatureDeleteArgs) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.FeatureDeleteArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, params
func (_m *IHandler) GetByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error) {
	ret := _m.Called(ctx, params)

	var r0 models.Feature
	if rf, ok := ret.Get(0).(func(context.Context, arguments.FeatureGetByIDArgs) models.Feature); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.Feature)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.FeatureGetByIDArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDs provides a mock function with given fields: ctx, params
func (_m *IHandler) GetByIDs(ctx context.Context, params arguments.FeatureGetByIDsArgs) ([]models.Feature, error) {
	ret := _m.Called(ctx, params)

	var r0 []models.Feature
	if rf, ok := ret.Get(0).(func(context.Context, arguments.FeatureGetByIDsArgs) []models.Feature); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Feature)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.FeatureGetByIDsArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, params
func (_m *IHandler) Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error) {
	ret := _m.Called(ctx, params)

	var r0 models.Feature
	if rf, ok := ret.Get(0).(func(context.Context, arguments.FeatureInsertArgs) models.Feature); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.Feature)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.FeatureInsertArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *IHandler) List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error) {
	ret := _m.Called(ctx, params)

	var r0 []models.Feature
	if rf, ok := ret.Get(0).(func(context.Context, arguments.FeatureListArgs) []models.Feature); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Feature)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.FeatureListArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, params
func (_m *IHandler) Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error) {
	ret := _m.Called(ctx, params)

	var r0 models.Feature
	if rf, ok := ret.Get(0).(func(context.Context, arguments.FeatureUpdateArgs) models.Feature); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(models.Feature)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, arguments.FeatureUpdateArgs) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
