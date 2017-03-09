// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeCreatedVolume struct {
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() dbng.VolumeType
	typeMutex       sync.RWMutex
	typeArgsForCall []struct{}
	typeReturns     struct {
		result1 dbng.VolumeType
	}
	typeReturnsOnCall map[int]struct {
		result1 dbng.VolumeType
	}
	CreateChildForContainerStub        func(dbng.CreatingContainer, string) (dbng.CreatingVolume, error)
	createChildForContainerMutex       sync.RWMutex
	createChildForContainerArgsForCall []struct {
		arg1 dbng.CreatingContainer
		arg2 string
	}
	createChildForContainerReturns struct {
		result1 dbng.CreatingVolume
		result2 error
	}
	createChildForContainerReturnsOnCall map[int]struct {
		result1 dbng.CreatingVolume
		result2 error
	}
	DestroyingStub        func() (dbng.DestroyingVolume, error)
	destroyingMutex       sync.RWMutex
	destroyingArgsForCall []struct{}
	destroyingReturns     struct {
		result1 dbng.DestroyingVolume
		result2 error
	}
	destroyingReturnsOnCall map[int]struct {
		result1 dbng.DestroyingVolume
		result2 error
	}
	WorkerStub        func() dbng.Worker
	workerMutex       sync.RWMutex
	workerArgsForCall []struct{}
	workerReturns     struct {
		result1 dbng.Worker
	}
	workerReturnsOnCall map[int]struct {
		result1 dbng.Worker
	}
	SizeInBytesStub        func() int64
	sizeInBytesMutex       sync.RWMutex
	sizeInBytesArgsForCall []struct{}
	sizeInBytesReturns     struct {
		result1 int64
	}
	sizeInBytesReturnsOnCall map[int]struct {
		result1 int64
	}
	InitializeStub        func() error
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct{}
	initializeReturns     struct {
		result1 error
	}
	initializeReturnsOnCall map[int]struct {
		result1 error
	}
	IsInitializedStub        func() (bool, error)
	isInitializedMutex       sync.RWMutex
	isInitializedArgsForCall []struct{}
	isInitializedReturns     struct {
		result1 bool
		result2 error
	}
	isInitializedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ContainerHandleStub        func() string
	containerHandleMutex       sync.RWMutex
	containerHandleArgsForCall []struct{}
	containerHandleReturns     struct {
		result1 string
	}
	containerHandleReturnsOnCall map[int]struct {
		result1 string
	}
	ParentHandleStub        func() string
	parentHandleMutex       sync.RWMutex
	parentHandleArgsForCall []struct{}
	parentHandleReturns     struct {
		result1 string
	}
	parentHandleReturnsOnCall map[int]struct {
		result1 string
	}
	ResourceTypeStub        func() (*dbng.VolumeResourceType, error)
	resourceTypeMutex       sync.RWMutex
	resourceTypeArgsForCall []struct{}
	resourceTypeReturns     struct {
		result1 *dbng.VolumeResourceType
		result2 error
	}
	resourceTypeReturnsOnCall map[int]struct {
		result1 *dbng.VolumeResourceType
		result2 error
	}
	BaseResourceTypeStub        func() (*dbng.UsedWorkerBaseResourceType, error)
	baseResourceTypeMutex       sync.RWMutex
	baseResourceTypeArgsForCall []struct{}
	baseResourceTypeReturns     struct {
		result1 *dbng.UsedWorkerBaseResourceType
		result2 error
	}
	baseResourceTypeReturnsOnCall map[int]struct {
		result1 *dbng.UsedWorkerBaseResourceType
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreatedVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.handleReturns.result1
}

func (fake *FakeCreatedVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeCreatedVolume) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pathReturns.result1
}

func (fake *FakeCreatedVolume) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeCreatedVolume) PathReturns(result1 string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) PathReturnsOnCall(i int, result1 string) {
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) Type() dbng.VolumeType {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.typeReturns.result1
}

func (fake *FakeCreatedVolume) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeCreatedVolume) TypeReturns(result1 dbng.VolumeType) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 dbng.VolumeType
	}{result1}
}

func (fake *FakeCreatedVolume) TypeReturnsOnCall(i int, result1 dbng.VolumeType) {
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 dbng.VolumeType
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 dbng.VolumeType
	}{result1}
}

