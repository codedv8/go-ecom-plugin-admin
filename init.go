package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
	EComStructs "github.com/codedv8/go-ecom-structs"
	_ "github.com/gin-gonic/gin"
)

func (admin *ADMIN) SysInit(app *EComApp.Application) {

}

func (admin *ADMIN) Init(app *EComApp.Application) {
	app.ListenToHook("ROUTER_WILDCARD", func(payload interface{}) (bool, error) {
		switch c := payload.(type) {
		case *EComStructs.RouterWildcard:
			path := c.Context.Request.URL.Path
			if len(path) == 11 {
				c.Context.String(200, "Path was 11 chars in length. Plugin Admin responded to this call")
				return false, nil
			}
			return true, nil
		}
		return true, nil
	})
}
