package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func routes(app *application) http.Handler {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://yourdomain.com", "http://*"}, // Allowed frontends
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	api := r.Group("/api/v1")
	api.GET("/healthcheck", app.healthCheckHandler)
	api.GET("/", app.createTodoHandler)

	// todo routes
	todo := api.Group("/todos")
	todo.POST("/", app.createTodoHandler)
	todo.GET("/", app.createTodoHandler)

	return r
}

func (app *application) healthCheckHandler(c *gin.Context) {
	health := healthCheckResponse{
		Msg:  "API is up and running",
		Code: http.StatusOK,
	}
	c.JSON(http.StatusOK, gin.H{"status": health})
}
