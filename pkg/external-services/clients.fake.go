// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"golang-rest-api-clean-architecture/pkg/external-services/authenticator"
	"golang-rest-api-clean-architecture/pkg/external-services/database"
	location_verifier "golang-rest-api-clean-architecture/pkg/external-services/location-verifier"
	service_bus "golang-rest-api-clean-architecture/pkg/external-services/service-bus"
	"sync"
)

// Ensure, that ClientSetMock does implement ClientSet.
// If this is not the case, regenerate this file with moq.
var _ ClientSet = &ClientSetMock{}

// ClientSetMock is a mock implementation of ClientSet.
//
// 	func TestSomethingThatUsesClientSet(t *testing.T) {
//
// 		// make and configure a mocked ClientSet
// 		mockedClientSet := &ClientSetMock{
// 			AuthenticationClientFunc: func() authenticator.AuthenticationClient {
// 				panic("mock out the AuthenticationClient method")
// 			},
// 			DbClientFunc: func() database.DbClient {
// 				panic("mock out the DbClient method")
// 			},
// 			LocationVerificationClientFunc: func() location_verifier.LocationVerificationClient {
// 				panic("mock out the LocationVerificationClient method")
// 			},
// 			ServiceBusClientFunc: func() service_bus.ServiceBusClient {
// 				panic("mock out the ServiceBusClient method")
// 			},
// 		}
//
// 		// use mockedClientSet in code that requires ClientSet
// 		// and then make assertions.
//
// 	}
type ClientSetMock struct {
	// AuthenticationClientFunc mocks the AuthenticationClient method.
	AuthenticationClientFunc func() authenticator.AuthenticationClient

	// DbClientFunc mocks the DbClient method.
	DbClientFunc func() database.DbClient

	// LocationVerificationClientFunc mocks the LocationVerificationClient method.
	LocationVerificationClientFunc func() location_verifier.LocationVerificationClient

	// ServiceBusClientFunc mocks the ServiceBusClient method.
	ServiceBusClientFunc func() service_bus.ServiceBusClient

	// calls tracks calls to the methods.
	calls struct {
		// AuthenticationClient holds details about calls to the AuthenticationClient method.
		AuthenticationClient []struct {
		}
		// DbClient holds details about calls to the DbClient method.
		DbClient []struct {
		}
		// LocationVerificationClient holds details about calls to the LocationVerificationClient method.
		LocationVerificationClient []struct {
		}
		// ServiceBusClient holds details about calls to the ServiceBusClient method.
		ServiceBusClient []struct {
		}
	}
	lockAuthenticationClient       sync.RWMutex
	lockDbClient                   sync.RWMutex
	lockLocationVerificationClient sync.RWMutex
	lockServiceBusClient           sync.RWMutex
}

// AuthenticationClient calls AuthenticationClientFunc.
func (mock *ClientSetMock) AuthenticationClient() authenticator.AuthenticationClient {
	if mock.AuthenticationClientFunc == nil {
		panic("ClientSetMock.AuthenticationClientFunc: method is nil but ClientSet.AuthenticationClient was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAuthenticationClient.Lock()
	mock.calls.AuthenticationClient = append(mock.calls.AuthenticationClient, callInfo)
	mock.lockAuthenticationClient.Unlock()
	return mock.AuthenticationClientFunc()
}

// AuthenticationClientCalls gets all the calls that were made to AuthenticationClient.
// Check the length with:
//     len(mockedClientSet.AuthenticationClientCalls())
func (mock *ClientSetMock) AuthenticationClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAuthenticationClient.RLock()
	calls = mock.calls.AuthenticationClient
	mock.lockAuthenticationClient.RUnlock()
	return calls
}

// DbClient calls DbClientFunc.
func (mock *ClientSetMock) DbClient() database.DbClient {
	if mock.DbClientFunc == nil {
		panic("ClientSetMock.DbClientFunc: method is nil but ClientSet.DbClient was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDbClient.Lock()
	mock.calls.DbClient = append(mock.calls.DbClient, callInfo)
	mock.lockDbClient.Unlock()
	return mock.DbClientFunc()
}

// DbClientCalls gets all the calls that were made to DbClient.
// Check the length with:
//     len(mockedClientSet.DbClientCalls())
func (mock *ClientSetMock) DbClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDbClient.RLock()
	calls = mock.calls.DbClient
	mock.lockDbClient.RUnlock()
	return calls
}

// LocationVerificationClient calls LocationVerificationClientFunc.
func (mock *ClientSetMock) LocationVerificationClient() location_verifier.LocationVerificationClient {
	if mock.LocationVerificationClientFunc == nil {
		panic("ClientSetMock.LocationVerificationClientFunc: method is nil but ClientSet.LocationVerificationClient was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLocationVerificationClient.Lock()
	mock.calls.LocationVerificationClient = append(mock.calls.LocationVerificationClient, callInfo)
	mock.lockLocationVerificationClient.Unlock()
	return mock.LocationVerificationClientFunc()
}

// LocationVerificationClientCalls gets all the calls that were made to LocationVerificationClient.
// Check the length with:
//     len(mockedClientSet.LocationVerificationClientCalls())
func (mock *ClientSetMock) LocationVerificationClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLocationVerificationClient.RLock()
	calls = mock.calls.LocationVerificationClient
	mock.lockLocationVerificationClient.RUnlock()
	return calls
}

// ServiceBusClient calls ServiceBusClientFunc.
func (mock *ClientSetMock) ServiceBusClient() service_bus.ServiceBusClient {
	if mock.ServiceBusClientFunc == nil {
		panic("ClientSetMock.ServiceBusClientFunc: method is nil but ClientSet.ServiceBusClient was just called")
	}
	callInfo := struct {
	}{}
	mock.lockServiceBusClient.Lock()
	mock.calls.ServiceBusClient = append(mock.calls.ServiceBusClient, callInfo)
	mock.lockServiceBusClient.Unlock()
	return mock.ServiceBusClientFunc()
}

// ServiceBusClientCalls gets all the calls that were made to ServiceBusClient.
// Check the length with:
//     len(mockedClientSet.ServiceBusClientCalls())
func (mock *ClientSetMock) ServiceBusClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockServiceBusClient.RLock()
	calls = mock.calls.ServiceBusClient
	mock.lockServiceBusClient.RUnlock()
	return calls
}
