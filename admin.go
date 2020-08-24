package main

import (
	EComApp "github.com/codedv8/go-ecom-app"
)

type ADMIN struct {
	App     *EComApp.Application
	Message string
}

var adminList []ADMIN

// Exports
func SysInit(app *EComApp.Application) error {
	admin := &ADMIN{
		App:     app,
		Message: "Welcome to the CMS plugin",
	}
	admin.SysInit(app)

	adminList = append(adminList, *admin)

	return nil
}

func Init(app *EComApp.Application) error {
	for _, admin := range adminList {
		admin.Init(app)
	}

	return nil
}

func Done(app *EComApp.Application) error {
	for _, admin := range adminList {
		admin.Done(app)
	}

	return nil
}

var admin ADMIN
