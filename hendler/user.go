package hendler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	var input user.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatterValidationError(err)
		errorMassage := gin.H{"errors":errors}

		respone := helper.ApiRespone("register account failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	newUser , err := h.userService.RegisterUser(input)	
	if err != nil{
		errors := helper.FormatterValidationError(err)
		errorMassage := gin.H{"errors":errors}
		respone := helper.ApiRespone("register account failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	formatter := user.FormatUser(newUser, "token")
	respone := helper.ApiRespone("Account has been regitered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK,respone)
}

func (h *userHandler) Login(c *gin.Context){
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	
	if err != nil{
		errors := helper.FormatterValidationError(err)
		errorMassage := gin.H{"errors":errors}
		respone := helper.ApiRespone("Login failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMassage := gin.H{"errors":err.Error()}
		respone := helper.ApiRespone("Login failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	} 

	formatter := user.FormatUser(loggedinUser, "token")
	respone := helper.ApiRespone("loggedin successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK,respone)
}

func (h *userHandler) CheckEmailAvability(c *gin.Context)  {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	
	if err != nil{
		errors := helper.FormatterValidationError(err)
		errorMassage := gin.H{"errors":errors}
		respone := helper.ApiRespone("Email checking failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	
	if err != nil {
		errorMassage := gin.H{"errors":"Server Error"}
		respone := helper.ApiRespone("Login failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, respone)
		return
	} 

	data := gin.H{
		"is_available" : isEmailAvailable,
	}
	metaMessage := "Email has been registered"

	if isEmailAvailable{
		metaMessage = "Email is Availabe"
	}
	respone := helper.ApiRespone(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK,respone)
}