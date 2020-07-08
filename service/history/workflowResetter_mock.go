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

// Code generated by MockGen. DO NOT EDIT.
// Source: workflowResetter.go

// Package history is a generated GoMock package.
package history

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	history "go.temporal.io/temporal-proto/history/v1"
)

// MockworkflowResetter is a mock of workflowResetter interface.
type MockworkflowResetter struct {
	ctrl     *gomock.Controller
	recorder *MockworkflowResetterMockRecorder
}

// MockworkflowResetterMockRecorder is the mock recorder for MockworkflowResetter.
type MockworkflowResetterMockRecorder struct {
	mock *MockworkflowResetter
}

// NewMockworkflowResetter creates a new mock instance.
func NewMockworkflowResetter(ctrl *gomock.Controller) *MockworkflowResetter {
	mock := &MockworkflowResetter{ctrl: ctrl}
	mock.recorder = &MockworkflowResetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockworkflowResetter) EXPECT() *MockworkflowResetterMockRecorder {
	return m.recorder
}

// resetWorkflow mocks base method.
func (m *MockworkflowResetter) resetWorkflow(ctx context.Context, namespaceID, workflowID, baseRunID string, baseBranchToken []byte, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID int64, resetRunID, resetRequestID string, currentWorkflow nDCWorkflow, resetReason string, additionalReapplyEvents []*history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "resetWorkflow", ctx, namespaceID, workflowID, baseRunID, baseBranchToken, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID, resetRunID, resetRequestID, currentWorkflow, resetReason, additionalReapplyEvents)
	ret0, _ := ret[0].(error)
	return ret0
}

// resetWorkflow indicates an expected call of resetWorkflow.
func (mr *MockworkflowResetterMockRecorder) resetWorkflow(ctx, namespaceID, workflowID, baseRunID, baseBranchToken, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID, resetRunID, resetRequestID, currentWorkflow, resetReason, additionalReapplyEvents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "resetWorkflow", reflect.TypeOf((*MockworkflowResetter)(nil).resetWorkflow), ctx, namespaceID, workflowID, baseRunID, baseBranchToken, baseRebuildLastEventID, baseRebuildLastEventVersion, baseNextEventID, resetRunID, resetRequestID, currentWorkflow, resetReason, additionalReapplyEvents)
}
