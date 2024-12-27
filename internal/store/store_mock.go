// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package store

import (
	"github.com/bhavanki/rewind/pkg/model"
	"sync"
)

// Ensure, that StoreMock does implement Store.
// If this is not the case, regenerate this file with moq.
var _ Store = &StoreMock{}

// StoreMock is a mock implementation of Store.
//
//	func TestSomethingThatUsesStore(t *testing.T) {
//
//		// make and configure a mocked Store
//		mockedStore := &StoreMock{
//			CreateAPIFunc: func(a model.API) (model.API, error) {
//				panic("mock out the CreateAPI method")
//			},
//			CreateComponentFunc: func(c model.Component) (model.Component, error) {
//				panic("mock out the CreateComponent method")
//			},
//			CreateGroupFunc: func(g model.Group) (model.Group, error) {
//				panic("mock out the CreateGroup method")
//			},
//			CreateUserFunc: func(u model.User) (model.User, error) {
//				panic("mock out the CreateUser method")
//			},
//			DeleteAPIFunc: func(ref model.EntityRef) (model.API, error) {
//				panic("mock out the DeleteAPI method")
//			},
//			DeleteComponentFunc: func(ref model.EntityRef) (model.Component, error) {
//				panic("mock out the DeleteComponent method")
//			},
//			DeleteGroupFunc: func(ref model.EntityRef) (model.Group, error) {
//				panic("mock out the DeleteGroup method")
//			},
//			DeleteUserFunc: func(ref model.EntityRef) (model.User, error) {
//				panic("mock out the DeleteUser method")
//			},
//			ReadAPIFunc: func(ref model.EntityRef) (model.API, error) {
//				panic("mock out the ReadAPI method")
//			},
//			ReadComponentFunc: func(ref model.EntityRef) (model.Component, error) {
//				panic("mock out the ReadComponent method")
//			},
//			ReadGroupFunc: func(ref model.EntityRef) (model.Group, error) {
//				panic("mock out the ReadGroup method")
//			},
//			ReadUserFunc: func(ref model.EntityRef) (model.User, error) {
//				panic("mock out the ReadUser method")
//			},
//			UpdateAPIFunc: func(a model.API) (model.API, error) {
//				panic("mock out the UpdateAPI method")
//			},
//			UpdateComponentFunc: func(c model.Component) (model.Component, error) {
//				panic("mock out the UpdateComponent method")
//			},
//			UpdateGroupFunc: func(g model.Group) (model.Group, error) {
//				panic("mock out the UpdateGroup method")
//			},
//			UpdateUserFunc: func(u model.User) (model.User, error) {
//				panic("mock out the UpdateUser method")
//			},
//		}
//
//		// use mockedStore in code that requires Store
//		// and then make assertions.
//
//	}
type StoreMock struct {
	// CreateAPIFunc mocks the CreateAPI method.
	CreateAPIFunc func(a model.API) (model.API, error)

	// CreateComponentFunc mocks the CreateComponent method.
	CreateComponentFunc func(c model.Component) (model.Component, error)

	// CreateGroupFunc mocks the CreateGroup method.
	CreateGroupFunc func(g model.Group) (model.Group, error)

	// CreateUserFunc mocks the CreateUser method.
	CreateUserFunc func(u model.User) (model.User, error)

	// DeleteAPIFunc mocks the DeleteAPI method.
	DeleteAPIFunc func(ref model.EntityRef) (model.API, error)

	// DeleteComponentFunc mocks the DeleteComponent method.
	DeleteComponentFunc func(ref model.EntityRef) (model.Component, error)

	// DeleteGroupFunc mocks the DeleteGroup method.
	DeleteGroupFunc func(ref model.EntityRef) (model.Group, error)

	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(ref model.EntityRef) (model.User, error)

	// ReadAPIFunc mocks the ReadAPI method.
	ReadAPIFunc func(ref model.EntityRef) (model.API, error)

	// ReadComponentFunc mocks the ReadComponent method.
	ReadComponentFunc func(ref model.EntityRef) (model.Component, error)

	// ReadGroupFunc mocks the ReadGroup method.
	ReadGroupFunc func(ref model.EntityRef) (model.Group, error)

	// ReadUserFunc mocks the ReadUser method.
	ReadUserFunc func(ref model.EntityRef) (model.User, error)

	// UpdateAPIFunc mocks the UpdateAPI method.
	UpdateAPIFunc func(a model.API) (model.API, error)

	// UpdateComponentFunc mocks the UpdateComponent method.
	UpdateComponentFunc func(c model.Component) (model.Component, error)

	// UpdateGroupFunc mocks the UpdateGroup method.
	UpdateGroupFunc func(g model.Group) (model.Group, error)

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(u model.User) (model.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateAPI holds details about calls to the CreateAPI method.
		CreateAPI []struct {
			// A is the a argument value.
			A model.API
		}
		// CreateComponent holds details about calls to the CreateComponent method.
		CreateComponent []struct {
			// C is the c argument value.
			C model.Component
		}
		// CreateGroup holds details about calls to the CreateGroup method.
		CreateGroup []struct {
			// G is the g argument value.
			G model.Group
		}
		// CreateUser holds details about calls to the CreateUser method.
		CreateUser []struct {
			// U is the u argument value.
			U model.User
		}
		// DeleteAPI holds details about calls to the DeleteAPI method.
		DeleteAPI []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// DeleteComponent holds details about calls to the DeleteComponent method.
		DeleteComponent []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// DeleteGroup holds details about calls to the DeleteGroup method.
		DeleteGroup []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// ReadAPI holds details about calls to the ReadAPI method.
		ReadAPI []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// ReadComponent holds details about calls to the ReadComponent method.
		ReadComponent []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// ReadGroup holds details about calls to the ReadGroup method.
		ReadGroup []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// ReadUser holds details about calls to the ReadUser method.
		ReadUser []struct {
			// Ref is the ref argument value.
			Ref model.EntityRef
		}
		// UpdateAPI holds details about calls to the UpdateAPI method.
		UpdateAPI []struct {
			// A is the a argument value.
			A model.API
		}
		// UpdateComponent holds details about calls to the UpdateComponent method.
		UpdateComponent []struct {
			// C is the c argument value.
			C model.Component
		}
		// UpdateGroup holds details about calls to the UpdateGroup method.
		UpdateGroup []struct {
			// G is the g argument value.
			G model.Group
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// U is the u argument value.
			U model.User
		}
	}
	lockCreateAPI       sync.RWMutex
	lockCreateComponent sync.RWMutex
	lockCreateGroup     sync.RWMutex
	lockCreateUser      sync.RWMutex
	lockDeleteAPI       sync.RWMutex
	lockDeleteComponent sync.RWMutex
	lockDeleteGroup     sync.RWMutex
	lockDeleteUser      sync.RWMutex
	lockReadAPI         sync.RWMutex
	lockReadComponent   sync.RWMutex
	lockReadGroup       sync.RWMutex
	lockReadUser        sync.RWMutex
	lockUpdateAPI       sync.RWMutex
	lockUpdateComponent sync.RWMutex
	lockUpdateGroup     sync.RWMutex
	lockUpdateUser      sync.RWMutex
}

