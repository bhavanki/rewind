package store

import (
	"github.com/bhavanki/rewind/pkg/model"
)

var (
	testOwnerEntityRef = model.EntityRef{
		Kind:      "user",
		Namespace: "default",
		Name:      "owner",
	}
	testSystemEntityRef = model.EntityRef{
		Kind:      "system",
		Namespace: "default",
		Name:      "down",
	}
	testGroupEntityRef = model.EntityRef{
		Kind:      "group",
		Namespace: "default",
		Name:      "group",
	}

	// fullTestDBEntity = Entity{
	// 	APIVersion:  "backstage.io/v1alpha1",
	// 	Kind:        "api",
	// 	Name:        "my-service",
	// 	Namespace:   "my-namespace",
	// 	Title:       ptr("my-title"),
	// 	Description: ptr("my-description"),
	// 	Labels: []Label{
	// 		{
	// 			Key:   "key1",
	// 			Value: "value1",
	// 		},
	// 		{
	// 			Key:   "key2",
	// 			Value: "value2",
	// 		},
	// 		{
	// 			Key:   "key3",
	// 			Value: "value3",
	// 		},
	// 	},
	// 	Annotations: []Annotation{
	// 		{
	// 			Key:   "keya",
	// 			Value: "valuea",
	// 		},
	// 		{
	// 			Key:   "keyb",
	// 			Value: "valueb",
	// 		},
	// 		{
	// 			Key:   "keyc",
	// 			Value: "valuec",
	// 		},
	// 	},
	// 	Tags: ptr("tag1,tag2,tag3"),
	// 	Links: []Link{
	// 		{
	// 			URL:   "http://example.com/url1",
	// 			Title: ptr("link1"),
	// 			Icon:  ptr("icon1"),
	// 			Type:  ptr("linktype1"),
	// 		},
	// 		{
	// 			URL:   "http://example.com/url2",
	// 			Title: ptr("link2"),
	// 			Icon:  ptr("icon2"),
	// 			Type:  ptr("linktype2"),
	// 		},
	// 	},
	// }
	// fullTestDBAPI = API{
	// 	Entity:     fullTestDBEntity,
	// 	Type:       "openapi",
	// 	Lifecycle:  "experimental",
	// 	Owner:      ownerEntityRef,
	// 	System:     &systemEntityRef,
	// 	Definition: "definition",
	// }
	// fullTestDBUser = User{
	// 	Entity:      fullTestDBEntity,
	// 	DisplayName: ptr("displayName"),
	// 	Email:       ptr("email"),
	// 	Picture:     ptr("picture"),
	// 	MemberOf: []model.EntityRef{
	// 		groupEntityRef,
	// 	},
	// }

	testFullEntity = model.Entity{
		APIVersion: "backstage.io/v1alpha1",
		Kind:       "unknown",
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
	testFullAPI = model.API{
		Entity: testFullEntity,
		Spec: model.APISpec{
			Type:       "openapi",
			Lifecycle:  "experimental",
			Owner:      testOwnerEntityRef,
			System:     testSystemEntityRef,
			Definition: "definition",
		},
	}
	testFullUser = model.User{
		Entity: testFullEntity,
		Spec: model.UserSpec{
			Profile: model.UserProfile{
				DisplayName: "displayName",
				Email:       "email",
				Picture:     "picture",
			},
			MemberOf: []model.EntityRef{
				testGroupEntityRef,
			},
		},
	}
	testFullGroup = model.Group{
		Entity: testFullEntity,
		Spec: model.GroupSpec{
			Type: "team",
			Profile: model.GroupProfile{
				DisplayName: "displayName",
				Email:       "email",
				Picture:     "picture",
			},
			Parent: testGroupEntityRef,
			Children: []model.EntityRef{
				{
					Kind:      "group",
					Namespace: "default",
					Name:      "child",
				},
			},
			Members: []model.EntityRef{
				testFullUser.EntityRef(),
			},
		},
	}

	// minimalTestDBEntity = Entity{
	// 	APIVersion: "backstage.io/v1alpha1",
	// 	Kind:       "api",
	// 	Name:       "my-service",
	// 	Namespace:  "my-namespace",
	// }

	testMinimalEntity = model.Entity{
		APIVersion: "backstage.io/v1alpha1",
		Kind:       "api",
		Metadata: model.Metadata{
			Name:      "my-service",
			Namespace: "my-namespace",
		},
	}
)

func init() {
	testFullAPI.Entity.Kind = "api"
	testFullUser.Entity.Kind = "user"
}
