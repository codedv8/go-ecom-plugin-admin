package main

import (
	ecomapp "github.com/codedv8/go-ecom-app"
)

// ADMIN - The main struct for this module
type ADMIN struct {
	App     *ecomapp.Application
	Message string
}

var adminList []ADMIN

// Exports

// SysInit - Pre initialization of this module
func SysInit(app *ecomapp.Application) error {
	admin := &ADMIN{
		App:     app,
		Message: "Welcome to the CMS plugin",
	}
	admin.SysInit(app)

	adminList = append(adminList, *admin)

	return nil
}

// Init - Initialization of this module
func Init(app *ecomapp.Application) error {
	for _, admin := range adminList {
		admin.Init(app)
	}

	return nil
}

// Done - Shut down of this module
func Done(app *ecomapp.Application) error {
	for _, admin := range adminList {
		admin.Done(app)
	}

	return nil
}

var admin ADMIN
