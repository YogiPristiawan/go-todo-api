// Code generated by counterfeiter. DO NOT EDIT.
package servicesfakes

import (
	"go_todo_api/src/account/dto"
	"go_todo_api/src/account/services"
	"go_todo_api/src/shared/entities"
	"sync"
)

type FakeAccountService struct {
	GetProfileStub        func(dto.ProfileRequest) entities.BaseResponse[dto.ProfileResponse]
	getProfileMutex       sync.RWMutex
	getProfileArgsForCall []struct {
		arg1 dto.ProfileRequest
	}
	getProfileReturns struct {
		result1 entities.BaseResponse[dto.ProfileResponse]
	}
	getProfileReturnsOnCall map[int]struct {
		result1 entities.BaseResponse[dto.ProfileResponse]
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccountService) GetProfile(arg1 dto.ProfileRequest) entities.BaseResponse[dto.ProfileResponse] {
	fake.getProfileMutex.Lock()
	ret, specificReturn := fake.getProfileReturnsOnCall[len(fake.getProfileArgsForCall)]
	fake.getProfileArgsForCall = append(fake.getProfileArgsForCall, struct {
		arg1 dto.ProfileRequest
	}{arg1})
	stub := fake.GetProfileStub
	fakeReturns := fake.getProfileReturns
	fake.recordInvocation("GetProfile", []interface{}{arg1})
	fake.getProfileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAccountService) GetProfileCallCount() int {
	fake.getProfileMutex.RLock()
	defer fake.getProfileMutex.RUnlock()
	return len(fake.getProfileArgsForCall)
}

func (fake *FakeAccountService) GetProfileCalls(stub func(dto.ProfileRequest) entities.BaseResponse[dto.ProfileResponse]) {
	fake.getProfileMutex.Lock()
	defer fake.getProfileMutex.Unlock()
	fake.GetProfileStub = stub
}

func (fake *FakeAccountService) GetProfileArgsForCall(i int) dto.ProfileRequest {
	fake.getProfileMutex.RLock()
	defer fake.getProfileMutex.RUnlock()
	argsForCall := fake.getProfileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccountService) GetProfileReturns(result1 entities.BaseResponse[dto.ProfileResponse]) {
	fake.getProfileMutex.Lock()
	defer fake.getProfileMutex.Unlock()
	fake.GetProfileStub = nil
	fake.getProfileReturns = struct {
		result1 entities.BaseResponse[dto.ProfileResponse]
	}{result1}
}

func (fake *FakeAccountService) GetProfileReturnsOnCall(i int, result1 entities.BaseResponse[dto.ProfileResponse]) {
	fake.getProfileMutex.Lock()
	defer fake.getProfileMutex.Unlock()
	fake.GetProfileStub = nil
	if fake.getProfileReturnsOnCall == nil {
		fake.getProfileReturnsOnCall = make(map[int]struct {
			result1 entities.BaseResponse[dto.ProfileResponse]
		})
	}
	fake.getProfileReturnsOnCall[i] = struct {
		result1 entities.BaseResponse[dto.ProfileResponse]
	}{result1}
}

func (fake *FakeAccountService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getProfileMutex.RLock()
	defer fake.getProfileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAccountService) recordInvocation(key string, args []interface{}) {
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

var _ services.AccountService = new(FakeAccountService)
