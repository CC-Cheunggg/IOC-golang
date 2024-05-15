// Code generated by mockery v2.12.2. DO NOT EDIT.

/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	autowire "github.com/cc-cheunggg/ioc-golang/autowire"
)

// FacadeAutowire is an autogenerated mock type for the FacadeAutowire type
type FacadeAutowire struct {
	mock.Mock
}

// GetAllStructDescriptors provides a mock function with given fields:
func (_m *FacadeAutowire) GetAllStructDescriptors() map[string]*autowire.StructDescriptor {
	ret := _m.Called()

	var r0 map[string]*autowire.StructDescriptor
	if rf, ok := ret.Get(0).(func() map[string]*autowire.StructDescriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*autowire.StructDescriptor)
		}
	}

	return r0
}

// TagKey provides a mock function with given fields:
func (_m *FacadeAutowire) TagKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewFacadeAutowire creates a new instance of FacadeAutowire. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewFacadeAutowire(t testing.TB) *FacadeAutowire {
	mock := &FacadeAutowire{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
