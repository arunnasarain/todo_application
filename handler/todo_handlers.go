package handler

import (
	"github.com/arunnasarain/todo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type TodoHandler struct {
	DB *gorm.DB
}

type todoRequest struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h TodoHandler) GetTodos(c *gin.Context) {
	var todos []models.TODO
	if res := h.DB.Find(&todos); res.Error != nil {
		c.AbortWithError(http.StatusNotFound, res.Error)
	}
	c.JSON(200, &todos)
}

func (h TodoHandler) CreateTodo(c *gin.Context) {
	var todo todoRequest

	if err := c.BindJSON(&todo); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, "TODO added")
}

func (h TodoHandler) GetTodoByID(c *gin.Context) {
	// Implement handler for GET /todos/:id
}

func (h TodoHandler) UpdateTodo(c *gin.Context) {
	// Implement handler for PUT /todos/:id
}

func (h TodoHandler) DeleteTodo(c *gin.Context) {
	// Implement handler for DELETE /todos/:id
}
