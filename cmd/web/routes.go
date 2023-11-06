package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(router *gin.Engine) {

	// Define the route handlers
	router.GET("/", app.index)
	router.GET("/login", app.login)
	router.POST("/login", app.login)
	router.GET("/signup", app.signup)
	router.POST("/signup", app.signup)
	router.GET("/homepage", app.homepage)

	// Define the static file server middleware
	router.Static("/static", "./ui/static/")

}
