// This file was generated by counterfeiter
package fakes

import (
	"os"
	"sync"

	"github.com/cloudfoundry-incubator/guardian/kawasaki"
)

type FakeNetnsExecer struct {
	ExecStub        func(netnsFD *os.File, cb func() error) error
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		netnsFD *os.File
		cb      func() error
	}
	execReturns struct {
		result1 error
	}
}

func (fake *FakeNetnsExecer) Exec(netnsFD *os.File, cb func() error) error {
	fake.execMutex.Lock()
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		netnsFD *os.File
		cb      func() error
	}{netnsFD, cb})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(netnsFD, cb)
	} else {
		return fake.execReturns.result1
	}
}

func (fake *FakeNetnsExecer) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeNetnsExecer) ExecArgsForCall(i int) (*os.File, func() error) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].netnsFD, fake.execArgsForCall[i].cb
}

func (fake *FakeNetnsExecer) ExecReturns(result1 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 error
	}{result1}
}

var _ kawasaki.NetnsExecer = new(FakeNetnsExecer)
