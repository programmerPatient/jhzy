/*
 * @Description:
 * @Author: mali
 * @Date: 2022-09-05 13:47:43
 * @LastEditTime: 2022-11-09 11:16:52
 * @LastEditors: VSCode
 * @Reference:
 */
package main

import (
	"file/core/database"
	"file/model"
	"file/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {

	router := gin.Default()
	// 设置静态文件路径
	router.Static("/static", "./static")
	// 注意此处的导入路径
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/index", service.FileS.Index)
	// 路由：文件下载
	router.GET("/download/:filename", service.FileS.Download)
	// 路由：文件上传
	router.POST("/upload", service.FileS.UploadFile)
	router.GET("/checkin/index", service.CheckinS.Index)
	router.POST("/checkin/click", service.CheckinS.Click)
	router.GET("/content/index", service.ContentS.Index)
	router.POST("/content/up", service.ContentS.Up)
	setupDB()
	//迁移
	model.Init()
	router.Run(":8998")
}

// 数据库连接建立
func setupDB() {
	var db gorm.Dialector
	db = mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			//"zmcs",
			//"kZYmGhBJApZ4XFyW",
			//"43.143.26.21",
			"root",
			"1234qwer!",
			"127.0.0.1",
			"3306",
			"zmcs",
			"utf8mb4",
		),
	})
	//使用gorm的默认日志服务
	database.Connect(db)
	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(10)
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(10) * time.Second)
}
