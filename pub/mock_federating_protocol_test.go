// Code generated by MockGen. DO NOT EDIT.
// Source: federating_protocol.go

// Package pub is a generated GoMock package.
package pub

import (
	context "context"
	vocab "github.com/go-fed/activity/streams/vocab"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockFederatingProtocol is a mock of FederatingProtocol interface
type MockFederatingProtocol struct {
	ctrl     *gomock.Controller
	recorder *MockFederatingProtocolMockRecorder
}

// MockFederatingProtocolMockRecorder is the mock recorder for MockFederatingProtocol
type MockFederatingProtocolMockRecorder struct {
	mock *MockFederatingProtocol
}

// NewMockFederatingProtocol creates a new mock instance
func NewMockFederatingProtocol(ctrl *gomock.Controller) *MockFederatingProtocol {
	mock := &MockFederatingProtocol{ctrl: ctrl}
	mock.recorder = &MockFederatingProtocolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFederatingProtocol) EXPECT() *MockFederatingProtocolMockRecorder {
	return m.recorder
}

// AuthenticatePostInbox mocks base method
func (m *MockFederatingProtocol) AuthenticatePostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticatePostInbox", c, w, r)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticatePostInbox indicates an expected call of AuthenticatePostInbox
func (mr *MockFederatingProtocolMockRecorder) AuthenticatePostInbox(c, w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticatePostInbox", reflect.TypeOf((*MockFederatingProtocol)(nil).AuthenticatePostInbox), c, w, r)
}

// Blocked mocks base method
func (m *MockFederatingProtocol) Blocked(c context.Context, actorIRIs []*url.URL) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Blocked", c, actorIRIs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Blocked indicates an expected call of Blocked
func (mr *MockFederatingProtocolMockRecorder) Blocked(c, actorIRIs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blocked", reflect.TypeOf((*MockFederatingProtocol)(nil).Blocked), c, actorIRIs)
}

// Callbacks mocks base method
func (m *MockFederatingProtocol) Callbacks(c context.Context) (FederatingWrappedCallbacks, []interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Callbacks", c)
	ret0, _ := ret[0].(FederatingWrappedCallbacks)
	ret1, _ := ret[1].([]interface{})
	return ret0, ret1
}

// Callbacks indicates an expected call of Callbacks
func (mr *MockFederatingProtocolMockRecorder) Callbacks(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Callbacks", reflect.TypeOf((*MockFederatingProtocol)(nil).Callbacks), c)
}

// MaxInboxForwardingRecursionDepth mocks base method
func (m *MockFederatingProtocol) MaxInboxForwardingRecursionDepth(c context.Context) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxInboxForwardingRecursionDepth", c)
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxInboxForwardingRecursionDepth indicates an expected call of MaxInboxForwardingRecursionDepth
func (mr *MockFederatingProtocolMockRecorder) MaxInboxForwardingRecursionDepth(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxInboxForwardingRecursionDepth", reflect.TypeOf((*MockFederatingProtocol)(nil).MaxInboxForwardingRecursionDepth), c)
}

// MaxDeliveryRecursionDepth mocks base method
func (m *MockFederatingProtocol) MaxDeliveryRecursionDepth(c context.Context) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxDeliveryRecursionDepth", c)
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxDeliveryRecursionDepth indicates an expected call of MaxDeliveryRecursionDepth
func (mr *MockFederatingProtocolMockRecorder) MaxDeliveryRecursionDepth(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxDeliveryRecursionDepth", reflect.TypeOf((*MockFederatingProtocol)(nil).MaxDeliveryRecursionDepth), c)
}

// FilterForwarding mocks base method
func (m *MockFederatingProtocol) FilterForwarding(c context.Context, potentialRecipients []*url.URL, a Activity) ([]*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterForwarding", c, potentialRecipients, a)
	ret0, _ := ret[0].([]*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterForwarding indicates an expected call of FilterForwarding
func (mr *MockFederatingProtocolMockRecorder) FilterForwarding(c, potentialRecipients, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterForwarding", reflect.TypeOf((*MockFederatingProtocol)(nil).FilterForwarding), c, potentialRecipients, a)
}

// GetInbox mocks base method
func (m *MockFederatingProtocol) GetInbox(c context.Context, r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInbox", c, r)
	ret0, _ := ret[0].(vocab.ActivityStreamsOrderedCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInbox indicates an expected call of GetInbox
func (mr *MockFederatingProtocolMockRecorder) GetInbox(c, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInbox", reflect.TypeOf((*MockFederatingProtocol)(nil).GetInbox), c, r)
}