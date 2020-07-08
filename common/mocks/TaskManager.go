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

// TaskManager is an autogenerated mock type for the TaskManager type
type TaskManager struct {
	mock.Mock
}

// GetName provides a mock function with given fields:
func (_m *TaskManager) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *TaskManager) Close() {
	_m.Called()
}

// LeaseTaskQueue provides a mock function with given fields: request
func (_m *TaskManager) LeaseTaskQueue(request *persistence.LeaseTaskQueueRequest) (*persistence.LeaseTaskQueueResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.LeaseTaskQueueResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*persistence.LeaseTaskQueueRequest) (*persistence.LeaseTaskQueueResponse, error)); ok {
		return rf(request)
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).(*persistence.LeaseTaskQueueResponse)
	}

	if rf, ok := ret.Get(1).(func(*persistence.LeaseTaskQueueRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTaskQueue provides a mock function with given fields: request
func (_m *TaskManager) UpdateTaskQueue(request *persistence.UpdateTaskQueueRequest) (*persistence.UpdateTaskQueueResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.UpdateTaskQueueResponse
	if rf, ok := ret.Get(0).(func(*persistence.UpdateTaskQueueRequest) *persistence.UpdateTaskQueueResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.UpdateTaskQueueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.UpdateTaskQueueRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteTask provides a mock function with given fields: request
func (_m *TaskManager) CompleteTask(request *persistence.CompleteTaskRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.CompleteTaskRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteTasksLessThan
func (_m *TaskManager) CompleteTasksLessThan(request *persistence.CompleteTasksLessThanRequest) (int, error) {
	ret := _m.Called(request)

	var r0 int
	if rf, ok := ret.Get(0).(func(*persistence.CompleteTasksLessThanRequest) int); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.CompleteTasksLessThanRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TaskManager) ListTaskQueue(request *persistence.ListTaskQueueRequest) (*persistence.ListTaskQueueResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListTaskQueueResponse
	if rf, ok := ret.Get(0).(func(request *persistence.ListTaskQueueRequest) *persistence.ListTaskQueueResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListTaskQueueResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListTaskQueueRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TaskManager) DeleteTaskQueue(request *persistence.DeleteTaskQueueRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.DeleteTaskQueueRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTasks provides a mock function with given fields: request
func (_m *TaskManager) CreateTasks(request *persistence.CreateTasksRequest) (*persistence.CreateTasksResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.CreateTasksResponse
	if rf, ok := ret.Get(0).(func(*persistence.CreateTasksRequest) *persistence.CreateTasksResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.CreateTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.CreateTasksRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields: request
func (_m *TaskManager) GetTasks(request *persistence.GetTasksRequest) (*persistence.GetTasksResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetTasksResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetTasksRequest) *persistence.GetTasksResponse); ok {
		r0 = rf(request)
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).(*persistence.GetTasksResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetTasksRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
