// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mocks

import (
	"github.com/stretchr/testify/mock"

	"go.temporal.io/server/common/persistence"
)

// ExecutionManagerFactory is an autogenerated mock type for the ExecutionManagerFactory type
type ExecutionManagerFactory struct {
	mock.Mock
}

// NewExecutionManager provides a mock function with given fields: shardID
func (_m *ExecutionManagerFactory) NewExecutionManager(shardID int) (persistence.ExecutionManager, error) {
	ret := _m.Called(shardID)

	var r0 persistence.ExecutionManager
	if rf, ok := ret.Get(0).(func(int) persistence.ExecutionManager); ok {
		r0 = rf(shardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(persistence.ExecutionManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(shardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close is mock implementation for Close of ExecutionManagerFactory
func (_m *ExecutionManagerFactory) Close() {
	_m.Called()
}
