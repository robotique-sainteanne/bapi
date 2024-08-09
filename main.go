package main

import (
	"fmt"

	"club.scimatic/bapi/controllers"
	"club.scimatic/bapi/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/posts/:id", controllers.ReadPost)
	r.GET("/posts", controllers.ReadPosts)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run(":5000")
}