// CreateAPI calls CreateAPIFunc.
func (mock *StoreMock) CreateAPI(a model.API) (model.API, error) {
	if mock.CreateAPIFunc == nil {
		panic("StoreMock.CreateAPIFunc: method is nil but Store.CreateAPI was just called")
	}
	callInfo := struct {
		A model.API
	}{
		A: a,
	}
	mock.lockCreateAPI.Lock()
	mock.calls.CreateAPI = append(mock.calls.CreateAPI, callInfo)
	mock.lockCreateAPI.Unlock()
	return mock.CreateAPIFunc(a)
}

// CreateAPICalls gets all the calls that were made to CreateAPI.
// Check the length with:
//
//	len(mockedStore.CreateAPICalls())
func (mock *StoreMock) CreateAPICalls() []struct {
	A model.API
} {
	var calls []struct {
		A model.API
	}
	mock.lockCreateAPI.RLock()
	calls = mock.calls.CreateAPI
	mock.lockCreateAPI.RUnlock()
	return calls
}

// CreateComponent calls CreateComponentFunc.
func (mock *StoreMock) CreateComponent(c model.Component) (model.Component, error) {
	if mock.CreateComponentFunc == nil {
		panic("StoreMock.CreateComponentFunc: method is nil but Store.CreateComponent was just called")
	}
	callInfo := struct {
		C model.Component
	}{
		C: c,
	}
	mock.lockCreateComponent.Lock()
	mock.calls.CreateComponent = append(mock.calls.CreateComponent, callInfo)
	mock.lockCreateComponent.Unlock()
	return mock.CreateComponentFunc(c)
}

