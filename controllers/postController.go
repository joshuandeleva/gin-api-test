package controllers

import (
	"crud-gin/initializers"
	"crud-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data from requet body
	var body struct {
		Body string 
		Title string 
	}

	c.Bind(&body)

	// create post

	post := models.Post{Title: body.Title, Body: body.Body}

	// save post  form the initializers get the db instance

	result  := initializers.DB.Create(&post)

	// return response
	 if result.Error != nil {
		c.Status(400)
		return
	 }
	 // successful creation

	 c.JSON(http.StatusCreated, gin.H{"post": post})


}

func PostsIndex(c *gin.Context) {
	// get all posts
	var posts []models.Post

	 initializers.DB.Find(&posts) // find all posts

	 c.JSON(http.StatusOK, gin.H{"posts": posts}) // return all posts

}

func UpdatePost(c *gin.Context) {
	// get post id
	id := c.Param("id")
	// get post
	var post models.Post
	initializers.DB.First(&post, id) // find post by id returns only one post
	// check if post exists
	if post.ID == 0 {
		c.Status(http.StatusNotFound) // post not found
	}
	// get data from request body
	var body struct {
		Body string
		Title string
	}
	c.Bind(&body) // bind data to body
	// update post
	post.Body = body.Body
	post.Title = body.Title
	result := initializers.DB.Save(&post) // save post
	// check for errors
	if result.Error != nil {
		c.Status(http.StatusBadRequest) // bad request
		return
	}
	// return response
	c.JSON(http.StatusOK, gin.H{"post": post}) // return updated post

}

func DeletePost(c *gin.Context) {
	// get post id as query
	id := c.Query("id")

	// GET THE POST
	var post models.Post

	result := initializers.DB.First(&post, id) // find post by id

	// check if post exists
	if result.Error != nil {
		c.Status(http.StatusNotFound) // post not found
		return
	}

	// delete post
	newResult := initializers.DB.Delete(&post)

	// check error deleting post
	if newResult.Error != nil {
		c.Status(http.StatusBadRequest) // bad request
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"}) // return deleted post
}


func GetPost(c *gin.Context) {
	// get post id
	id := c.Query("id")

	var post models.Post
	// find post by id
	result := initializers.DB.First(&post, id) // find post by id

	// check if post exists

	if result.Error != nil {
		c.Status(http.StatusNotFound) // post not found
	}
	// return response
	c.JSON(http.StatusOK, gin.H{"post": post}) // return post
}