package model

var (
	TestOwnerEntityRef = EntityRef{
		Kind:      KindUser,
		Namespace: "default",
		Name:      "owner",
	}
	TestOwner2EntityRef = EntityRef{
		Kind:      KindUser,
		Namespace: "default",
		Name:      "owner2",
	}
	TestSystemEntityRef = EntityRef{
		Kind:      KindSystem,
		Namespace: "default",
		Name:      "down",
	}
	TestSystem2EntityRef = EntityRef{
		Kind:      KindSystem,
		Namespace: "default",
		Name:      "shock",
	}
	TestComponentEntityRef = EntityRef{
		Kind:      KindComponent,
		Namespace: "default",
		Name:      "component",
	}
	TestComponent2EntityRef = EntityRef{
		Kind:      KindComponent,
		Namespace: "default",
		Name:      "component2",
	}
	TestAPI1EntityRef = EntityRef{
		Kind:      KindAPI,
		Namespace: "default",
		Name:      "api1",
	}
	TestAPI2EntityRef = EntityRef{
		Kind:      KindAPI,
		Namespace: "default",
		Name:      "api2",
	}
	TestUser2EntityRef = EntityRef{
		Kind:      KindUser,
		Namespace: "default",
		Name:      "user2",
	}
	TestGroupEntityRef = EntityRef{
		Kind:      KindGroup,
		Namespace: "default",
		Name:      "group",
	}
	TestGroup2EntityRef = EntityRef{
		Kind:      KindGroup,
		Namespace: "default",
		Name:      "group2",
	}
	TestResource1EntityRef = EntityRef{
		Kind:      KindResource,
		Namespace: "default",
		Name:      "resource1",
	}
	TestResource2EntityRef = EntityRef{
		Kind:      KindResource,
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
			Type:           ComponentTypeService,
			Lifecycle:      ComponentLifecycleExperimental,
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
			Type:       APITypeOpenAPI,
			Lifecycle:  APILifecycleExperimental,
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
					Kind:      KindGroup,
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
		Kind:       KindAPI,
		Metadata: Metadata{
			Name:      "my-service",
			Namespace: "my-namespace",
		},
	}
)

func init() {
	TestFullComponent.Entity.Kind = KindComponent
	TestFullAPI.Entity.Kind = KindAPI
	TestFullUser.Entity.Kind = KindUser
	TestFullGroup.Entity.Kind = KindGroup
}
