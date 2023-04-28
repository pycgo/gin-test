package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
//中间件函数1
func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我再方法前面1")
		c.Next()
		fmt.Println("我再方法后面1")
	}
}
//中间件函数2
func middle2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我再方法前面2")
		c.Next()
		fmt.Println("我再方法后面2")
	}
}

func main() {
	router := gin.Default()
	v1 := router.Group("v1").Use(middle()).Use(middle2())

	v1.GET("/test", func(c *gin.Context) {
		fmt.Println("我在v1路由中")
		// Multipart form
		c.JSON(200, gin.H{
			"success": true,
		})

	})
	router.Run(":8080")
}
