package mocks

import "github.com/stretchr/testify/mock"

// DiagnosticsMessageSubscriber is an autogenerated mock type for the DiagnosticsMessageSubscriber type
type DiagnosticsMessageListener struct {
	mock.Mock
}

func (_m *DiagnosticsMessageListener) Remove() {
	_m.Called()
}