// CreateComponentCalls gets all the calls that were made to CreateComponent.
// Check the length with:
//
//	len(mockedStore.CreateComponentCalls())
func (mock *StoreMock) CreateComponentCalls() []struct {
	C model.Component
} {
	var calls []struct {
		C model.Component
	}
	mock.lockCreateComponent.RLock()
	calls = mock.calls.CreateComponent
	mock.lockCreateComponent.RUnlock()
	return calls
}

// CreateGroup calls CreateGroupFunc.
func (mock *StoreMock) CreateGroup(g model.Group) (model.Group, error) {
	if mock.CreateGroupFunc == nil {
		panic("StoreMock.CreateGroupFunc: method is nil but Store.CreateGroup was just called")
	}
	callInfo := struct {
		G model.Group
	}{
		G: g,
	}
	mock.lockCreateGroup.Lock()
	mock.calls.CreateGroup = append(mock.calls.CreateGroup, callInfo)
	mock.lockCreateGroup.Unlock()
	return mock.CreateGroupFunc(g)
}

// CreateGroupCalls gets all the calls that were made to CreateGroup.
// Check the length with:
//
//	len(mockedStore.CreateGroupCalls())
func (mock *StoreMock) CreateGroupCalls() []struct {
	G model.Group
} {
	var calls []struct {
		G model.Group
	}
	mock.lockCreateGroup.RLock()
	calls = mock.calls.CreateGroup
	mock.lockCreateGroup.RUnlock()
	return calls
}

// CreateUser calls CreateUserFunc.
func (mock *StoreMock) CreateUser(u model.User) (model.User, error) {
	if mock.CreateUserFunc == nil {
		panic("StoreMock.CreateUserFunc: method is nil but Store.CreateUser was just called")
	}
	callInfo := struct {
		U model.User
	}{
		U: u,
	}
	mock.lockCreateUser.Lock()
	mock.calls.CreateUser = append(mock.calls.CreateUser, callInfo)
	mock.lockCreateUser.Unlock()
	return mock.CreateUserFunc(u)
}

// CreateUserCalls gets all the calls that were made to CreateUser.
// Check the length with:
//
//	len(mockedStore.CreateUserCalls())
func (mock *StoreMock) CreateUserCalls() []struct {
	U model.User
} {
	var calls []struct {
		U model.User
	}
	mock.lockCreateUser.RLock()
	calls = mock.calls.CreateUser
	mock.lockCreateUser.RUnlock()
	return calls
}

