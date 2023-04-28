package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PersonMessage struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,agebig28"`
	Sex  string `json:"sex" binding:"required"`
}

// 自定义验证
func agebig28(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(int) < 29 {
		return false
	}
	return true
}

func main() {

	r := gin.Default()
    
	//自定义验证注册
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("agebig28", agebig28)
	}

	r.POST("/ping", func(c *gin.Context) {
		var person PersonMessage
		err := c.ShouldBindJSON(&person)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "bind json error",
				"data":    err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success",
				"data":    person,
			})
		}

	})

	r.Run(":8008") // listen and serve on 0.0.0.0:8080
}
