// Code generated by counterfeiter. DO NOT EDIT.
package servicesfakes

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/services"
	"go_todo_api/src/shared/entities"
	"sync"
)

type FakeAuthService struct {
	LoginStub        func(dto.LoginRequest) entities.BaseResponse[dto.LoginResponse]
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 dto.LoginRequest
	}
	loginReturns struct {
		result1 entities.BaseResponse[dto.LoginResponse]
	}
	loginReturnsOnCall map[int]struct {
		result1 entities.BaseResponse[dto.LoginResponse]
	}
	RegisterStub        func(dto.RegisterRequest) entities.BaseResponse[dto.RegisterResponse]
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 dto.RegisterRequest
	}
	registerReturns struct {
		result1 entities.BaseResponse[dto.RegisterResponse]
	}
	registerReturnsOnCall map[int]struct {
		result1 entities.BaseResponse[dto.RegisterResponse]
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthService) Login(arg1 dto.LoginRequest) entities.BaseResponse[dto.LoginResponse] {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 dto.LoginRequest
	}{arg1})
	stub := fake.LoginStub
	fakeReturns := fake.loginReturns
	fake.recordInvocation("Login", []interface{}{arg1})
	fake.loginMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthService) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeAuthService) LoginCalls(stub func(dto.LoginRequest) entities.BaseResponse[dto.LoginResponse]) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeAuthService) LoginArgsForCall(i int) dto.LoginRequest {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthService) LoginReturns(result1 entities.BaseResponse[dto.LoginResponse]) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 entities.BaseResponse[dto.LoginResponse]
	}{result1}
}

func (fake *FakeAuthService) LoginReturnsOnCall(i int, result1 entities.BaseResponse[dto.LoginResponse]) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 entities.BaseResponse[dto.LoginResponse]
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 entities.BaseResponse[dto.LoginResponse]
	}{result1}
}

func (fake *FakeAuthService) Register(arg1 dto.RegisterRequest) entities.BaseResponse[dto.RegisterResponse] {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 dto.RegisterRequest
	}{arg1})
	stub := fake.RegisterStub
	fakeReturns := fake.registerReturns
	fake.recordInvocation("Register", []interface{}{arg1})
	fake.registerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthService) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeAuthService) RegisterCalls(stub func(dto.RegisterRequest) entities.BaseResponse[dto.RegisterResponse]) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = stub
}

func (fake *FakeAuthService) RegisterArgsForCall(i int) dto.RegisterRequest {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	argsForCall := fake.registerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAuthService) RegisterReturns(result1 entities.BaseResponse[dto.RegisterResponse]) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 entities.BaseResponse[dto.RegisterResponse]
	}{result1}
}

func (fake *FakeAuthService) RegisterReturnsOnCall(i int, result1 entities.BaseResponse[dto.RegisterResponse]) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 entities.BaseResponse[dto.RegisterResponse]
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 entities.BaseResponse[dto.RegisterResponse]
	}{result1}
}

func (fake *FakeAuthService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ services.AuthService = new(FakeAuthService)