// DeleteAPI calls DeleteAPIFunc.
func (mock *StoreMock) DeleteAPI(ref model.EntityRef) (model.API, error) {
	if mock.DeleteAPIFunc == nil {
		panic("StoreMock.DeleteAPIFunc: method is nil but Store.DeleteAPI was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockDeleteAPI.Lock()
	mock.calls.DeleteAPI = append(mock.calls.DeleteAPI, callInfo)
	mock.lockDeleteAPI.Unlock()
	return mock.DeleteAPIFunc(ref)
}

// DeleteAPICalls gets all the calls that were made to DeleteAPI.
// Check the length with:
//
//	len(mockedStore.DeleteAPICalls())
func (mock *StoreMock) DeleteAPICalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockDeleteAPI.RLock()
	calls = mock.calls.DeleteAPI
	mock.lockDeleteAPI.RUnlock()
	return calls
}

// DeleteComponent calls DeleteComponentFunc.
func (mock *StoreMock) DeleteComponent(ref model.EntityRef) (model.Component, error) {
	if mock.DeleteComponentFunc == nil {
		panic("StoreMock.DeleteComponentFunc: method is nil but Store.DeleteComponent was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockDeleteComponent.Lock()
	mock.calls.DeleteComponent = append(mock.calls.DeleteComponent, callInfo)
	mock.lockDeleteComponent.Unlock()
	return mock.DeleteComponentFunc(ref)
}

// DeleteComponentCalls gets all the calls that were made to DeleteComponent.
// Check the length with:
//
//	len(mockedStore.DeleteComponentCalls())
func (mock *StoreMock) DeleteComponentCalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockDeleteComponent.RLock()
	calls = mock.calls.DeleteComponent
	mock.lockDeleteComponent.RUnlock()
	return calls
}

// DeleteGroup calls DeleteGroupFunc.
func (mock *StoreMock) DeleteGroup(ref model.EntityRef) (model.Group, error) {
	if mock.DeleteGroupFunc == nil {
		panic("StoreMock.DeleteGroupFunc: method is nil but Store.DeleteGroup was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockDeleteGroup.Lock()
	mock.calls.DeleteGroup = append(mock.calls.DeleteGroup, callInfo)
	mock.lockDeleteGroup.Unlock()
	return mock.DeleteGroupFunc(ref)
}

// DeleteGroupCalls gets all the calls that were made to DeleteGroup.
// Check the length with:
//
//	len(mockedStore.DeleteGroupCalls())
func (mock *StoreMock) DeleteGroupCalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockDeleteGroup.RLock()
	calls = mock.calls.DeleteGroup
	mock.lockDeleteGroup.RUnlock()
	return calls
}

// DeleteUser calls DeleteUserFunc.
func (mock *StoreMock) DeleteUser(ref model.EntityRef) (model.User, error) {
	if mock.DeleteUserFunc == nil {
		panic("StoreMock.DeleteUserFunc: method is nil but Store.DeleteUser was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	mock.lockDeleteUser.Unlock()
	return mock.DeleteUserFunc(ref)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//
//	len(mockedStore.DeleteUserCalls())
func (mock *StoreMock) DeleteUserCalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	mock.lockDeleteUser.RUnlock()
	return calls
}

// ReadAPI calls ReadAPIFunc.
func (mock *StoreMock) ReadAPI(ref model.EntityRef) (model.API, error) {
	if mock.ReadAPIFunc == nil {
		panic("StoreMock.ReadAPIFunc: method is nil but Store.ReadAPI was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockReadAPI.Lock()
	mock.calls.ReadAPI = append(mock.calls.ReadAPI, callInfo)
	mock.lockReadAPI.Unlock()
	return mock.ReadAPIFunc(ref)
}

// ReadAPICalls gets all the calls that were made to ReadAPI.
// Check the length with:
//
//	len(mockedStore.ReadAPICalls())
func (mock *StoreMock) ReadAPICalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockReadAPI.RLock()
	calls = mock.calls.ReadAPI
	mock.lockReadAPI.RUnlock()
	return calls
}

// ReadComponent calls ReadComponentFunc.
func (mock *StoreMock) ReadComponent(ref model.EntityRef) (model.Component, error) {
	if mock.ReadComponentFunc == nil {
		panic("StoreMock.ReadComponentFunc: method is nil but Store.ReadComponent was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockReadComponent.Lock()
	mock.calls.ReadComponent = append(mock.calls.ReadComponent, callInfo)
	mock.lockReadComponent.Unlock()
	return mock.ReadComponentFunc(ref)
}

// ReadComponentCalls gets all the calls that were made to ReadComponent.
// Check the length with:
//
//	len(mockedStore.ReadComponentCalls())
func (mock *StoreMock) ReadComponentCalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockReadComponent.RLock()
	calls = mock.calls.ReadComponent
	mock.lockReadComponent.RUnlock()
	return calls
}

// ReadGroup calls ReadGroupFunc.
func (mock *StoreMock) ReadGroup(ref model.EntityRef) (model.Group, error) {
	if mock.ReadGroupFunc == nil {
		panic("StoreMock.ReadGroupFunc: method is nil but Store.ReadGroup was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockReadGroup.Lock()
	mock.calls.ReadGroup = append(mock.calls.ReadGroup, callInfo)
	mock.lockReadGroup.Unlock()
	return mock.ReadGroupFunc(ref)
}

// ReadGroupCalls gets all the calls that were made to ReadGroup.
// Check the length with:
//
//	len(mockedStore.ReadGroupCalls())
func (mock *StoreMock) ReadGroupCalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockReadGroup.RLock()
	calls = mock.calls.ReadGroup
	mock.lockReadGroup.RUnlock()
	return calls
}

// ReadUser calls ReadUserFunc.
func (mock *StoreMock) ReadUser(ref model.EntityRef) (model.User, error) {
	if mock.ReadUserFunc == nil {
		panic("StoreMock.ReadUserFunc: method is nil but Store.ReadUser was just called")
	}
	callInfo := struct {
		Ref model.EntityRef
	}{
		Ref: ref,
	}
	mock.lockReadUser.Lock()
	mock.calls.ReadUser = append(mock.calls.ReadUser, callInfo)
	mock.lockReadUser.Unlock()
	return mock.ReadUserFunc(ref)
}

// ReadUserCalls gets all the calls that were made to ReadUser.
// Check the length with:
//
//	len(mockedStore.ReadUserCalls())
func (mock *StoreMock) ReadUserCalls() []struct {
	Ref model.EntityRef
} {
	var calls []struct {
		Ref model.EntityRef
	}
	mock.lockReadUser.RLock()
	calls = mock.calls.ReadUser
	mock.lockReadUser.RUnlock()
	return calls
}

// UpdateAPI calls UpdateAPIFunc.
func (mock *StoreMock) UpdateAPI(a model.API) (model.API, error) {
	if mock.UpdateAPIFunc == nil {
		panic("StoreMock.UpdateAPIFunc: method is nil but Store.UpdateAPI was just called")
	}
	callInfo := struct {
		A model.API
	}{
		A: a,
	}
	mock.lockUpdateAPI.Lock()
	mock.calls.UpdateAPI = append(mock.calls.UpdateAPI, callInfo)
	mock.lockUpdateAPI.Unlock()
	return mock.UpdateAPIFunc(a)
}

// UpdateAPICalls gets all the calls that were made to UpdateAPI.
// Check the length with:
//
//	len(mockedStore.UpdateAPICalls())
func (mock *StoreMock) UpdateAPICalls() []struct {
	A model.API
} {
	var calls []struct {
		A model.API
	}
	mock.lockUpdateAPI.RLock()
	calls = mock.calls.UpdateAPI
	mock.lockUpdateAPI.RUnlock()
	return calls
}

// UpdateComponent calls UpdateComponentFunc.
func (mock *StoreMock) UpdateComponent(c model.Component) (model.Component, error) {
	if mock.UpdateComponentFunc == nil {
		panic("StoreMock.UpdateComponentFunc: method is nil but Store.UpdateComponent was just called")
	}
	callInfo := struct {
		C model.Component
	}{
		C: c,
	}
	mock.lockUpdateComponent.Lock()
	mock.calls.UpdateComponent = append(mock.calls.UpdateComponent, callInfo)
	mock.lockUpdateComponent.Unlock()
	return mock.UpdateComponentFunc(c)
}

// UpdateComponentCalls gets all the calls that were made to UpdateComponent.
// Check the length with:
//
//	len(mockedStore.UpdateComponentCalls())
func (mock *StoreMock) UpdateComponentCalls() []struct {
	C model.Component
} {
	var calls []struct {
		C model.Component
	}
	mock.lockUpdateComponent.RLock()
	calls = mock.calls.UpdateComponent
	mock.lockUpdateComponent.RUnlock()
	return calls
}

// UpdateGroup calls UpdateGroupFunc.
func (mock *StoreMock) UpdateGroup(g model.Group) (model.Group, error) {
	if mock.UpdateGroupFunc == nil {
		panic("StoreMock.UpdateGroupFunc: method is nil but Store.UpdateGroup was just called")
	}
	callInfo := struct {
		G model.Group
	}{
		G: g,
	}
	mock.lockUpdateGroup.Lock()
	mock.calls.UpdateGroup = append(mock.calls.UpdateGroup, callInfo)
	mock.lockUpdateGroup.Unlock()
	return mock.UpdateGroupFunc(g)
}

// UpdateGroupCalls gets all the calls that were made to UpdateGroup.
// Check the length with:
//
//	len(mockedStore.UpdateGroupCalls())
func (mock *StoreMock) UpdateGroupCalls() []struct {
	G model.Group
} {
	var calls []struct {
		G model.Group
	}
	mock.lockUpdateGroup.RLock()
	calls = mock.calls.UpdateGroup
	mock.lockUpdateGroup.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *StoreMock) UpdateUser(u model.User) (model.User, error) {
	if mock.UpdateUserFunc == nil {
		panic("StoreMock.UpdateUserFunc: method is nil but Store.UpdateUser was just called")
	}
	callInfo := struct {
		U model.User
	}{
		U: u,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(u)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//
//	len(mockedStore.UpdateUserCalls())
func (mock *StoreMock) UpdateUserCalls() []struct {
	U model.User
} {
	var calls []struct {
		U model.User
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}
