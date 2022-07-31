package main

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/hendler"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	api.POST("/avatar",authMiddleware(authService,userService),userHendler.UploadAvatar)



	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context){
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader, "Bearer"){
			respone := helper.ApiRespone("Unauthorized", http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,respone)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2{
			tokenString = arrayToken[1]
		}
	
		token , err := authService.ValidateToken(tokenString)
		if err != nil {
			respone := helper.ApiRespone("Unauthorized", http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,respone)
			return
		}

		claim, oke := token.Claims.(jwt.MapClaims)

		if !oke || !token.Valid{
			respone := helper.ApiRespone("Unauthorized", http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,respone)
			return
		}
		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserByid(userId)
		if err != nil {
			respone := helper.ApiRespone("Unauthorized", http.StatusUnauthorized,"error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,respone)
			return
		}
		c.Set("currentUser", user)
	}
}


