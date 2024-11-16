package main

import (
	"chitchatv3/data"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	var DB *gorm.DB
	// dsn := "host=/var/run/postgresql user=postgres password=892289 dbname=chitchat sslmode=disable"
	dsn := "lumia:892289@tcp(127.0.0.1:3306)/chitchat?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&data.User{}, &data.Session{}, &data.Thread{}, &data.Post{})
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

func main() {
	p("ChitChat Blog", version(), "running at", config.Address)
	//初始化数据库
	DB := initDB()
	data.SetDB(DB)
	// 初始化Gin引擎
	r := gin.Default()

	// 处理静态资源
	r.Static("/static", config.Static)

	// 定义路由和处理函数
	//defined in route_main.go
	r.GET("/", index)
	r.GET("/err", err)

	// defined in route_auth.go
	r.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/signup", signup)
	r.POST("/signup_account", signupAccount)
	r.POST("/authenticate", authenticate)

	// defined in route_thread.go
	r.GET("/thread/new", newThread)
	r.POST("/thread/create", createThread)
	r.POST("/thread/post", postThread)
	r.GET("/thread/read", readThread)

	// 启动Gin服务器
	s := &http.Server{
		Addr:           config.Address,
		Handler:        r,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
