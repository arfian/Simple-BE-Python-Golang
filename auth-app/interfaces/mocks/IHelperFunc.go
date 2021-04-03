package mocks

import mock "github.com/stretchr/testify/mock"

type IHelperFunc struct {
	mock.Mock
}

func (_m *IHelperFunc) GeneratePass() (string) {
	args := _m.Called()
	result := args.Get(0)
	return result.(string)
}

func (_m *IHelperFunc) GenerateUsername(name string) (string) {
	args := _m.Called(name)
	result := args.Get(0)
	return result.(string)
}