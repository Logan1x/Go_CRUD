package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/logan1x/go_crud/initializers"
	"github.com/logan1x/go_crud/models"
)

func PostsCreate(c *gin.Context) {

	//  get data from request
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//  create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get data from request
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get post id from url
	id := c.Param("id")

	// find post
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get post id from url
	id := c.Param("id")

	// get post data from request
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find post
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// update post

	result = initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get post id from url
	id := c.Param("id")

	// delete the post
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// respond back with 200

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
