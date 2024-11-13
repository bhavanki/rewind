package model

// import (
// 	"gopkg.in/yaml.v3"
// )

type Entity struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
}

type Metadata struct {
	Name        string            `yaml:"name"`
	Namespace   string            `yaml:"namespace"`
	Title       string            `yaml:"title,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
	Tags        []string          `yaml:"tags,omitempty"`
	Links       []Link            `yaml:"links,omitempty"`
}

type Link struct {
	URL   string `yaml:"url"`
	Title string `yaml:"title,omitempty"`
	Icon  string `yaml:"icon,omitempty"`
	Type  string `yaml:"type,omitempty"`
}

func (e Entity) EntityRef() EntityRef {
	return EntityRef{
		Kind:      e.Kind,
		Namespace: e.Metadata.Namespace,
		Name:      e.Metadata.Name,
	}
}
