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

func TestUpdateComponent(t *testing.T) {
	store := testStore(t)

	c, err := store.CreateComponent(model.TestFullComponent)
	assert.NoError(t, err)

	c.Metadata.Title = "my-new-title"
	c.Metadata.Description = "my-new-description"
	c.Metadata.Labels = map[string]string{
		"key0": "value0",
		"key1": "value1a",
	}
	c.Metadata.Annotations = map[string]string{
		"keya": "valuea1",
	}
	c.Metadata.Tags = []string{"tag0"}
	c.Metadata.Links = []model.Link{
		{
			URL:   "http://example.com/url0",
			Title: "link0",
			Icon:  "icon0",
			Type:  "linktype0",
		},
		{
			URL:   "http://example.com/url1",
			Title: "link1a",
			Icon:  "icon1a",
			Type:  "linktype1a",
		},
	}
	c.Spec.Type = model.ComponentTypeLibrary
	c.Spec.Lifecycle = model.ComponentLifecycleProduction
	c.Spec.Owner = model.TestOwner2EntityRef
	c.Spec.System = model.TestSystem2EntityRef
	c.Spec.SubcomponentOf = model.TestComponent2EntityRef
	c.Spec.ProvidesAPIs = []model.EntityRef{
		model.TestAPI2EntityRef,
	}
	c.Spec.ConsumesAPIs = nil
	c.Spec.DependsOn = []model.EntityRef{
		model.TestResource2EntityRef,
	}
	c.Spec.DependencyOf = nil

	u, err := store.UpdateComponent(c)
	assert.NoError(t, err)
	assert.Equal(t, c, u)

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

	_, err = store.readEntity(c.EntityRef())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(componentSelectStatement, id)
	defer rows.Close()
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

func TestUpdateAPI(t *testing.T) {
	store := testStore(t)

	a, err := store.CreateAPI(model.TestFullAPI)
	assert.NoError(t, err)

	// metadata updates tested for component
	a.Spec.Type = model.APITypeGRPC
	a.Spec.Lifecycle = model.APILifecycleProduction
	a.Spec.Owner = model.TestOwner2EntityRef
	a.Spec.System = model.TestSystem2EntityRef
	a.Spec.Definition = "my-new-definition"

	u, err := store.UpdateAPI(a)
	assert.NoError(t, err)
	assert.Equal(t, a, u)

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

	_, err = store.readEntity(a.EntityRef())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(apiSelectStatement, id)
	defer rows.Close()
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

func TestUpdateUser(t *testing.T) {
	store := testStore(t)

	u, err := store.CreateUser(model.TestFullUser)
	assert.NoError(t, err)

	// metadata updates tested for component
	u.Spec.Profile.DisplayName = "new-displayName"
	u.Spec.Profile.Email = "new-email"
	u.Spec.Profile.Picture = "new-picture"
	u.Spec.MemberOf = []model.EntityRef{
		model.TestGroup2EntityRef,
	}

	uu, err := store.UpdateUser(u)
	assert.NoError(t, err)
	assert.Equal(t, u, uu)

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

	_, err = store.readEntity(u.EntityRef())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(userSelectStatement, id)
	defer rows.Close()
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

func TestUpdateGroup(t *testing.T) {
	store := testStore(t)

	g, err := store.CreateGroup(model.TestFullGroup)
	assert.NoError(t, err)

	// metadata updates tested for component
	g.Spec.Type = "business-unit"
	g.Spec.Profile.DisplayName = "new-displayName"
	g.Spec.Profile.Email = "new-email"
	g.Spec.Profile.Picture = "new-picture"
	g.Spec.Parent = model.TestGroup2EntityRef
	g.Spec.Children = []model.EntityRef{
		model.TestGroupEntityRef,
	}
	g.Spec.Members = []model.EntityRef{
		model.TestUser2EntityRef,
	}

	u, err := store.UpdateGroup(g)
	assert.NoError(t, err)
	assert.Equal(t, g, u)

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

	_, err = store.readEntity(g.EntityRef())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	rows, err := store.db.Queryx(groupSelectStatement, id)
	defer rows.Close()
	assert.NoError(t, err)
	assert.False(t, rows.Next())
}
