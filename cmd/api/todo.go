package main

import (
	"net/http"

	"github.com/cedrickewi/hotel-reservation/internals/services"
	"github.com/gin-gonic/gin"
)

var payload services.Todo

func (app *application) createTodoHandler(c *gin.Context) {
	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := app.store.Todo.Insert(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create todo",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"todo":    payload,
	})
}

func (app *application) listTodosHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List of todos",
	})
}
