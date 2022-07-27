package main

import (
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
	
	userService.SaveAvatar(4, "images/1-profile.png")

	userHendler := hendler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users",userHendler.RegisterUser)
	api.POST("/sessions",userHendler.Login)
	api.POST("/email_checker",userHendler.CheckEmailAvability)
	api.POST("/avatar",userHendler.UploadAvatar)



	router.Run()
}

