package model

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

//go:embed testdata/component.yaml
var componentYAMLBytes []byte

func TestComponentYAML(t *testing.T) {
	var component Component
	err := yaml.Unmarshal(componentYAMLBytes, &component)

	assert.NoError(t, err)
	assert.Equal(t, TestFullComponent, component)

	out, err := yaml.Marshal(component)

	assert.NoError(t, err)

	var component2 Component
	err = yaml.Unmarshal(out, &component2)

	assert.NoError(t, err)
	assert.Equal(t, TestFullComponent, component2)
}
