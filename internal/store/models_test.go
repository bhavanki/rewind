package store

import (
	"testing"

	"github.com/bhavanki/rewind/pkg/model"
	"github.com/stretchr/testify/assert"
)

var (
	fullTestDBEntity = Entity{
		APIVersion:  "backstage.io/v1alpha1",
		Kind:        "Component",
		Name:        "my-service",
		Namespace:   "my-namespace",
		Title:       ptr("my-title"),
		Description: ptr("my-description"),
		Labels: []Label{
			{
				Key:   "key1",
				Value: "value1",
			},
			{
				Key:   "key2",
				Value: "value2",
			},
			{
				Key:   "key3",
				Value: "value3",
			},
		},
		Annotations: []Annotation{
			{
				Key:   "keya",
				Value: "valuea",
			},
			{
				Key:   "keyb",
				Value: "valueb",
			},
			{
				Key:   "keyc",
				Value: "valuec",
			},
		},
		Tags: ptr("tag1,tag2,tag3"),
		Links: []Link{
			{
				URL:   "http://example.com/url1",
				Title: ptr("link1"),
				Icon:  ptr("icon1"),
				Type:  ptr("linktype1"),
			},
			{
				URL:   "http://example.com/url2",
				Title: ptr("link2"),
				Icon:  ptr("icon2"),
				Type:  ptr("linktype2"),
			},
		},
	}

	fullTestModelEntity = model.Entity{
		APIVersion: "backstage.io/v1alpha1",
		Kind:       "Component",
		Metadata: model.Metadata{
			Name:        "my-service",
			Namespace:   "my-namespace",
			Title:       "my-title",
			Description: "my-description",
			Labels: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			Annotations: map[string]string{
				"keya": "valuea",
				"keyb": "valueb",
				"keyc": "valuec",
			},
			Tags: []string{"tag1", "tag2", "tag3"},
			Links: []model.Link{
				{
					URL:   "http://example.com/url1",
					Title: "link1",
					Icon:  "icon1",
					Type:  "linktype1",
				},
				{
					URL:   "http://example.com/url2",
					Title: "link2",
					Icon:  "icon2",
					Type:  "linktype2",
				},
			},
		},
	}
)

func TestFromDBModel(t *testing.T) {
	type testCase struct {
		dbe         Entity
		me          model.Entity
		description string
	}
	tcs := []testCase{
		{
			dbe:         fullTestDBEntity,
			me:          fullTestModelEntity,
			description: "full model",
		},
		{
			dbe: Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Name:       "my-service",
				Namespace:  "my-namespace",
			},
			me: model.Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Metadata: model.Metadata{
					Name:      "my-service",
					Namespace: "my-namespace",
				},
			},
			description: "minimal model",
		},
		{
			dbe: Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Name:       "my-service",
				Namespace:  "my-namespace",
				Links: []Link{
					{
						URL: "http://example.com/url1",
					},
					{
						URL: "http://example.com/url2",
					},
				},
			},
			me: model.Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Metadata: model.Metadata{
					Name:      "my-service",
					Namespace: "my-namespace",
					Links: []model.Link{
						{
							URL: "http://example.com/url1",
						},
						{
							URL: "http://example.com/url2",
						},
					},
				},
			},
			description: "minimal links",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			e, err := FromDBModel(tc.dbe)
			assert.NoError(t, err)
			assert.Equal(t, tc.me, e)
		})
	}
}

func TestToDBModel(t *testing.T) {
	type testCase struct {
		me          model.Entity
		dbe         Entity
		description string
	}
	tcs := []testCase{
		{
			me:          fullTestModelEntity,
			dbe:         fullTestDBEntity,
			description: "full model",
		},
		{
			me: model.Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Metadata: model.Metadata{
					Name:      "my-service",
					Namespace: "my-namespace",
				},
			},
			dbe: Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Name:       "my-service",
				Namespace:  "my-namespace",
			},
			description: "minimal model",
		},
		{
			me: model.Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Metadata: model.Metadata{
					Name:      "my-service",
					Namespace: "my-namespace",
					Links: []model.Link{
						{
							URL: "http://example.com/url1",
						},
						{
							URL: "http://example.com/url2",
						},
					},
				},
			},
			dbe: Entity{
				APIVersion: "backstage.io/v1alpha1",
				Kind:       "Component",
				Name:       "my-service",
				Namespace:  "my-namespace",
				Links: []Link{
					{
						URL: "http://example.com/url1",
					},
					{
						URL: "http://example.com/url2",
					},
				},
			},
			description: "minimal links",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			e, err := ToDBModel(tc.me)
			assert.NoError(t, err)
			assert.Equal(t, tc.dbe, e)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
