package main

import (
	"log"

	EComApp "github.com/codedv8/go-ecom-app"
	EComStructs "github.com/codedv8/go-ecom-structs"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

// SysInit - Pre initialization of this object
func (admin *ADMIN) SysInit(app *EComApp.Application) {
	// Create apiRouter
	adminRouter := app.Router.Group("/admin")
	// Create basic auth object
	basicAuth := app.UseBasicAuth("Admin")
	// Add basic auth to adminRouter
	adminRouter.Use(basicAuth.Use())
	// Add handlers
	adminRouter.Handle("GET", "/", func(c *gin.Context) {
		log.Printf("Welcome to this admin %+v\n", c)
		c.String(200, "Welcome to this admin")
	})
}

// Init - Initialization of this object
func (admin *ADMIN) Init(app *EComApp.Application) {
	app.ListenToHook("ROUTER_WILDCARD", func(payload interface{}) (bool, error) {
		switch c := payload.(type) {
		case *EComStructs.RouterWildcard:
			path := c.Context.Request.URL.Path
			if len(path) < 7 {
				return true, nil
			}
			if path[:7] == "/admin/" {
				c.Context.String(200, c.Context.Request.Method+" "+path[7:]+"\n")
				c.Context.String(200, "Plugin Admin responded to this call: "+path)
				return false, nil
			}
			return true, nil
		}
		return true, nil
	})
}
