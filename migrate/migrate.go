package main

import (
	"crud-gin/initializers"
	"crud-gin/models"
)

func init(){
	initializers.LoadEnvVaraibles()
	initializers.ConnectToDB()

}


func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}