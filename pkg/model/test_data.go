package model

var (
	TestOwnerEntityRef = EntityRef{
		Kind:      "user",
		Namespace: "default",
		Name:      "owner",
	}
	TestSystemEntityRef = EntityRef{
		Kind:      "system",
		Namespace: "default",
		Name:      "down",
	}
	TestComponentEntityRef = EntityRef{
		Kind:      "component",
		Namespace: "default",
		Name:      "component",
	}
	TestAPI1EntityRef = EntityRef{
		Kind:      "api",
		Namespace: "default",
		Name:      "api1",
	}
	TestAPI2EntityRef = EntityRef{
		Kind:      "api",
		Namespace: "default",
		Name:      "api2",
	}
	TestGroupEntityRef = EntityRef{
		Kind:      "group",
		Namespace: "default",
		Name:      "group",
	}
	TestResource1EntityRef = EntityRef{
		Kind:      "resource",
		Namespace: "default",
		Name:      "resource1",
	}
	TestResource2EntityRef = EntityRef{
		Kind:      "resource",
		Namespace: "default",
		Name:      "resource2",
	}

	TestFullEntity = Entity{
		APIVersion: "backstage.io/v1alpha1",
		Kind:       "unknown",
		Metadata: Metadata{
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
			Links: []Link{
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
	TestFullComponent = Component{
		Entity: TestFullEntity,
		Spec: ComponentSpec{
			Type:           "service",
			Lifecycle:      "experimental",
			Owner:          TestOwnerEntityRef,
			System:         TestSystemEntityRef,
			SubcomponentOf: TestComponentEntityRef,
			ProvidesAPIs: []EntityRef{
				TestAPI1EntityRef,
			},
			ConsumesAPIs: []EntityRef{
				TestAPI2EntityRef,
			},
			DependsOn: []EntityRef{
				TestResource1EntityRef,
			},
			DependencyOf: []EntityRef{
				TestResource2EntityRef,
			},
		},
	}
	TestFullAPI = API{
		Entity: TestFullEntity,
		Spec: APISpec{
			Type:       "openapi",
			Lifecycle:  "experimental",
			Owner:      TestOwnerEntityRef,
			System:     TestSystemEntityRef,
			Definition: "definition",
		},
	}
	TestFullUser = User{
		Entity: TestFullEntity,
		Spec: UserSpec{
			Profile: UserProfile{
				DisplayName: "displayName",
				Email:       "email",
				Picture:     "picture",
			},
			MemberOf: []EntityRef{
				TestGroupEntityRef,
			},
		},
	}
	TestFullGroup = Group{
		Entity: TestFullEntity,
		Spec: GroupSpec{
			Type: "team",
			Profile: GroupProfile{
				DisplayName: "displayName",
				Email:       "email",
				Picture:     "picture",
			},
			Parent: TestGroupEntityRef,
			Children: []EntityRef{
				{
					Kind:      "group",
					Namespace: "default",
					Name:      "child",
				},
			},
			Members: []EntityRef{
				TestFullUser.EntityRef(),
			},
		},
	}

	TestMinimalEntity = Entity{
		APIVersion: "backstage.io/v1alpha1",
		Kind:       "api",
		Metadata: Metadata{
			Name:      "my-service",
			Namespace: "my-namespace",
		},
	}
)

func init() {
	TestFullComponent.Entity.Kind = "component"
	TestFullAPI.Entity.Kind = "api"
	TestFullUser.Entity.Kind = "user"
	TestFullGroup.Entity.Kind = "user"
}
