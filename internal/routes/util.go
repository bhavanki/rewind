package routes

import (
	"fmt"
	// "io"
	// "log/slog"
	"net/http"

	"github.com/bhavanki/rewind/pkg/model"
	"github.com/gin-gonic/gin"
)

func expectedEntityRef(c *gin.Context) model.EntityRef {
	kind := c.Param("kind")
	namespace := c.Param("namespace")
	name := c.Param("name")
	return model.EntityRef{
		Kind:      kind,
		Namespace: namespace,
		Name:      name,
	}
}

func verifyEntityRef(c *gin.Context, expected model.EntityRef, actual model.EntityRef) bool {
	if expected != actual {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("expected entity ref %s, got %s", expected, actual)})
		return false
	}
	return true
}

// func logRequestBody(c *gin.Context) {
// 	b, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		slog.Debug("failed to log request body", "error", err.Error())
// 	}
// 	slog.Debug(string(b))
// }
