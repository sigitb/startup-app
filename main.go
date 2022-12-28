package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/hendler"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
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
	campaignRespository := campaign.NewRepository(db)
	transactionRespository := transaction.NewRepository(db)

	campaignService := campaign.NewService(campaignRespository)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRespository, campaignRespository, paymentService)

	userHendler := hendler.NewUserHandler(userService, authService)
	campaignHendler := hendler.NewCampaignHendler(campaignService)
	transactionHendler := hendler.NewTransactionHendler(transactionService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users",userHendler.RegisterUser)
	api.POST("/sessions",userHendler.Login)
	api.POST("/email_checker",userHendler.CheckEmailAvability)
	api.POST("/avatar",authMiddleware(authService,userService),userHendler.UploadAvatar)
	
	// campaign
	api.GET("/campaigns", campaignHendler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHendler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService,userService),campaignHendler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService,userService),campaignHendler.UpdateCampaign)
	api.POST("/campaign-image", authMiddleware(authService,userService),campaignHendler.CreateCampaignImage)

	// transaction
	api.GET("/campaign/:id/transaction", authMiddleware(authService,userService), transactionHendler.GetCampaignTransaction)
	api.GET("/transaction", authMiddleware(authService,userService), transactionHendler.GetUserTransactions)
	api.POST("/transaction", authMiddleware(authService,userService), transactionHendler.CreateTransaction)

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


