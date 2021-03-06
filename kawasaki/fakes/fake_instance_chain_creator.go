// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/cloudfoundry-incubator/guardian/kawasaki"
	"github.com/pivotal-golang/lager"
)

type FakeInstanceChainCreator struct {
	CreateStub        func(logger lager.Logger, instanceChain, bridgeName string, ip net.IP, network *net.IPNet) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		logger        lager.Logger
		instanceChain string
		bridgeName    string
		ip            net.IP
		network       *net.IPNet
	}
	createReturns struct {
		result1 error
	}
	DestroyStub        func(logger lager.Logger, instanceChain string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		logger        lager.Logger
		instanceChain string
	}
	destroyReturns struct {
		result1 error
	}
}

func (fake *FakeInstanceChainCreator) Create(logger lager.Logger, instanceChain string, bridgeName string, ip net.IP, network *net.IPNet) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		logger        lager.Logger
		instanceChain string
		bridgeName    string
		ip            net.IP
		network       *net.IPNet
	}{logger, instanceChain, bridgeName, ip, network})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(logger, instanceChain, bridgeName, ip, network)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeInstanceChainCreator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeInstanceChainCreator) CreateArgsForCall(i int) (lager.Logger, string, string, net.IP, *net.IPNet) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].logger, fake.createArgsForCall[i].instanceChain, fake.createArgsForCall[i].bridgeName, fake.createArgsForCall[i].ip, fake.createArgsForCall[i].network
}

func (fake *FakeInstanceChainCreator) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstanceChainCreator) Destroy(logger lager.Logger, instanceChain string) error {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		logger        lager.Logger
		instanceChain string
	}{logger, instanceChain})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(logger, instanceChain)
	} else {
		return fake.destroyReturns.result1
	}
}

func (fake *FakeInstanceChainCreator) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeInstanceChainCreator) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].logger, fake.destroyArgsForCall[i].instanceChain
}

func (fake *FakeInstanceChainCreator) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

var _ kawasaki.InstanceChainCreator = new(FakeInstanceChainCreator)
