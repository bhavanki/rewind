package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bhavanki/rewind/internal/store"
	"github.com/bhavanki/rewind/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestCreateEntity_Component(t *testing.T) {
	r := gin.Default()
	var component model.Component
	s := &store.StoreMock{
		CreateComponentFunc: func(c model.Component) (model.Component, error) {
			component = c
			return c, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	componentYAML, err := yaml.Marshal(model.TestFullComponent)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/component/my-namespace/my-service", strings.NewReader(string(componentYAML)))
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, model.TestFullComponent, component)
}

func TestReadEntity_Component(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		ReadComponentFunc: func(ref model.EntityRef) (model.Component, error) {
			return model.TestFullComponent, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/component/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var component model.Component
	err = yaml.Unmarshal(w.Body.Bytes(), &component)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullComponent, component)
}
