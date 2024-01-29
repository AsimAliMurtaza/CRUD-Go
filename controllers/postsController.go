package controllers

import (
	"github.com/asim/go-crud/initializers"
	"github.com/asim/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	c.JSON(200, gin.H{
		"post": post,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get id off of url
	var post models.Post
	id := c.Param("id")
	//get data off of url
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	//find post by id
	initializers.DB.First(&post, id)
	//update post
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body})
	//return updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)
	initializers.DB.Delete(&post)
	c.JSON(200, gin.H{
		"post": post,
	})
}
