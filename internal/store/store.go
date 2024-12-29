package store

import (
	"github.com/bhavanki/rewind/pkg/model"
)

//go:generate moq -out store_mock.go . Store

type Store interface {
	CreateComponent(c model.Component) (model.Component, error)
	ReadComponent(ref model.EntityRef) (model.Component, error)
	UpdateComponent(c model.Component) (model.Component, error)
	DeleteComponent(ref model.EntityRef) (model.Component, error)
	ListComponents(filters []Filter, ordering Ordering, pagination Pagination) ([]model.EntityRef, Pagination, error)

	CreateAPI(a model.API) (model.API, error)
	ReadAPI(ref model.EntityRef) (model.API, error)
	UpdateAPI(a model.API) (model.API, error)
	DeleteAPI(ref model.EntityRef) (model.API, error)

	CreateUser(u model.User) (model.User, error)
	ReadUser(ref model.EntityRef) (model.User, error)
	UpdateUser(u model.User) (model.User, error)
	DeleteUser(ref model.EntityRef) (model.User, error)

	CreateGroup(g model.Group) (model.Group, error)
	ReadGroup(ref model.EntityRef) (model.Group, error)
	UpdateGroup(g model.Group) (model.Group, error)
	DeleteGroup(ref model.EntityRef) (model.Group, error)
}

type Filter struct {
	Key   string
	Value string
}

type OrderingField string

const (
	OrderByNamespace = OrderingField("entity.namespace")
	OrderByName      = OrderingField("entity.name")
)

type Ordering struct {
	OrderBy    OrderingField
	Descending bool
}

type Pagination struct {
	Limit  int
	Offset int
}
