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

func TestListComponents(t *testing.T) {
	component1 := model.Component{
		Entity: model.Entity{
			APIVersion: "backstage.io/v1alpha1",
			Kind:       model.KindComponent,
			Metadata: model.Metadata{
				Namespace: "default",
				Name:      "component1",
			},
		},
		Spec: model.ComponentSpec{
			Type:      model.ComponentTypeService,
			Lifecycle: model.ComponentLifecycleExperimental,
			Owner:     model.TestFullUser.EntityRef(),
		},
	}

	component2 := component1
	component2.Entity.Metadata.Namespace = "ns1"
	component2.Entity.Metadata.Name = "component2"
	component2.Spec.Type = model.ComponentTypeLibrary

	component3 := component1
	component3.Entity.Metadata.Namespace = "ns1"
	component3.Entity.Metadata.Name = "component3"
	component3.Spec.Type = model.ComponentTypeWebsite

	type testCase struct {
		components         []model.Component
		filters            []Filter
		ordering           Ordering
		pagination         Pagination
		expectedEntityRefs []model.EntityRef
		description        string
	}
	tcs := []testCase{
		{
			components:         nil,
			expectedEntityRefs: nil,
			description:        "empty",
		},
		{
			components: []model.Component{
				component1,
			},
			expectedEntityRefs: []model.EntityRef{
				component1.EntityRef(),
			},
			description: "single hit, no filters",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			expectedEntityRefs: []model.EntityRef{
				component1.EntityRef(),
				component2.EntityRef(),
			},
			description: "multiple hits, no filters",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			filters: []Filter{
				{
					Key:   "entity.namespace",
					Value: "ns1",
				},
			},
			expectedEntityRefs: []model.EntityRef{
				component2.EntityRef(),
			},
			description: "filter on namespace",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			filters: []Filter{
				{
					Key:   "entity.name",
					Value: "component1",
				},
			},
			expectedEntityRefs: []model.EntityRef{
				component1.EntityRef(),
			},
			description: "filter on name",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			filters: []Filter{
				{
					Key:   "entity.namespace",
					Value: "default",
				},
				{
					Key:   "entity.name",
					Value: "component1",
				},
			},
			expectedEntityRefs: []model.EntityRef{
				component1.EntityRef(),
			},
			description: "filter on namespace and name",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			filters: []Filter{
				{
					Key:   "entity.namespace",
					Value: "ns1",
				},
				{
					Key:   "entity.name",
					Value: "component1",
				},
			},
			expectedEntityRefs: nil,
			description:        "filter on namespace and name, no hits",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			ordering: Ordering{
				OrderBy: OrderByName,
			},
			expectedEntityRefs: []model.EntityRef{
				component1.EntityRef(),
				component2.EntityRef(),
			},
			description: "multiple hits, ascending name order",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			ordering: Ordering{
				OrderBy:    OrderByName,
				Descending: true,
			},
			expectedEntityRefs: []model.EntityRef{
				component2.EntityRef(),
				component1.EntityRef(),
			},
			description: "multiple hits, descending name order",
		},
		{
			components: []model.Component{
				component1,
				component2,
			},
			ordering: Ordering{
				OrderBy: OrderByNamespace,
			},
			expectedEntityRefs: []model.EntityRef{
				component1.EntityRef(),
				component2.EntityRef(),
			},
			description: "multiple hits, ascending namespace order",
		},
		{
			components: []model.Component{
				component1,
				component2,
				component3,
			},
			pagination: Pagination{
				Limit: 2,
			},
			expectedEntityRefs: []model.EntityRef{
				component1.EntityRef(),
				component2.EntityRef(),
			},
			description: "multiple hits, first page",
		},
		{
			components: []model.Component{
				component1,
				component2,
				component3,
			},
			pagination: Pagination{
				Limit:  2,
				Offset: 2,
			},
			expectedEntityRefs: []model.EntityRef{
				component3.EntityRef(),
			},
			description: "multiple hits, second page",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			store := testStore(t)
			for _, c := range tc.components {
				_, err := store.CreateComponent(c)
				assert.NoError(t, err)
			}

			refs, pagination, err := store.ListComponents(tc.filters, tc.ordering, tc.pagination)
			assert.NoError(t, err)
			if tc.ordering.OrderBy != "" {
				assert.Equal(t, tc.expectedEntityRefs, refs)
			} else {
				assert.ElementsMatch(t, tc.expectedEntityRefs, refs)
			}
			if tc.pagination.Limit > 0 {
				assert.Equal(t, tc.pagination.Limit, pagination.Limit)
				assert.Equal(t, tc.pagination.Offset+len(refs), pagination.Offset)
			}
		})
	}
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
