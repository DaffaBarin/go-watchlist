// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	admins "go-watchlist/business/admins"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Login provides a mock function with given fields: username, password
func (_m *Repository) Login(username string, password string) (admins.Domain, error) {
	ret := _m.Called(username, password)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(string, string) admins.Domain); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: domain
func (_m *Repository) Register(domain *admins.Domain) (admins.Domain, error) {
	ret := _m.Called(domain)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(*admins.Domain) admins.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admins.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
