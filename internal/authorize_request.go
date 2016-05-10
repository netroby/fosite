// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: AuthorizeRequester)

package internal

import (
	url "net/url"
	time "time"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
)

// Mock of AuthorizeRequester interface
type MockAuthorizeRequester struct {
	ctrl     *gomock.Controller
	recorder *_MockAuthorizeRequesterRecorder
}

// Recorder for MockAuthorizeRequester (not exported)
type _MockAuthorizeRequesterRecorder struct {
	mock *MockAuthorizeRequester
}

func NewMockAuthorizeRequester(ctrl *gomock.Controller) *MockAuthorizeRequester {
	mock := &MockAuthorizeRequester{ctrl: ctrl}
	mock.recorder = &_MockAuthorizeRequesterRecorder{mock}
	return mock
}

func (_m *MockAuthorizeRequester) EXPECT() *_MockAuthorizeRequesterRecorder {
	return _m.recorder
}

func (_m *MockAuthorizeRequester) DidHandleAllResponseTypes() bool {
	ret := _m.ctrl.Call(_m, "DidHandleAllResponseTypes")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) DidHandleAllResponseTypes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DidHandleAllResponseTypes")
}

func (_m *MockAuthorizeRequester) GetClient() fosite.Client {
	ret := _m.ctrl.Call(_m, "GetClient")
	ret0, _ := ret[0].(fosite.Client)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClient")
}

func (_m *MockAuthorizeRequester) GetGrantedScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetGrantedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetGrantedScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGrantedScopes")
}

func (_m *MockAuthorizeRequester) GetRedirectURI() *url.URL {
	ret := _m.ctrl.Call(_m, "GetRedirectURI")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetRedirectURI() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRedirectURI")
}

func (_m *MockAuthorizeRequester) GetRequestForm() url.Values {
	ret := _m.ctrl.Call(_m, "GetRequestForm")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetRequestForm() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestForm")
}

func (_m *MockAuthorizeRequester) GetRequestedAt() time.Time {
	ret := _m.ctrl.Call(_m, "GetRequestedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetRequestedAt() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestedAt")
}

func (_m *MockAuthorizeRequester) GetResponseTypes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetResponseTypes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetResponseTypes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetResponseTypes")
}

func (_m *MockAuthorizeRequester) GetScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetScopes")
}

func (_m *MockAuthorizeRequester) GetSession() interface{} {
	ret := _m.ctrl.Call(_m, "GetSession")
	ret0, _ := ret[0].(interface{})
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetSession() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSession")
}

func (_m *MockAuthorizeRequester) GetState() string {
	ret := _m.ctrl.Call(_m, "GetState")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) GetState() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetState")
}

func (_m *MockAuthorizeRequester) GrantScope(_param0 string) {
	_m.ctrl.Call(_m, "GrantScope", _param0)
}

func (_mr *_MockAuthorizeRequesterRecorder) GrantScope(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GrantScope", arg0)
}

func (_m *MockAuthorizeRequester) IsRedirectURIValid() bool {
	ret := _m.ctrl.Call(_m, "IsRedirectURIValid")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockAuthorizeRequesterRecorder) IsRedirectURIValid() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsRedirectURIValid")
}

func (_m *MockAuthorizeRequester) Merge(_param0 fosite.Requester) {
	_m.ctrl.Call(_m, "Merge", _param0)
}

func (_mr *_MockAuthorizeRequesterRecorder) Merge(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Merge", arg0)
}

func (_m *MockAuthorizeRequester) SetResponseTypeHandled(_param0 string) {
	_m.ctrl.Call(_m, "SetResponseTypeHandled", _param0)
}

func (_mr *_MockAuthorizeRequesterRecorder) SetResponseTypeHandled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetResponseTypeHandled", arg0)
}

func (_m *MockAuthorizeRequester) SetScopes(_param0 fosite.Arguments) {
	_m.ctrl.Call(_m, "SetScopes", _param0)
}

func (_mr *_MockAuthorizeRequesterRecorder) SetScopes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetScopes", arg0)
}

func (_m *MockAuthorizeRequester) SetSession(_param0 interface{}) {
	_m.ctrl.Call(_m, "SetSession", _param0)
}

func (_mr *_MockAuthorizeRequesterRecorder) SetSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSession", arg0)
}