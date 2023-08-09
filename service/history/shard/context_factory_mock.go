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
// Source: context_factory.go

// Package shard is a generated GoMock package.
package shard

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockContextFactory is a mock of ContextFactory interface.
type MockContextFactory struct {
	ctrl     *gomock.Controller
	recorder *MockContextFactoryMockRecorder
}

// MockContextFactoryMockRecorder is the mock recorder for MockContextFactory.
type MockContextFactoryMockRecorder struct {
	mock *MockContextFactory
}

// NewMockContextFactory creates a new mock instance.
func NewMockContextFactory(ctrl *gomock.Controller) *MockContextFactory {
	mock := &MockContextFactory{ctrl: ctrl}
	mock.recorder = &MockContextFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextFactory) EXPECT() *MockContextFactoryMockRecorder {
	return m.recorder
}

// CreateContext mocks base method.
func (m *MockContextFactory) CreateContext(shardID int32, closeCallback CloseCallback) (ControllableContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContext", shardID, closeCallback)
	ret0, _ := ret[0].(ControllableContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContext indicates an expected call of CreateContext.
func (mr *MockContextFactoryMockRecorder) CreateContext(shardID, closeCallback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContext", reflect.TypeOf((*MockContextFactory)(nil).CreateContext), shardID, closeCallback)
}