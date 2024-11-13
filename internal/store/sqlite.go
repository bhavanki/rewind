package store

import (
	"fmt"

	"github.com/bhavanki/rewind/pkg/model"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sqliteStore struct {
	db *gorm.DB
}

func NewSqliteStore(connString string) (*sqliteStore, error) {
	db, err := gorm.Open(sqlite.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database at %s: %w", connString, err)
	}

	err = db.AutoMigrate(&Entity{}, &Label{}, &Annotation{}, &Link{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database: %w", err)
	}

	return &sqliteStore{
		db: db,
	}, nil
}

func (s sqliteStore) CreateEntity(e model.Entity) (model.Entity, error) {
	dbe, err := ToDBModel(e)
	if err != nil {
		return e, err
	}
	result := s.db.Create(&dbe)
	if result.Error != nil {
		return e, fmt.Errorf("failed to create entity: %w", result.Error)
	}

	return e, nil
}
func (s sqliteStore) ReadEntity(ref model.EntityRef) (model.Entity, error) {
	dbe := Entity{}
	result := s.db.Where(&Entity{
		Kind:      ref.Kind,
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}).First(&dbe)
	if result.Error != nil {
		return model.Entity{}, fmt.Errorf("failed to read entity: %w", result.Error)
	}

	e, err := FromDBModel(dbe)
	if err != nil {
		return model.Entity{}, err
	}
	return e, err
}
func (s sqliteStore) DeleteEntity(ref model.EntityRef) (model.Entity, error) {
	dbe := Entity{}
	result := s.db.Clauses(clause.Returning{}).Where(&Entity{
		Kind:      ref.Kind,
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}).Delete(&dbe)
	if result.Error != nil {
		return model.Entity{}, fmt.Errorf("failed to delete entity: %w", result.Error)
	}

	e, err := FromDBModel(dbe)
	if err != nil {
		return model.Entity{}, err
	}
	return e, err
}
