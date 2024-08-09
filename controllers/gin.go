package controllers

import (
	"club.scimatic/bapi/database"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post *database.Post
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(post)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
	return
}

func ReadPost(c *gin.Context) {
	var post database.Post
	id := c.Param("id")
	res := database.DB.Find(&post, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
	return
}

func ReadPosts(c *gin.Context) {
	var posts []database.Post
	res := database.DB.Find(&posts)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
	return
}

func UpdatePost(c *gin.Context) {
	var post database.Post
	id := c.Param("id")
	err := c.ShouldBind(&post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updatePost database.Post
	res := database.DB.Model(&updatePost).Where("id = ?", id).Updates(post)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "post not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
	return
}

func DeletePost(c *gin.Context) {
	var post database.Post
	id := c.Param("id")
	res := database.DB.Find(&post, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}
	database.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "post deleted successfully",
	})
	return
}
