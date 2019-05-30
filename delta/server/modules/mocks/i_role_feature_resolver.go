// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import graphql "github.com/graphql-go/graphql"
import mock "github.com/stretchr/testify/mock"

// IRoleFeatureResolver is an autogenerated mock type for the IRoleFeatureResolver type
type IRoleFeatureResolver struct {
	mock.Mock
}

// Count provides a mock function with given fields: p
func (_m *IRoleFeatureResolver) Count(p graphql.ResolveParams) (interface{}, error) {
	ret := _m.Called(p)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(graphql.ResolveParams) interface{}); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.ResolveParams) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: p
func (_m *IRoleFeatureResolver) Delete(p graphql.ResolveParams) (interface{}, error) {
	ret := _m.Called(p)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(graphql.ResolveParams) interface{}); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.ResolveParams) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForwardParams provides a mock function with given fields: p
func (_m *IRoleFeatureResolver) ForwardParams(p graphql.ResolveParams) (interface{}, error) {
	ret := _m.Called(p)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(graphql.ResolveParams) interface{}); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.ResolveParams) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: p
func (_m *IRoleFeatureResolver) GetByID(p graphql.ResolveParams) (interface{}, error) {
	ret := _m.Called(p)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(graphql.ResolveParams) interface{}); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.ResolveParams) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: p
func (_m *IRoleFeatureResolver) Insert(p graphql.ResolveParams) (interface{}, error) {
	ret := _m.Called(p)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(graphql.ResolveParams) interface{}); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.ResolveParams) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: p
func (_m *IRoleFeatureResolver) List(p graphql.ResolveParams) (interface{}, error) {
	ret := _m.Called(p)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(graphql.ResolveParams) interface{}); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.ResolveParams) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: p
func (_m *IRoleFeatureResolver) Update(p graphql.ResolveParams) (interface{}, error) {
	ret := _m.Called(p)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(graphql.ResolveParams) interface{}); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.ResolveParams) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}