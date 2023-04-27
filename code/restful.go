package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping/:id", func(c *gin.Context) {
		getId := c.Param("id")
		user := c.DefaultQuery("user", "test")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"message": "pong",
			"id":      getId,
			"user":    user,
			"pwd":     pwd,
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "test")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	// DELETE 和 get基本一致
	r.DELETE("/ping/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//PUT和POST基本一致
	r.PUT("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8008") // listen and serve on 0.0.0.0:8080
}
