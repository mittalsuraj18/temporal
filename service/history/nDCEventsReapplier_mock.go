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
// Source: nDCEventsReapplier.go

// Package history is a generated GoMock package.
package history

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	history "go.temporal.io/api/history/v1"
)

// MocknDCEventsReapplier is a mock of nDCEventsReapplier interface
type MocknDCEventsReapplier struct {
	ctrl     *gomock.Controller
	recorder *MocknDCEventsReapplierMockRecorder
}

// MocknDCEventsReapplierMockRecorder is the mock recorder for MocknDCEventsReapplier
type MocknDCEventsReapplierMockRecorder struct {
	mock *MocknDCEventsReapplier
}

// NewMocknDCEventsReapplier creates a new mock instance
func NewMocknDCEventsReapplier(ctrl *gomock.Controller) *MocknDCEventsReapplier {
	mock := &MocknDCEventsReapplier{ctrl: ctrl}
	mock.recorder = &MocknDCEventsReapplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocknDCEventsReapplier) EXPECT() *MocknDCEventsReapplierMockRecorder {
	return m.recorder
}

// reapplyEvents mocks base method
func (m *MocknDCEventsReapplier) reapplyEvents(ctx context.Context, msBuilder mutableState, historyEvents []*history.HistoryEvent, runID string) ([]*history.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "reapplyEvents", ctx, msBuilder, historyEvents, runID)
	ret0, _ := ret[0].([]*history.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// reapplyEvents indicates an expected call of reapplyEvents
func (mr *MocknDCEventsReapplierMockRecorder) reapplyEvents(ctx, msBuilder, historyEvents, runID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "reapplyEvents", reflect.TypeOf((*MocknDCEventsReapplier)(nil).reapplyEvents), ctx, msBuilder, historyEvents, runID)
}
