package routes

import (
	"net/http"

	"github.com/bhavanki/rewind/internal/store"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, store store.Store) {
	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/v1/:kind/:namespace/:name", withStore(store, ReadEntity))
	r.POST("/api/v1/:kind/:namespace/:name", withStore(store, CreateEntity))
	r.PUT("/api/v1/:kind/:namespace/:name", withStore(store, UpdateEntity))
	r.DELETE("/api/v1/:kind/:namespace/:name", withStore(store, DeleteEntity))
}

type storeHandlerFunc func(*gin.Context, store.Store)

func withStore(store store.Store, f storeHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c, store)
	}
}
