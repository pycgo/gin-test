//参考 https://gorm.io/zh_CN/docs/
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Sex  bool
	Age  int
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(172.16.86.128:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	//自动根据结构体创建表
	db.AutoMigrate(&HelloWorld{})
	//增加数据
	//db.Create(&HelloWorld{
	//	Name: "zxx",
	//	Sex:  true,
	//	Age:  182,
	//})

	//查询单条数据

	//var hello1 HelloWorld
	//查询第一条数据 // SELECT * FROM users ORDER BY id LIMIT 1;
	//db.First(&hello1)
	//查询最后一条数据 // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//db.Last(&hello1)
	//SELECT * FROM users LIMIT 1;
	//db.Take(&hello1)
	//fmt.Println(hello1)

	//查询多条数据 find

	var findReust []HelloWorld
	////db.Find(&findReust)
	//fmt.Println(findReust)
	db.Where("name = ?", "zxx1").Find(&findReust)
	fmt.Println(findReust)

}
