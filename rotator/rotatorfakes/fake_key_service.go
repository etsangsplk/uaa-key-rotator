// Code generated by counterfeiter. DO NOT EDIT.
package rotatorfakes

import (
	"sync"

	"github.com/cloudfoundry/uaa-key-rotator/crypto"
	"github.com/cloudfoundry/uaa-key-rotator/rotator"
)

type FakeKeyService struct {
	KeyStub        func(keyLabel string) crypto.Decryptor
	keyMutex       sync.RWMutex
	keyArgsForCall []struct {
		keyLabel string
	}
	keyReturns struct {
		result1 crypto.Decryptor
	}
	keyReturnsOnCall map[int]struct {
		result1 crypto.Decryptor
	}
	ActiveKeyStub        func() (string, crypto.Encryptor)
	activeKeyMutex       sync.RWMutex
	activeKeyArgsForCall []struct{}
	activeKeyReturns     struct {
		result1 string
		result2 crypto.Encryptor
	}
	activeKeyReturnsOnCall map[int]struct {
		result1 string
		result2 crypto.Encryptor
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKeyService) Key(keyLabel string) crypto.Decryptor {
	fake.keyMutex.Lock()
	ret, specificReturn := fake.keyReturnsOnCall[len(fake.keyArgsForCall)]
	fake.keyArgsForCall = append(fake.keyArgsForCall, struct {
		keyLabel string
	}{keyLabel})
	fake.recordInvocation("Key", []interface{}{keyLabel})
	fake.keyMutex.Unlock()
	if fake.KeyStub != nil {
		return fake.KeyStub(keyLabel)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.keyReturns.result1
}

func (fake *FakeKeyService) KeyCallCount() int {
	fake.keyMutex.RLock()
	defer fake.keyMutex.RUnlock()
	return len(fake.keyArgsForCall)
}

func (fake *FakeKeyService) KeyArgsForCall(i int) string {
	fake.keyMutex.RLock()
	defer fake.keyMutex.RUnlock()
	return fake.keyArgsForCall[i].keyLabel
}

func (fake *FakeKeyService) KeyReturns(result1 crypto.Decryptor) {
	fake.KeyStub = nil
	fake.keyReturns = struct {
		result1 crypto.Decryptor
	}{result1}
}

func (fake *FakeKeyService) KeyReturnsOnCall(i int, result1 crypto.Decryptor) {
	fake.KeyStub = nil
	if fake.keyReturnsOnCall == nil {
		fake.keyReturnsOnCall = make(map[int]struct {
			result1 crypto.Decryptor
		})
	}
	fake.keyReturnsOnCall[i] = struct {
		result1 crypto.Decryptor
	}{result1}
}

func (fake *FakeKeyService) ActiveKey() (string, crypto.Encryptor) {
	fake.activeKeyMutex.Lock()
	ret, specificReturn := fake.activeKeyReturnsOnCall[len(fake.activeKeyArgsForCall)]
	fake.activeKeyArgsForCall = append(fake.activeKeyArgsForCall, struct{}{})
	fake.recordInvocation("ActiveKey", []interface{}{})
	fake.activeKeyMutex.Unlock()
	if fake.ActiveKeyStub != nil {
		return fake.ActiveKeyStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.activeKeyReturns.result1, fake.activeKeyReturns.result2
}

func (fake *FakeKeyService) ActiveKeyCallCount() int {
	fake.activeKeyMutex.RLock()
	defer fake.activeKeyMutex.RUnlock()
	return len(fake.activeKeyArgsForCall)
}

func (fake *FakeKeyService) ActiveKeyReturns(result1 string, result2 crypto.Encryptor) {
	fake.ActiveKeyStub = nil
	fake.activeKeyReturns = struct {
		result1 string
		result2 crypto.Encryptor
	}{result1, result2}
}

func (fake *FakeKeyService) ActiveKeyReturnsOnCall(i int, result1 string, result2 crypto.Encryptor) {
	fake.ActiveKeyStub = nil
	if fake.activeKeyReturnsOnCall == nil {
		fake.activeKeyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 crypto.Encryptor
		})
	}
	fake.activeKeyReturnsOnCall[i] = struct {
		result1 string
		result2 crypto.Encryptor
	}{result1, result2}
}

func (fake *FakeKeyService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.keyMutex.RLock()
	defer fake.keyMutex.RUnlock()
	fake.activeKeyMutex.RLock()
	defer fake.activeKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKeyService) recordInvocation(key string, args []interface{}) {
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

var _ rotator.KeyService = new(FakeKeyService)