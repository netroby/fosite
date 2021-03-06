// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: RefreshTokenStrategy)

// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of RefreshTokenStrategy interface
type MockRefreshTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockRefreshTokenStrategyRecorder
}

// Recorder for MockRefreshTokenStrategy (not exported)
type _MockRefreshTokenStrategyRecorder struct {
	mock *MockRefreshTokenStrategy
}

func NewMockRefreshTokenStrategy(ctrl *gomock.Controller) *MockRefreshTokenStrategy {
	mock := &MockRefreshTokenStrategy{ctrl: ctrl}
	mock.recorder = &_MockRefreshTokenStrategyRecorder{mock}
	return mock
}

func (_m *MockRefreshTokenStrategy) EXPECT() *_MockRefreshTokenStrategyRecorder {
	return _m.recorder
}

func (_m *MockRefreshTokenStrategy) GenerateRefreshToken(_param0 context.Context, _param1 fosite.Requester) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GenerateRefreshToken", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockRefreshTokenStrategyRecorder) GenerateRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateRefreshToken", arg0, arg1)
}

func (_m *MockRefreshTokenStrategy) RefreshTokenSignature(_param0 string) string {
	ret := _m.ctrl.Call(_m, "RefreshTokenSignature", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRefreshTokenStrategyRecorder) RefreshTokenSignature(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RefreshTokenSignature", arg0)
}

func (_m *MockRefreshTokenStrategy) ValidateRefreshToken(_param0 context.Context, _param1 fosite.Requester, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ValidateRefreshToken", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRefreshTokenStrategyRecorder) ValidateRefreshToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateRefreshToken", arg0, arg1, arg2)
}
