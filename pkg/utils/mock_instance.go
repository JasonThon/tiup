package utils



import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v0manifest "github.com/pingcap/tiup/pkg/repository/v0manifest"
	reflect "reflect"

)

// MockInstance is a mock of Instance interface
type MockInstance struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMockRecorder
}

// MockInstanceMockRecorder is the mock recorder for MockInstance
type MockInstanceMockRecorder struct {
	mock *MockInstance
}

// NewMockInstance creates a new mock instance
func NewMockInstance(ctrl *gomock.Controller) *MockInstance {
	mock := &MockInstance{ctrl: ctrl}
	mock.recorder = &MockInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstance) EXPECT() *MockInstanceMockRecorder {
	return m.recorder
}

// Component mocks base method
func (m *MockInstance) Component() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Component")
	ret0, _ := ret[0].(string)
	return ret0
}

// Component indicates an expected call of Component
func (mr *MockInstanceMockRecorder) Component() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Component", reflect.TypeOf((*MockInstance)(nil).Component))
}

// LogFile mocks base method
func (m *MockInstance) LogFile() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogFile")
	ret0, _ := ret[0].(string)
	return ret0
}

// LogFile indicates an expected call of LogFile
func (mr *MockInstanceMockRecorder) LogFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogFile", reflect.TypeOf((*MockInstance)(nil).LogFile))
}

// Pid mocks base method
func (m *MockInstance) Pid() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pid")
	ret0, _ := ret[0].(int)
	return ret0
}

// Pid indicates an expected call of Pid
func (mr *MockInstanceMockRecorder) Pid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pid", reflect.TypeOf((*MockInstance)(nil).Pid))
}

// Start mocks base method
func (m *MockInstance) Start(arg0 context.Context, arg1 v0manifest.Version) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockInstanceMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockInstance)(nil).Start), arg0, arg1)
}

// StatusAddrs mocks base method
func (m *MockInstance) StatusAddrs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusAddrs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// StatusAddrs indicates an expected call of StatusAddrs
func (mr *MockInstanceMockRecorder) StatusAddrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusAddrs", reflect.TypeOf((*MockInstance)(nil).StatusAddrs))
}

// Uptime mocks base method
func (m *MockInstance) Uptime() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uptime")
	ret0, _ := ret[0].(string)
	return ret0
}

// Uptime indicates an expected call of Uptime
func (mr *MockInstanceMockRecorder) Uptime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uptime", reflect.TypeOf((*MockInstance)(nil).Uptime))
}

// Wait mocks base method
func (m *MockInstance) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockInstanceMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockInstance)(nil).Wait))
}

