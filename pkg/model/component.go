package model

type Component struct {
	Entity
	Spec ComponentSpec `yaml:"spec"`
}

type ComponentSpec struct {
	Type           string      `yaml:"type"`
	Lifecycle      string      `yaml:"lifecycle"`
	Owner          string      `yaml:"owner"`
	System         EntityRef   `yaml:"system,omitempty"`
	SubcomponentOf EntityRef   `yaml:"subcomponentOf,omitempty"`
	ProvidesAPIs   []EntityRef `yaml:"providesApis,omitempty"`
	ConsumesAPIs   []EntityRef `yaml:"consumesApis,omitempty"`
	DependsOn      []EntityRef `yaml:"dependsOn,omitempty"`
	DependencyOf   []EntityRef `yaml:"dependencyOf,omitempty"`
}
