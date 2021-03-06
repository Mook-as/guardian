// This file was generated by counterfeiter
package fakes

import (
	"os"
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc/bundlerules"
)

type FakeMkdirChowner struct {
	MkdirChownStub        func(path string, perms os.FileMode, uid, gid int) error
	mkdirChownMutex       sync.RWMutex
	mkdirChownArgsForCall []struct {
		path  string
		perms os.FileMode
		uid   int
		gid   int
	}
	mkdirChownReturns struct {
		result1 error
	}
}

func (fake *FakeMkdirChowner) MkdirChown(path string, perms os.FileMode, uid int, gid int) error {
	fake.mkdirChownMutex.Lock()
	fake.mkdirChownArgsForCall = append(fake.mkdirChownArgsForCall, struct {
		path  string
		perms os.FileMode
		uid   int
		gid   int
	}{path, perms, uid, gid})
	fake.mkdirChownMutex.Unlock()
	if fake.MkdirChownStub != nil {
		return fake.MkdirChownStub(path, perms, uid, gid)
	} else {
		return fake.mkdirChownReturns.result1
	}
}

func (fake *FakeMkdirChowner) MkdirChownCallCount() int {
	fake.mkdirChownMutex.RLock()
	defer fake.mkdirChownMutex.RUnlock()
	return len(fake.mkdirChownArgsForCall)
}

func (fake *FakeMkdirChowner) MkdirChownArgsForCall(i int) (string, os.FileMode, int, int) {
	fake.mkdirChownMutex.RLock()
	defer fake.mkdirChownMutex.RUnlock()
	return fake.mkdirChownArgsForCall[i].path, fake.mkdirChownArgsForCall[i].perms, fake.mkdirChownArgsForCall[i].uid, fake.mkdirChownArgsForCall[i].gid
}

func (fake *FakeMkdirChowner) MkdirChownReturns(result1 error) {
	fake.MkdirChownStub = nil
	fake.mkdirChownReturns = struct {
		result1 error
	}{result1}
}

var _ bundlerules.MkdirChowner = new(FakeMkdirChowner)
