package model

type Group struct {
	Entity
	Spec GroupSpec `yaml:"spec"`
}

type GroupSpec struct {
	Type     string       `yaml:"type"`
	Profile  GroupProfile `yaml:"profile,omitempty"`
	Parent   EntityRef    `yaml:"parent,omitempty"`
	Children []EntityRef  `yaml:"children"`
	Members  []EntityRef  `yaml:"members,omitempty"`
}

type GroupProfile struct {
	DisplayName string `yaml:"displayName,omitempty"`
	Email       string `yaml:"email,omitempty"`
	Picture     string `yaml:"picture,omitempty"`
}
