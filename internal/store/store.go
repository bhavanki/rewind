package store

import (
	"github.com/bhavanki/rewind/pkg/model"
)

type Store interface {
	CreateEntity(e model.Entity) (model.Entity, error)
	ReadEntity(ref model.EntityRef) (model.Entity, error)
	DeleteEntity(ref model.EntityRef) (model.Entity, error)
}
