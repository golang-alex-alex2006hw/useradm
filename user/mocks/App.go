// Copyright 2017 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import context "context"
import jwt "github.com/mendersoftware/useradm/jwt"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/useradm/model"
import useradm "github.com/mendersoftware/useradm/user"

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *App) CreateUser(ctx context.Context, u *model.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUserInitial provides a mock function with given fields: ctx, u
func (_m *App) CreateUserInitial(ctx context.Context, u *model.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: ctx, email, pass
func (_m *App) Login(ctx context.Context, email string, pass string) (*jwt.Token, error) {
	ret := _m.Called(ctx, email, pass)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *jwt.Token); ok {
		r0 = rf(ctx, email, pass)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, pass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignToken provides a mock function with given fields: ctx
func (_m *App) SignToken(ctx context.Context) jwt.SignFunc {
	ret := _m.Called(ctx)

	var r0 jwt.SignFunc
	if rf, ok := ret.Get(0).(func(context.Context) jwt.SignFunc); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.SignFunc)
		}
	}

	return r0
}

// Verify provides a mock function with given fields: ctx, token
func (_m *App) Verify(ctx context.Context, token *jwt.Token) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ useradm.App = (*App)(nil)
