package main

import (
	"github.com/arunnasarain/todo/db"
	"github.com/arunnasarain/todo/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("./configs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	//dbUrl := os.Getenv("DATABASE_URL")
	dbUrl := viper.Get("DB_URL").(string)
	// Initialize Gin router
	router := gin.Default()

	// Initialise Db and add dummy TODOs
	dbHandler := db.Init(dbUrl)
	db.InitialData(dbHandler)

	h := &handler.TodoHandler{
		DB: dbHandler,
	}
	// Routes
	router.GET("/todos", h.GetTodos)
	router.POST("/todos", h.CreateTodo)
	router.GET("/todos/:id", h.GetTodoByID)
	router.PUT("/todos/:id", h.UpdateTodo)
	router.DELETE("/todos/:id", h.DeleteTodo)

	// Start server
	router.Run(port)
}
