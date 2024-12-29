package routes

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/bhavanki/rewind/internal/store"
	"github.com/bhavanki/rewind/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreateEntity(c *gin.Context, store store.Store) {
	expectedEntityRef := expectedEntityRef(c)
	kind := expectedEntityRef.Kind

	switch kind {
	case model.KindComponent:
		var component model.Component
		if err := c.ShouldBindYAML(&component); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, component.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.CreateComponent(component); err != nil {
			slog.Error("failed to store component", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store component"})
			return
		}
	case model.KindAPI:
		var api model.API
		if err := c.ShouldBindYAML(&api); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, api.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.CreateAPI(api); err != nil {
			slog.Error("failed to store API", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store API"})
			return
		}
	case model.KindUser:
		var user model.User
		if err := c.ShouldBindYAML(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, user.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.CreateUser(user); err != nil {
			slog.Error("failed to store user", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store user"})
			return
		}
	case model.KindGroup:
		var group model.Group
		if err := c.ShouldBindYAML(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, group.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.CreateGroup(group); err != nil {
			slog.Error("failed to store group", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store group"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unsupported kind %s", kind)})
		return
	}

	c.Status(http.StatusCreated)
}

func ReadEntity(c *gin.Context, store store.Store) {
	expectedEntityRef := expectedEntityRef(c)
	kind := expectedEntityRef.Kind

	switch kind {
	case model.KindComponent:
		component, err := store.ReadComponent(expectedEntityRef)
		if err != nil {
			slog.Error("failed to read component", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read component"})
			return
		}
		c.YAML(http.StatusOK, component)
	case model.KindAPI:
		api, err := store.ReadAPI(expectedEntityRef)
		if err != nil {
			slog.Error("failed to read API", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read API"})
			return
		}
		c.YAML(http.StatusOK, api)
	case model.KindUser:
		user, err := store.ReadUser(expectedEntityRef)
		if err != nil {
			slog.Error("failed to read user", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read user"})
			return
		}
		c.YAML(http.StatusOK, user)
	case model.KindGroup:
		group, err := store.ReadGroup(expectedEntityRef)
		if err != nil {
			slog.Error("failed to read group", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read group"})
			return
		}
		c.YAML(http.StatusOK, group)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unsupported kind %s", kind)})
		return
	}
}

func UpdateEntity(c *gin.Context, store store.Store) {
	expectedEntityRef := expectedEntityRef(c)
	kind := expectedEntityRef.Kind

	switch kind {
	case model.KindComponent:
		var component model.Component
		if err := c.ShouldBindYAML(&component); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, component.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.UpdateComponent(component); err != nil {
			slog.Error("failed to update component", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update component"})
			return
		}
	case model.KindAPI:
		var api model.API
		if err := c.ShouldBindYAML(&api); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, api.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.UpdateAPI(api); err != nil {
			slog.Error("failed to update API", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update API"})
			return
		}
	case model.KindUser:
		var user model.User
		if err := c.ShouldBindYAML(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, user.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.UpdateUser(user); err != nil {
			slog.Error("failed to update user", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
			return
		}
	case model.KindGroup:
		var group model.Group
		if err := c.ShouldBindYAML(&group); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !verifyEntityRef(c, group.Entity.EntityRef(), expectedEntityRef) {
			return
		}
		if _, err := store.UpdateGroup(group); err != nil {
			slog.Error("failed to update group", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update group"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unsupported kind %s", kind)})
		return
	}

	c.Status(http.StatusAccepted)
}

func DeleteEntity(c *gin.Context, store store.Store) {
	expectedEntityRef := expectedEntityRef(c)
	kind := expectedEntityRef.Kind

	switch kind {
	case model.KindComponent:
		component, err := store.DeleteComponent(expectedEntityRef)
		if err != nil {
			slog.Error("failed to delete component", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete component"})
			return
		}
		c.YAML(http.StatusOK, component)
	case model.KindAPI:
		api, err := store.DeleteAPI(expectedEntityRef)
		if err != nil {
			slog.Error("failed to delete API", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete API"})
			return
		}
		c.YAML(http.StatusOK, api)
	case model.KindUser:
		user, err := store.DeleteUser(expectedEntityRef)
		if err != nil {
			slog.Error("failed to delete user", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
			return
		}
		c.YAML(http.StatusOK, user)
	case model.KindGroup:
		group, err := store.DeleteGroup(expectedEntityRef)
		if err != nil {
			slog.Error("failed to delete group", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete group"})
			return
		}
		c.YAML(http.StatusOK, group)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unsupported kind %s", kind)})
		return
	}
}

func ListEntities(c *gin.Context, st store.Store) {
	filters, ordering, pagination, err := processListParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("bad list parameter: %s", err)})
		return
	}

	var refs []model.EntityRef
	var nextPagination store.Pagination
	kind := c.Param("kind")
	switch kind {
	case model.KindComponent:
		refs, nextPagination, err = st.ListComponents(filters, ordering, pagination)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unsupported kind %s", kind)})
		return
	}
	if err != nil {
		slog.Error("failed to list entities", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list entities"})
		return
	}

	c.JSON(http.StatusOK, model.SearchResults{
		Results:    refs,
		Limit:      nextPagination.Limit,
		NextOffset: nextPagination.Offset,
	})
}

const (
	defaultLimit = "50"
)

func processListParams(c *gin.Context) ([]store.Filter, store.Ordering, store.Pagination, error) {
	filters := []store.Filter{}
	namespace := c.Query("namespace")
	if namespace != "" {
		filters = append(filters, store.Filter{
			Key:   "entity.namespace",
			Value: namespace,
		})
	}
	name := c.Query("name")
	if name != "" {
		filters = append(filters, store.Filter{
			Key:   "entity.name",
			Value: name,
		})
	}

	ordering := store.Ordering{}
	pagination := store.Pagination{}

	orderBy := c.Query("orderBy")
	if orderBy != "" {
		switch orderBy {
		case "namespace":
			ordering.OrderBy = store.OrderByNamespace
		case "name":
			ordering.OrderBy = store.OrderByName
		default:
			return filters, ordering, pagination, fmt.Errorf("invalid orderBy %s", orderBy)
		}
		descending := c.Query("descending")
		if descending == "true" {
			ordering.Descending = true
		}
	}

	limit := c.Query("limit")
	if limit == "" {
		limit = defaultLimit
	}
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt <= 0 {
			return filters, ordering, pagination, fmt.Errorf("invalid limit %s", limit)
		}
		pagination.Limit = limitInt
		offset := c.Query("offset")
		if offset != "" {
			offsetInt, err := strconv.Atoi(offset)
			if err != nil || offsetInt < 0 {
				return filters, ordering, pagination, fmt.Errorf("invalid offset %s", offset)
			}
			pagination.Offset = offsetInt
		}
	}

	return filters, ordering, pagination, nil
}
