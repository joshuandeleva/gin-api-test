package main

import (
	"crud-gin/controllers"
	"crud-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVaraibles()
	initializers.ConnectToDB()
}


func main(){
	// create a new gin instance

	r := gin.Default() // THIS COMES WITH Default logger and middleware

	// define a handler

	r.POST("/api/v1/create-post" , controllers.PostsCreate)
	r.GET("/api/v1/get-all-post" , controllers.PostsIndex)
	r.POST("/api/v1/get-post/:id" , controllers.UpdatePost)
	r.DELETE("/api/v1/delete-post" , controllers.DeletePost)
	r.GET("/api/v1/get-post" , controllers.GetPost)


	// start the server

	r.Run()
}