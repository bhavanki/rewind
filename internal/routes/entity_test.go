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

// func TestUpdateEntity_Component(t *testing.T) {

// }

func TestDeleteEntity_Component(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		DeleteComponentFunc: func(ref model.EntityRef) (model.Component, error) {
			return model.TestFullComponent, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/component/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var component model.Component
	err = yaml.Unmarshal(w.Body.Bytes(), &component)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullComponent, component)
}

func TestCreateEntity_API(t *testing.T) {
	r := gin.Default()
	var api model.API
	s := &store.StoreMock{
		CreateAPIFunc: func(a model.API) (model.API, error) {
			api = a
			return a, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	apiYAML, err := yaml.Marshal(model.TestFullAPI)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/api/my-namespace/my-service", strings.NewReader(string(apiYAML)))
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, model.TestFullAPI, api)
}

func TestReadEntity_API(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		ReadAPIFunc: func(ref model.EntityRef) (model.API, error) {
			return model.TestFullAPI, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var api model.API
	err = yaml.Unmarshal(w.Body.Bytes(), &api)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullAPI, api)
}

func TestDeleteEntity_API(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		DeleteAPIFunc: func(ref model.EntityRef) (model.API, error) {
			return model.TestFullAPI, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/api/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var api model.API
	err = yaml.Unmarshal(w.Body.Bytes(), &api)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullAPI, api)
}

func TestCreateEntity_User(t *testing.T) {
	r := gin.Default()
	var user model.User
	s := &store.StoreMock{
		CreateUserFunc: func(u model.User) (model.User, error) {
			user = u
			return u, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	userYAML, err := yaml.Marshal(model.TestFullUser)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/user/my-namespace/my-service", strings.NewReader(string(userYAML)))
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, model.TestFullUser, user)
}

func TestReadEntity_User(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		ReadUserFunc: func(ref model.EntityRef) (model.User, error) {
			return model.TestFullUser, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var user model.User
	err = yaml.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullUser, user)
}

func TestDeleteEntity_User(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		DeleteUserFunc: func(ref model.EntityRef) (model.User, error) {
			return model.TestFullUser, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/user/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var user model.User
	err = yaml.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullUser, user)
}

func TestCreateEntity_Group(t *testing.T) {
	r := gin.Default()
	var group model.Group
	s := &store.StoreMock{
		CreateGroupFunc: func(g model.Group) (model.Group, error) {
			group = g
			return g, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	groupYAML, err := yaml.Marshal(model.TestFullGroup)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/group/my-namespace/my-service", strings.NewReader(string(groupYAML)))
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, model.TestFullGroup, group)
}

func TestReadEntity_Group(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		ReadGroupFunc: func(ref model.EntityRef) (model.Group, error) {
			return model.TestFullGroup, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/group/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var group model.Group
	err = yaml.Unmarshal(w.Body.Bytes(), &group)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullGroup, group)
}

func TestDeleteEntity_Group(t *testing.T) {
	r := gin.Default()
	s := &store.StoreMock{
		DeleteGroupFunc: func(ref model.EntityRef) (model.Group, error) {
			return model.TestFullGroup, nil
		},
	}
	SetupRoutes(r, s)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/group/my-namespace/my-service", nil)
	require.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var group model.Group
	err = yaml.Unmarshal(w.Body.Bytes(), &group)
	assert.NoError(t, err)
	assert.Equal(t, model.TestFullGroup, group)
}
