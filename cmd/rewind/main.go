package main

import (
	"github.com/bhavanki/rewind/internal/routes"
	"github.com/bhavanki/rewind/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	// store, err := store.NewSqliteStore("file::memory:?cache=shared")
	store, err := store.NewSqliteStore("file::memory:")
	if err != nil {
		panic(err)
	}

	routes.SetupRoutes(r, store)

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
