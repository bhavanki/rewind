package store

import (
	"github.com/bhavanki/rewind/pkg/model"
)

type Store interface {
	CreateAPI(a model.API) (model.API, error)
	ReadAPI(ref model.EntityRef) (model.API, error)
	DeleteAPI(ref model.EntityRef) (model.API, error)

	CreateUser(u model.User) (model.User, error)
	ReadUser(ref model.EntityRef) (model.User, error)
	DeleteUser(ref model.EntityRef) (model.User, error)

	CreateGroup(g model.Group) (model.Group, error)
	ReadGroup(ref model.EntityRef) (model.Group, error)
	DeleteGroup(ref model.EntityRef) (model.Group, error)
}
