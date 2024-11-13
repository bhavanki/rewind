package store

import (
	"testing"

	"github.com/bhavanki/rewind/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testStore(t *testing.T) *sqliteStore {
	store, err := NewSqliteStore("file::memory:")
	require.NoError(t, err)
	return store
}

func TestCreateEntityAndReadEntity(t *testing.T) {
	store := testStore(t)

	e := model.Entity{
		APIVersion: "backstage.io/v1alpha1",
		Kind:       "Component",
		Metadata: model.Metadata{
			Name:      "my-service",
			Namespace: "my-namespace",
		},
	}

	_, err := store.CreateEntity(e)
	assert.NoError(t, err)

	r, err := store.ReadEntity(e.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, e, r)
}

func TestDeleteEntity(t *testing.T) {
	store := testStore(t)

	e := model.Entity{
		APIVersion: "backstage.io/v1alpha1",
		Kind:       "Component",
		Metadata: model.Metadata{
			Name:      "my-service",
			Namespace: "my-namespace",
		},
	}

	_, err := store.CreateEntity(e)
	assert.NoError(t, err)

	d, err := store.DeleteEntity(e.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, e, d)
}