func (fake *FakeCreatedVolume) CreateChildForContainer(arg1 dbng.CreatingContainer, arg2 string) (dbng.CreatingVolume, error) {
	fake.createChildForContainerMutex.Lock()
	ret, specificReturn := fake.createChildForContainerReturnsOnCall[len(fake.createChildForContainerArgsForCall)]
	fake.createChildForContainerArgsForCall = append(fake.createChildForContainerArgsForCall, struct {
		arg1 dbng.CreatingContainer
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateChildForContainer", []interface{}{arg1, arg2})
	fake.createChildForContainerMutex.Unlock()
	if fake.CreateChildForContainerStub != nil {
		return fake.CreateChildForContainerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createChildForContainerReturns.result1, fake.createChildForContainerReturns.result2
}

func (fake *FakeCreatedVolume) CreateChildForContainerCallCount() int {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	return len(fake.createChildForContainerArgsForCall)
}

func (fake *FakeCreatedVolume) CreateChildForContainerArgsForCall(i int) (dbng.CreatingContainer, string) {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	return fake.createChildForContainerArgsForCall[i].arg1, fake.createChildForContainerArgsForCall[i].arg2
}

func (fake *FakeCreatedVolume) CreateChildForContainerReturns(result1 dbng.CreatingVolume, result2 error) {
	fake.CreateChildForContainerStub = nil
	fake.createChildForContainerReturns = struct {
		result1 dbng.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) CreateChildForContainerReturnsOnCall(i int, result1 dbng.CreatingVolume, result2 error) {
	fake.CreateChildForContainerStub = nil
	if fake.createChildForContainerReturnsOnCall == nil {
		fake.createChildForContainerReturnsOnCall = make(map[int]struct {
			result1 dbng.CreatingVolume
			result2 error
		})
	}
	fake.createChildForContainerReturnsOnCall[i] = struct {
		result1 dbng.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) Destroying() (dbng.DestroyingVolume, error) {
	fake.destroyingMutex.Lock()
	ret, specificReturn := fake.destroyingReturnsOnCall[len(fake.destroyingArgsForCall)]
	fake.destroyingArgsForCall = append(fake.destroyingArgsForCall, struct{}{})
	fake.recordInvocation("Destroying", []interface{}{})
	fake.destroyingMutex.Unlock()
	if fake.DestroyingStub != nil {
		return fake.DestroyingStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.destroyingReturns.result1, fake.destroyingReturns.result2
}

func (fake *FakeCreatedVolume) DestroyingCallCount() int {
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	return len(fake.destroyingArgsForCall)
}

func (fake *FakeCreatedVolume) DestroyingReturns(result1 dbng.DestroyingVolume, result2 error) {
	fake.DestroyingStub = nil
	fake.destroyingReturns = struct {
		result1 dbng.DestroyingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) DestroyingReturnsOnCall(i int, result1 dbng.DestroyingVolume, result2 error) {
	fake.DestroyingStub = nil
	if fake.destroyingReturnsOnCall == nil {
		fake.destroyingReturnsOnCall = make(map[int]struct {
			result1 dbng.DestroyingVolume
			result2 error
		})
	}
	fake.destroyingReturnsOnCall[i] = struct {
		result1 dbng.DestroyingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) Worker() dbng.Worker {
	fake.workerMutex.Lock()
	ret, specificReturn := fake.workerReturnsOnCall[len(fake.workerArgsForCall)]
	fake.workerArgsForCall = append(fake.workerArgsForCall, struct{}{})
	fake.recordInvocation("Worker", []interface{}{})
	fake.workerMutex.Unlock()
	if fake.WorkerStub != nil {
		return fake.WorkerStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.workerReturns.result1
}

func (fake *FakeCreatedVolume) WorkerCallCount() int {
	fake.workerMutex.RLock()
	defer fake.workerMutex.RUnlock()
	return len(fake.workerArgsForCall)
}

func (fake *FakeCreatedVolume) WorkerReturns(result1 dbng.Worker) {
	fake.WorkerStub = nil
	fake.workerReturns = struct {
		result1 dbng.Worker
	}{result1}
}

func (fake *FakeCreatedVolume) WorkerReturnsOnCall(i int, result1 dbng.Worker) {
	fake.WorkerStub = nil
	if fake.workerReturnsOnCall == nil {
		fake.workerReturnsOnCall = make(map[int]struct {
			result1 dbng.Worker
		})
	}
	fake.workerReturnsOnCall[i] = struct {
		result1 dbng.Worker
	}{result1}
}

func (fake *FakeCreatedVolume) SizeInBytes() int64 {
	fake.sizeInBytesMutex.Lock()
	ret, specificReturn := fake.sizeInBytesReturnsOnCall[len(fake.sizeInBytesArgsForCall)]
	fake.sizeInBytesArgsForCall = append(fake.sizeInBytesArgsForCall, struct{}{})
	fake.recordInvocation("SizeInBytes", []interface{}{})
	fake.sizeInBytesMutex.Unlock()
	if fake.SizeInBytesStub != nil {
		return fake.SizeInBytesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sizeInBytesReturns.result1
}

func (fake *FakeCreatedVolume) SizeInBytesCallCount() int {
	fake.sizeInBytesMutex.RLock()
	defer fake.sizeInBytesMutex.RUnlock()
	return len(fake.sizeInBytesArgsForCall)
}

func (fake *FakeCreatedVolume) SizeInBytesReturns(result1 int64) {
	fake.SizeInBytesStub = nil
	fake.sizeInBytesReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeCreatedVolume) SizeInBytesReturnsOnCall(i int, result1 int64) {
	fake.SizeInBytesStub = nil
	if fake.sizeInBytesReturnsOnCall == nil {
		fake.sizeInBytesReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.sizeInBytesReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeCreatedVolume) Initialize() error {
	fake.initializeMutex.Lock()
	ret, specificReturn := fake.initializeReturnsOnCall[len(fake.initializeArgsForCall)]
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct{}{})
	fake.recordInvocation("Initialize", []interface{}{})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		return fake.InitializeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.initializeReturns.result1
}

func (fake *FakeCreatedVolume) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *FakeCreatedVolume) InitializeReturns(result1 error) {
	fake.InitializeStub = nil
	fake.initializeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedVolume) InitializeReturnsOnCall(i int, result1 error) {
	fake.InitializeStub = nil
	if fake.initializeReturnsOnCall == nil {
		fake.initializeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedVolume) IsInitialized() (bool, error) {
	fake.isInitializedMutex.Lock()
	ret, specificReturn := fake.isInitializedReturnsOnCall[len(fake.isInitializedArgsForCall)]
	fake.isInitializedArgsForCall = append(fake.isInitializedArgsForCall, struct{}{})
	fake.recordInvocation("IsInitialized", []interface{}{})
	fake.isInitializedMutex.Unlock()
	if fake.IsInitializedStub != nil {
		return fake.IsInitializedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isInitializedReturns.result1, fake.isInitializedReturns.result2
}

func (fake *FakeCreatedVolume) IsInitializedCallCount() int {
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	return len(fake.isInitializedArgsForCall)
}

func (fake *FakeCreatedVolume) IsInitializedReturns(result1 bool, result2 error) {
	fake.IsInitializedStub = nil
	fake.isInitializedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) IsInitializedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsInitializedStub = nil
	if fake.isInitializedReturnsOnCall == nil {
		fake.isInitializedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isInitializedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) ContainerHandle() string {
	fake.containerHandleMutex.Lock()
	ret, specificReturn := fake.containerHandleReturnsOnCall[len(fake.containerHandleArgsForCall)]
	fake.containerHandleArgsForCall = append(fake.containerHandleArgsForCall, struct{}{})
	fake.recordInvocation("ContainerHandle", []interface{}{})
	fake.containerHandleMutex.Unlock()
	if fake.ContainerHandleStub != nil {
		return fake.ContainerHandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.containerHandleReturns.result1
}

func (fake *FakeCreatedVolume) ContainerHandleCallCount() int {
	fake.containerHandleMutex.RLock()
	defer fake.containerHandleMutex.RUnlock()
	return len(fake.containerHandleArgsForCall)
}

func (fake *FakeCreatedVolume) ContainerHandleReturns(result1 string) {
	fake.ContainerHandleStub = nil
	fake.containerHandleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) ContainerHandleReturnsOnCall(i int, result1 string) {
	fake.ContainerHandleStub = nil
	if fake.containerHandleReturnsOnCall == nil {
		fake.containerHandleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.containerHandleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) ParentHandle() string {
	fake.parentHandleMutex.Lock()
	ret, specificReturn := fake.parentHandleReturnsOnCall[len(fake.parentHandleArgsForCall)]
	fake.parentHandleArgsForCall = append(fake.parentHandleArgsForCall, struct{}{})
	fake.recordInvocation("ParentHandle", []interface{}{})
	fake.parentHandleMutex.Unlock()
	if fake.ParentHandleStub != nil {
		return fake.ParentHandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.parentHandleReturns.result1
}

func (fake *FakeCreatedVolume) ParentHandleCallCount() int {
	fake.parentHandleMutex.RLock()
	defer fake.parentHandleMutex.RUnlock()
	return len(fake.parentHandleArgsForCall)
}

func (fake *FakeCreatedVolume) ParentHandleReturns(result1 string) {
	fake.ParentHandleStub = nil
	fake.parentHandleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) ParentHandleReturnsOnCall(i int, result1 string) {
	fake.ParentHandleStub = nil
	if fake.parentHandleReturnsOnCall == nil {
		fake.parentHandleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.parentHandleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedVolume) ResourceType() (*dbng.VolumeResourceType, error) {
	fake.resourceTypeMutex.Lock()
	ret, specificReturn := fake.resourceTypeReturnsOnCall[len(fake.resourceTypeArgsForCall)]
	fake.resourceTypeArgsForCall = append(fake.resourceTypeArgsForCall, struct{}{})
	fake.recordInvocation("ResourceType", []interface{}{})
	fake.resourceTypeMutex.Unlock()
	if fake.ResourceTypeStub != nil {
		return fake.ResourceTypeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resourceTypeReturns.result1, fake.resourceTypeReturns.result2
}

func (fake *FakeCreatedVolume) ResourceTypeCallCount() int {
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	return len(fake.resourceTypeArgsForCall)
}

func (fake *FakeCreatedVolume) ResourceTypeReturns(result1 *dbng.VolumeResourceType, result2 error) {
	fake.ResourceTypeStub = nil
	fake.resourceTypeReturns = struct {
		result1 *dbng.VolumeResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) ResourceTypeReturnsOnCall(i int, result1 *dbng.VolumeResourceType, result2 error) {
	fake.ResourceTypeStub = nil
	if fake.resourceTypeReturnsOnCall == nil {
		fake.resourceTypeReturnsOnCall = make(map[int]struct {
			result1 *dbng.VolumeResourceType
			result2 error
		})
	}
	fake.resourceTypeReturnsOnCall[i] = struct {
		result1 *dbng.VolumeResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) BaseResourceType() (*dbng.UsedWorkerBaseResourceType, error) {
	fake.baseResourceTypeMutex.Lock()
	ret, specificReturn := fake.baseResourceTypeReturnsOnCall[len(fake.baseResourceTypeArgsForCall)]
	fake.baseResourceTypeArgsForCall = append(fake.baseResourceTypeArgsForCall, struct{}{})
	fake.recordInvocation("BaseResourceType", []interface{}{})
	fake.baseResourceTypeMutex.Unlock()
	if fake.BaseResourceTypeStub != nil {
		return fake.BaseResourceTypeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.baseResourceTypeReturns.result1, fake.baseResourceTypeReturns.result2
}

func (fake *FakeCreatedVolume) BaseResourceTypeCallCount() int {
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	return len(fake.baseResourceTypeArgsForCall)
}

func (fake *FakeCreatedVolume) BaseResourceTypeReturns(result1 *dbng.UsedWorkerBaseResourceType, result2 error) {
	fake.BaseResourceTypeStub = nil
	fake.baseResourceTypeReturns = struct {
		result1 *dbng.UsedWorkerBaseResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) BaseResourceTypeReturnsOnCall(i int, result1 *dbng.UsedWorkerBaseResourceType, result2 error) {
	fake.BaseResourceTypeStub = nil
	if fake.baseResourceTypeReturnsOnCall == nil {
		fake.baseResourceTypeReturnsOnCall = make(map[int]struct {
			result1 *dbng.UsedWorkerBaseResourceType
			result2 error
		})
	}
	fake.baseResourceTypeReturnsOnCall[i] = struct {
		result1 *dbng.UsedWorkerBaseResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	fake.workerMutex.RLock()
	defer fake.workerMutex.RUnlock()
	fake.sizeInBytesMutex.RLock()
	defer fake.sizeInBytesMutex.RUnlock()
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	fake.containerHandleMutex.RLock()
	defer fake.containerHandleMutex.RUnlock()
	fake.parentHandleMutex.RLock()
	defer fake.parentHandleMutex.RUnlock()
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCreatedVolume) recordInvocation(key string, args []interface{}) {
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

var _ dbng.CreatedVolume = new(FakeCreatedVolume)
