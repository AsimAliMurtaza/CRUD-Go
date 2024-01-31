package controllers

import (
	"github.com/asim/go-crud/initializers"
	"github.com/asim/go-crud/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {

	var body struct {
		Name     string
		Password string
		Email    string
	}
	c.Bind(&body)

	post := models.User{Name: body.Name, Password: body.Password, Email: body.Email}

	result := initializers.DB.Create(&post)
	c.JSON(200, gin.H{
		"createUser": post,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}
}

func UserIndex(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"posts": users,
	})
}

func PostsShow(c *gin.Context) {
	var post models.User
	id := c.Param("id")
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get id off of url
	var post models.User
	id := c.Param("id")
	//get data off of url
	var body struct {
		Name     string
		Password string
		Email    string
	}
	c.Bind(&body)
	//find post by id
	initializers.DB.First(&post, id)
	//update post
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&post).Updates(models.User{
		Name:     body.Name,
		Password: body.Password,
		Email:    body.Email})
	//return updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	var post models.User
	id := c.Param("id")
	initializers.DB.First(&post, id)
	initializers.DB.Delete(&post)
	c.JSON(200, gin.H{
		"post": post,
	})
}
