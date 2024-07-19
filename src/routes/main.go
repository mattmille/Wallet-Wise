package routes

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Set the router as the default one shipped with Gin
var router = gin.Default()

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	api := router.Group("/api")
	addExpenseRoutes(api)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}

// Run will start the server
func Run() {
	getRoutes()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run(":3000")
}
