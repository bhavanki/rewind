package routes

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bhavanki/rewind/internal/store"
	"github.com/bhavanki/rewind/pkg/model"
	"github.com/gin-gonic/gin"
)

func CreateEntity(c *gin.Context, store store.Store) {
	expectedEntityRef := expectedEntityRef(c)
	kind := expectedEntityRef.Kind

	switch kind {
	case "component":
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
	case "component":
		component, err := store.ReadComponent(expectedEntityRef)
		if err != nil {
			slog.Error("failed to read component", "entityRef", expectedEntityRef.String(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read component"})
			return
		}
		c.YAML(http.StatusOK, component)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("unsupported kind %s", kind)})
		return
	}
}
