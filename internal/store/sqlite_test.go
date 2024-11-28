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

func TestCreateComponentAndReadComponent(t *testing.T) {
	store := testStore(t)

	c, err := store.CreateComponent(model.TestFullComponent)
	assert.NoError(t, err)
	id := c.ID
	c = model.TestFullComponent
	c.ID = id

	r, err := store.ReadComponent(model.TestFullComponent.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, c, r)
}

func TestDeleteComponent(t *testing.T) {
	store := testStore(t)

	c, err := store.CreateComponent(model.TestFullComponent)
	assert.NoError(t, err)
	id := c.ID

	d, err := store.DeleteComponent(model.TestFullComponent.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, c, d)

	_, err = readEntity(c.EntityRef(), store.db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(componentSelectStatement, id)
	assert.NoError(t, err)
	assert.False(t, rows.Next())
}

// ---

func TestCreateAPIAndReadAPI(t *testing.T) {
	store := testStore(t)

	a, err := store.CreateAPI(model.TestFullAPI)
	assert.NoError(t, err)
	id := a.ID
	a = model.TestFullAPI
	a.ID = id

	r, err := store.ReadAPI(model.TestFullAPI.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, a, r)
}

func TestDeleteAPI(t *testing.T) {
	store := testStore(t)

	a, err := store.CreateAPI(model.TestFullAPI)
	assert.NoError(t, err)
	id := a.ID

	d, err := store.DeleteAPI(model.TestFullAPI.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, a, d)

	_, err = readEntity(a.EntityRef(), store.db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(apiSelectStatement, id)
	assert.NoError(t, err)
	assert.False(t, rows.Next())
}

// ---

func TestCreateUserAndReadUser(t *testing.T) {
	store := testStore(t)

	u, err := store.CreateUser(model.TestFullUser)
	assert.NoError(t, err)
	id := u.ID
	u = model.TestFullUser
	u.ID = id

	r, err := store.ReadUser(model.TestFullUser.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, u, r)
}

func TestDeleteUser(t *testing.T) {
	store := testStore(t)

	u, err := store.CreateUser(model.TestFullUser)
	assert.NoError(t, err)
	id := u.ID

	d, err := store.DeleteUser(model.TestFullUser.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, u, d)

	_, err = readEntity(u.EntityRef(), store.db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(userSelectStatement, id)
	assert.NoError(t, err)
	assert.False(t, rows.Next())
}

// ---

func TestCreateGroupAndReadGroup(t *testing.T) {
	store := testStore(t)

	g, err := store.CreateGroup(model.TestFullGroup)
	assert.NoError(t, err)
	id := g.ID
	g = model.TestFullGroup
	g.ID = id

	r, err := store.ReadGroup(model.TestFullGroup.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, g, r)
}

func TestDeleteGroup(t *testing.T) {
	store := testStore(t)

	g, err := store.CreateGroup(model.TestFullGroup)
	assert.NoError(t, err)
	id := g.ID

	d, err := store.DeleteGroup(model.TestFullGroup.EntityRef())
	assert.NoError(t, err)

	assert.Equal(t, g, d)

	_, err = readEntity(g.EntityRef(), store.db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(groupSelectStatement, id)
	assert.NoError(t, err)
	assert.False(t, rows.Next())
}
