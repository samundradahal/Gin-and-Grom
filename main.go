package main

import (
	"go-gin/controllers"  // recently added
	"go-gin/initializers" // recenlty added

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB() // recently added
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate) // recently added
	r.GET("/posts", controllers.PostsFindAll)

	r.Run()
}
