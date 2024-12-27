package model

const (
	APITypeOpenAPI  = "openapi"
	APITypeAsyncAPI = "asyncapi"
	APITypeGraphQL  = "graphql"
	APITypeGRPC     = "grpc"

	APILifecycleExperimental = "experimental"
	APILifecycleProduction   = "production"
	APILifecycleDeprecated   = "deprecated"
)

type API struct {
	Entity `yaml:"entity,inline"`
	Spec   APISpec `yaml:"spec"`
}

type APISpec struct {
	Type       string    `yaml:"type"`
	Lifecycle  string    `yaml:"lifecycle"`
	Owner      EntityRef `yaml:"owner"`
	System     EntityRef `yaml:"system,omitempty"`
	Definition string    `yaml:"definition"`
}
