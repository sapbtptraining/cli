// Code generated by counterfeiter. DO NOT EDIT.
package v7pushactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7pushaction"
	"code.cloudfoundry.org/cli/util/manifestparser"
)

type FakeManifestParser struct {
	AppsStub        func() []manifestparser.Application
	appsMutex       sync.RWMutex
	appsArgsForCall []struct {
	}
	appsReturns struct {
		result1 []manifestparser.Application
	}
	appsReturnsOnCall map[int]struct {
		result1 []manifestparser.Application
	}
	ContainsManifestStub        func() bool
	containsManifestMutex       sync.RWMutex
	containsManifestArgsForCall []struct {
	}
	containsManifestReturns struct {
		result1 bool
	}
	containsManifestReturnsOnCall map[int]struct {
		result1 bool
	}
	FullRawManifestStub        func() []byte
	fullRawManifestMutex       sync.RWMutex
	fullRawManifestArgsForCall []struct {
	}
	fullRawManifestReturns struct {
		result1 []byte
	}
	fullRawManifestReturnsOnCall map[int]struct {
		result1 []byte
	}
	RawAppManifestStub        func(string) ([]byte, error)
	rawAppManifestMutex       sync.RWMutex
	rawAppManifestArgsForCall []struct {
		arg1 string
	}
	rawAppManifestReturns struct {
		result1 []byte
		result2 error
	}
	rawAppManifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestParser) Apps() []manifestparser.Application {
	fake.appsMutex.Lock()
	ret, specificReturn := fake.appsReturnsOnCall[len(fake.appsArgsForCall)]
	fake.appsArgsForCall = append(fake.appsArgsForCall, struct {
	}{})
	fake.recordInvocation("Apps", []interface{}{})
	fake.appsMutex.Unlock()
	if fake.AppsStub != nil {
		return fake.AppsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.appsReturns
	return fakeReturns.result1
}

func (fake *FakeManifestParser) AppsCallCount() int {
	fake.appsMutex.RLock()
	defer fake.appsMutex.RUnlock()
	return len(fake.appsArgsForCall)
}

func (fake *FakeManifestParser) AppsCalls(stub func() []manifestparser.Application) {
	fake.appsMutex.Lock()
	defer fake.appsMutex.Unlock()
	fake.AppsStub = stub
}

func (fake *FakeManifestParser) AppsReturns(result1 []manifestparser.Application) {
	fake.appsMutex.Lock()
	defer fake.appsMutex.Unlock()
	fake.AppsStub = nil
	fake.appsReturns = struct {
		result1 []manifestparser.Application
	}{result1}
}

func (fake *FakeManifestParser) AppsReturnsOnCall(i int, result1 []manifestparser.Application) {
	fake.appsMutex.Lock()
	defer fake.appsMutex.Unlock()
	fake.AppsStub = nil
	if fake.appsReturnsOnCall == nil {
		fake.appsReturnsOnCall = make(map[int]struct {
			result1 []manifestparser.Application
		})
	}
	fake.appsReturnsOnCall[i] = struct {
		result1 []manifestparser.Application
	}{result1}
}

func (fake *FakeManifestParser) ContainsManifest() bool {
	fake.containsManifestMutex.Lock()
	ret, specificReturn := fake.containsManifestReturnsOnCall[len(fake.containsManifestArgsForCall)]
	fake.containsManifestArgsForCall = append(fake.containsManifestArgsForCall, struct {
	}{})
	fake.recordInvocation("ContainsManifest", []interface{}{})
	fake.containsManifestMutex.Unlock()
	if fake.ContainsManifestStub != nil {
		return fake.ContainsManifestStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.containsManifestReturns
	return fakeReturns.result1
}

func (fake *FakeManifestParser) ContainsManifestCallCount() int {
	fake.containsManifestMutex.RLock()
	defer fake.containsManifestMutex.RUnlock()
	return len(fake.containsManifestArgsForCall)
}

func (fake *FakeManifestParser) ContainsManifestCalls(stub func() bool) {
	fake.containsManifestMutex.Lock()
	defer fake.containsManifestMutex.Unlock()
	fake.ContainsManifestStub = stub
}

func (fake *FakeManifestParser) ContainsManifestReturns(result1 bool) {
	fake.containsManifestMutex.Lock()
	defer fake.containsManifestMutex.Unlock()
	fake.ContainsManifestStub = nil
	fake.containsManifestReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifestParser) ContainsManifestReturnsOnCall(i int, result1 bool) {
	fake.containsManifestMutex.Lock()
	defer fake.containsManifestMutex.Unlock()
	fake.ContainsManifestStub = nil
	if fake.containsManifestReturnsOnCall == nil {
		fake.containsManifestReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.containsManifestReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifestParser) FullRawManifest() []byte {
	fake.fullRawManifestMutex.Lock()
	ret, specificReturn := fake.fullRawManifestReturnsOnCall[len(fake.fullRawManifestArgsForCall)]
	fake.fullRawManifestArgsForCall = append(fake.fullRawManifestArgsForCall, struct {
	}{})
	fake.recordInvocation("FullRawManifest", []interface{}{})
	fake.fullRawManifestMutex.Unlock()
	if fake.FullRawManifestStub != nil {
		return fake.FullRawManifestStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.fullRawManifestReturns
	return fakeReturns.result1
}

func (fake *FakeManifestParser) FullRawManifestCallCount() int {
	fake.fullRawManifestMutex.RLock()
	defer fake.fullRawManifestMutex.RUnlock()
	return len(fake.fullRawManifestArgsForCall)
}

func (fake *FakeManifestParser) FullRawManifestCalls(stub func() []byte) {
	fake.fullRawManifestMutex.Lock()
	defer fake.fullRawManifestMutex.Unlock()
	fake.FullRawManifestStub = stub
}

func (fake *FakeManifestParser) FullRawManifestReturns(result1 []byte) {
	fake.fullRawManifestMutex.Lock()
	defer fake.fullRawManifestMutex.Unlock()
	fake.FullRawManifestStub = nil
	fake.fullRawManifestReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeManifestParser) FullRawManifestReturnsOnCall(i int, result1 []byte) {
	fake.fullRawManifestMutex.Lock()
	defer fake.fullRawManifestMutex.Unlock()
	fake.FullRawManifestStub = nil
	if fake.fullRawManifestReturnsOnCall == nil {
		fake.fullRawManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.fullRawManifestReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeManifestParser) RawAppManifest(arg1 string) ([]byte, error) {
	fake.rawAppManifestMutex.Lock()
	ret, specificReturn := fake.rawAppManifestReturnsOnCall[len(fake.rawAppManifestArgsForCall)]
	fake.rawAppManifestArgsForCall = append(fake.rawAppManifestArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RawAppManifest", []interface{}{arg1})
	fake.rawAppManifestMutex.Unlock()
	if fake.RawAppManifestStub != nil {
		return fake.RawAppManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.rawAppManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifestParser) RawAppManifestCallCount() int {
	fake.rawAppManifestMutex.RLock()
	defer fake.rawAppManifestMutex.RUnlock()
	return len(fake.rawAppManifestArgsForCall)
}

func (fake *FakeManifestParser) RawAppManifestCalls(stub func(string) ([]byte, error)) {
	fake.rawAppManifestMutex.Lock()
	defer fake.rawAppManifestMutex.Unlock()
	fake.RawAppManifestStub = stub
}

func (fake *FakeManifestParser) RawAppManifestArgsForCall(i int) string {
	fake.rawAppManifestMutex.RLock()
	defer fake.rawAppManifestMutex.RUnlock()
	argsForCall := fake.rawAppManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifestParser) RawAppManifestReturns(result1 []byte, result2 error) {
	fake.rawAppManifestMutex.Lock()
	defer fake.rawAppManifestMutex.Unlock()
	fake.RawAppManifestStub = nil
	fake.rawAppManifestReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) RawAppManifestReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.rawAppManifestMutex.Lock()
	defer fake.rawAppManifestMutex.Unlock()
	fake.RawAppManifestStub = nil
	if fake.rawAppManifestReturnsOnCall == nil {
		fake.rawAppManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.rawAppManifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appsMutex.RLock()
	defer fake.appsMutex.RUnlock()
	fake.containsManifestMutex.RLock()
	defer fake.containsManifestMutex.RUnlock()
	fake.fullRawManifestMutex.RLock()
	defer fake.fullRawManifestMutex.RUnlock()
	fake.rawAppManifestMutex.RLock()
	defer fake.rawAppManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifestParser) recordInvocation(key string, args []interface{}) {
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

var _ v7pushaction.ManifestParser = new(FakeManifestParser)
