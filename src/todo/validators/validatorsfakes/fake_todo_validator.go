// Code generated by counterfeiter. DO NOT EDIT.
package validatorsfakes

import (
	"go_todo_api/src/todo/dto"
	"go_todo_api/src/todo/validators"
	"sync"
)

type FakeTodoValidator struct {
	ValidateDetailStub        func(dto.DetailTodoRequest) error
	validateDetailMutex       sync.RWMutex
	validateDetailArgsForCall []struct {
		arg1 dto.DetailTodoRequest
	}
	validateDetailReturns struct {
		result1 error
	}
	validateDetailReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateStoreStub        func(dto.StoreTodoRequest) error
	validateStoreMutex       sync.RWMutex
	validateStoreArgsForCall []struct {
		arg1 dto.StoreTodoRequest
	}
	validateStoreReturns struct {
		result1 error
	}
	validateStoreReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTodoValidator) ValidateDetail(arg1 dto.DetailTodoRequest) error {
	fake.validateDetailMutex.Lock()
	ret, specificReturn := fake.validateDetailReturnsOnCall[len(fake.validateDetailArgsForCall)]
	fake.validateDetailArgsForCall = append(fake.validateDetailArgsForCall, struct {
		arg1 dto.DetailTodoRequest
	}{arg1})
	stub := fake.ValidateDetailStub
	fakeReturns := fake.validateDetailReturns
	fake.recordInvocation("ValidateDetail", []interface{}{arg1})
	fake.validateDetailMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTodoValidator) ValidateDetailCallCount() int {
	fake.validateDetailMutex.RLock()
	defer fake.validateDetailMutex.RUnlock()
	return len(fake.validateDetailArgsForCall)
}

func (fake *FakeTodoValidator) ValidateDetailCalls(stub func(dto.DetailTodoRequest) error) {
	fake.validateDetailMutex.Lock()
	defer fake.validateDetailMutex.Unlock()
	fake.ValidateDetailStub = stub
}

func (fake *FakeTodoValidator) ValidateDetailArgsForCall(i int) dto.DetailTodoRequest {
	fake.validateDetailMutex.RLock()
	defer fake.validateDetailMutex.RUnlock()
	argsForCall := fake.validateDetailArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTodoValidator) ValidateDetailReturns(result1 error) {
	fake.validateDetailMutex.Lock()
	defer fake.validateDetailMutex.Unlock()
	fake.ValidateDetailStub = nil
	fake.validateDetailReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTodoValidator) ValidateDetailReturnsOnCall(i int, result1 error) {
	fake.validateDetailMutex.Lock()
	defer fake.validateDetailMutex.Unlock()
	fake.ValidateDetailStub = nil
	if fake.validateDetailReturnsOnCall == nil {
		fake.validateDetailReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateDetailReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTodoValidator) ValidateStore(arg1 dto.StoreTodoRequest) error {
	fake.validateStoreMutex.Lock()
	ret, specificReturn := fake.validateStoreReturnsOnCall[len(fake.validateStoreArgsForCall)]
	fake.validateStoreArgsForCall = append(fake.validateStoreArgsForCall, struct {
		arg1 dto.StoreTodoRequest
	}{arg1})
	stub := fake.ValidateStoreStub
	fakeReturns := fake.validateStoreReturns
	fake.recordInvocation("ValidateStore", []interface{}{arg1})
	fake.validateStoreMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTodoValidator) ValidateStoreCallCount() int {
	fake.validateStoreMutex.RLock()
	defer fake.validateStoreMutex.RUnlock()
	return len(fake.validateStoreArgsForCall)
}

func (fake *FakeTodoValidator) ValidateStoreCalls(stub func(dto.StoreTodoRequest) error) {
	fake.validateStoreMutex.Lock()
	defer fake.validateStoreMutex.Unlock()
	fake.ValidateStoreStub = stub
}

func (fake *FakeTodoValidator) ValidateStoreArgsForCall(i int) dto.StoreTodoRequest {
	fake.validateStoreMutex.RLock()
	defer fake.validateStoreMutex.RUnlock()
	argsForCall := fake.validateStoreArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTodoValidator) ValidateStoreReturns(result1 error) {
	fake.validateStoreMutex.Lock()
	defer fake.validateStoreMutex.Unlock()
	fake.ValidateStoreStub = nil
	fake.validateStoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTodoValidator) ValidateStoreReturnsOnCall(i int, result1 error) {
	fake.validateStoreMutex.Lock()
	defer fake.validateStoreMutex.Unlock()
	fake.ValidateStoreStub = nil
	if fake.validateStoreReturnsOnCall == nil {
		fake.validateStoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateStoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTodoValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateDetailMutex.RLock()
	defer fake.validateDetailMutex.RUnlock()
	fake.validateStoreMutex.RLock()
	defer fake.validateStoreMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTodoValidator) recordInvocation(key string, args []interface{}) {
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

var _ validators.TodoValidator = new(FakeTodoValidator)
