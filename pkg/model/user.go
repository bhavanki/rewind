package model

type User struct {
	Entity
	Spec UserSpec `yaml:"spec"`
}

type UserSpec struct {
	Profile  UserProfile `yaml:"profile,omitempty"`
	MemberOf []EntityRef `yaml:"memberOf"`
}

type UserProfile struct {
	DisplayName string `yaml:"displayName,omitempty"`
	Email       string `yaml:"email,omitempty"`
	Picture     string `yaml:"picture,omitempty"`
}
