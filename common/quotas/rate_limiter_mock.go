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
// Source: rate_limiter.go

// Package quotas is a generated GoMock package.
package quotas

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRateLimiter is a mock of RateLimiter interface
type MockRateLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockRateLimiterMockRecorder
}

// MockRateLimiterMockRecorder is the mock recorder for MockRateLimiter
type MockRateLimiterMockRecorder struct {
	mock *MockRateLimiter
}

// NewMockRateLimiter creates a new mock instance
func NewMockRateLimiter(ctrl *gomock.Controller) *MockRateLimiter {
	mock := &MockRateLimiter{ctrl: ctrl}
	mock.recorder = &MockRateLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRateLimiter) EXPECT() *MockRateLimiterMockRecorder {
	return m.recorder
}

// Allow mocks base method
func (m *MockRateLimiter) Allow() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allow")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Allow indicates an expected call of Allow
func (mr *MockRateLimiterMockRecorder) Allow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allow", reflect.TypeOf((*MockRateLimiter)(nil).Allow))
}

// AllowN mocks base method
func (m *MockRateLimiter) AllowN(now time.Time, numToken int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllowN", now, numToken)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AllowN indicates an expected call of AllowN
func (mr *MockRateLimiterMockRecorder) AllowN(now, numToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowN", reflect.TypeOf((*MockRateLimiter)(nil).AllowN), now, numToken)
}

// Reserve mocks base method
func (m *MockRateLimiter) Reserve() Reservation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reserve")
	ret0, _ := ret[0].(Reservation)
	return ret0
}

// Reserve indicates an expected call of Reserve
func (mr *MockRateLimiterMockRecorder) Reserve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reserve", reflect.TypeOf((*MockRateLimiter)(nil).Reserve))
}

// ReserveN mocks base method
func (m *MockRateLimiter) ReserveN(now time.Time, numToken int) Reservation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveN", now, numToken)
	ret0, _ := ret[0].(Reservation)
	return ret0
}

// ReserveN indicates an expected call of ReserveN
func (mr *MockRateLimiterMockRecorder) ReserveN(now, numToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveN", reflect.TypeOf((*MockRateLimiter)(nil).ReserveN), now, numToken)
}

// Wait mocks base method
func (m *MockRateLimiter) Wait(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockRateLimiterMockRecorder) Wait(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockRateLimiter)(nil).Wait), ctx)
}

// WaitN mocks base method
func (m *MockRateLimiter) WaitN(ctx context.Context, numToken int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitN", ctx, numToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitN indicates an expected call of WaitN
func (mr *MockRateLimiterMockRecorder) WaitN(ctx, numToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitN", reflect.TypeOf((*MockRateLimiter)(nil).WaitN), ctx, numToken)
}

// Rate mocks base method
func (m *MockRateLimiter) Rate() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rate")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Rate indicates an expected call of Rate
func (mr *MockRateLimiterMockRecorder) Rate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rate", reflect.TypeOf((*MockRateLimiter)(nil).Rate))
}

// Burst mocks base method
func (m *MockRateLimiter) Burst() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Burst")
	ret0, _ := ret[0].(int)
	return ret0
}

// Burst indicates an expected call of Burst
func (mr *MockRateLimiterMockRecorder) Burst() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Burst", reflect.TypeOf((*MockRateLimiter)(nil).Burst))
}

// MockReservation is a mock of Reservation interface
type MockReservation struct {
	ctrl     *gomock.Controller
	recorder *MockReservationMockRecorder
}

// MockReservationMockRecorder is the mock recorder for MockReservation
type MockReservationMockRecorder struct {
	mock *MockReservation
}

// NewMockReservation creates a new mock instance
func NewMockReservation(ctrl *gomock.Controller) *MockReservation {
	mock := &MockReservation{ctrl: ctrl}
	mock.recorder = &MockReservationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReservation) EXPECT() *MockReservationMockRecorder {
	return m.recorder
}

// OK mocks base method
func (m *MockReservation) OK() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OK")
	ret0, _ := ret[0].(bool)
	return ret0
}

// OK indicates an expected call of OK
func (mr *MockReservationMockRecorder) OK() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OK", reflect.TypeOf((*MockReservation)(nil).OK))
}

// Cancel mocks base method
func (m *MockReservation) Cancel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cancel")
}

// Cancel indicates an expected call of Cancel
func (mr *MockReservationMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockReservation)(nil).Cancel))
}

// CancelAt mocks base method
func (m *MockReservation) CancelAt(now time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelAt", now)
}

// CancelAt indicates an expected call of CancelAt
func (mr *MockReservationMockRecorder) CancelAt(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelAt", reflect.TypeOf((*MockReservation)(nil).CancelAt), now)
}

// Delay mocks base method
func (m *MockReservation) Delay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Delay indicates an expected call of Delay
func (mr *MockReservationMockRecorder) Delay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockReservation)(nil).Delay))
}

// DelayFrom mocks base method
func (m *MockReservation) DelayFrom(now time.Time) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelayFrom", now)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DelayFrom indicates an expected call of DelayFrom
func (mr *MockReservationMockRecorder) DelayFrom(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelayFrom", reflect.TypeOf((*MockReservation)(nil).DelayFrom), now)
}
