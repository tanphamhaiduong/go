// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database "github.com/tanphamhaiduong/go/delta/internal/database"
)

// IDB is an autogenerated mock type for the IDB type
type IDB struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *IDB) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *IDB) Ping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrepareContext provides a mock function with given fields: ctx, query
func (_m *IDB) PrepareContext(ctx context.Context, query string) (database.IStmt, error) {
	ret := _m.Called(ctx, query)

	var r0 database.IStmt
	if rf, ok := ret.Get(0).(func(context.Context, string) database.IStmt); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.IStmt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
