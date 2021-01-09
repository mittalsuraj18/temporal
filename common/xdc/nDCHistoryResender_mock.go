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
// Source: nDCHistoryResender.go

// Package xdc is a generated GoMock package.
package xdc

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNDCHistoryResender is a mock of NDCHistoryResender interface
type MockNDCHistoryResender struct {
	ctrl     *gomock.Controller
	recorder *MockNDCHistoryResenderMockRecorder
}

// MockNDCHistoryResenderMockRecorder is the mock recorder for MockNDCHistoryResender
type MockNDCHistoryResenderMockRecorder struct {
	mock *MockNDCHistoryResender
}

// NewMockNDCHistoryResender creates a new mock instance
func NewMockNDCHistoryResender(ctrl *gomock.Controller) *MockNDCHistoryResender {
	mock := &MockNDCHistoryResender{ctrl: ctrl}
	mock.recorder = &MockNDCHistoryResenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNDCHistoryResender) EXPECT() *MockNDCHistoryResenderMockRecorder {
	return m.recorder
}

// SendSingleWorkflowHistory mocks base method
func (m *MockNDCHistoryResender) SendSingleWorkflowHistory(namespaceID, workflowID, runID string, startEventID, startEventVersion, endEventID, endEventVersion int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSingleWorkflowHistory", namespaceID, workflowID, runID, startEventID, startEventVersion, endEventID, endEventVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSingleWorkflowHistory indicates an expected call of SendSingleWorkflowHistory
func (mr *MockNDCHistoryResenderMockRecorder) SendSingleWorkflowHistory(namespaceID, workflowID, runID, startEventID, startEventVersion, endEventID, endEventVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSingleWorkflowHistory", reflect.TypeOf((*MockNDCHistoryResender)(nil).SendSingleWorkflowHistory), namespaceID, workflowID, runID, startEventID, startEventVersion, endEventID, endEventVersion)
}
