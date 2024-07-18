package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"wallet-wise/src/models"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	_ "modernc.org/sqlite"
)

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Jokes handler not implemented yet",
	})
}

func main() {

	println("blah")

	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	categories, err := models.Categories().All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	println(categories)

	os.Exit(1)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Start and run the server
	router.Run(":3000")
}
