package main

import (
	"bwastartup/auth"
	"bwastartup/hendler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()


	userHendler := hendler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users",userHendler.RegisterUser)
	api.POST("/sessions",userHendler.Login)
	api.POST("/email_checker",userHendler.CheckEmailAvability)
	api.POST("/avatar",userHendler.UploadAvatar)



	router.Run()
}

