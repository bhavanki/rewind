package model

type API struct {
	Entity
	Spec APISpec `yaml:"spec"`
}

type APISpec struct {
	Type       string    `yaml:"type"`
	Lifecycle  string    `yaml:"lifecycle"`
	Owner      EntityRef `yaml:"owner"`
	System     EntityRef `yaml:"system,omitempty"`
	Definition string    `yaml:"definition"`
}
