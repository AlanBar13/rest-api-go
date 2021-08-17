package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"Title" binding:"required"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed"`
	Due         string `json:"Due"`
}

var todos = []todo{
	{ID: "1", Title: "Buy groceries", Description: "Buy ketchup, and Meat", Completed: false, Due: time.Now().String()},
	{ID: "2", Title: "Go get a Haircut", Description: "Haircut price $20", Completed: false, Due: time.Now().String()},
	{ID: "3", Title: "Take the dog out", Description: "", Completed: false, Due: time.Now().String()},
	{ID: "4", Title: "Build REST API with GO", Description: "See tutotrial", Completed: false, Due: time.Now().String()},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to REST API",
		})
	})
	router.Run(":5000")
}
