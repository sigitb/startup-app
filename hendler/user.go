package hendler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token ,err := h.authService.GenerateToken(newUser.Id)
	if err != nil{
		respone := helper.ApiRespone("register account failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	formatter := user.FormatUser(newUser, token)
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
	token ,err := h.authService.GenerateToken(loggedinUser.Id)
	if err != nil{
		respone := helper.ApiRespone("Login failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)
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

func (s *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil{
		data := gin.H{
			"is_uploaded" : false, 
		}
		respone := helper.ApiRespone("File to upload avatar image", http.StatusBadRequest,"error",data)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id
	path := fmt.Sprintf("images/%d-%s",userId,file.Filename) 
	err = c.SaveUploadedFile(file, path)
	if err != nil{
		data := gin.H{
			"is_uploaded" : false, 
		}
		respone := helper.ApiRespone("File to upload avatar image", http.StatusBadRequest,"error",data)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	
	_ ,err = s.userService.SaveAvatar(userId, path) 
	if err != nil{
		data := gin.H{
			"is_uploaded" : false, 
		}
		respone := helper.ApiRespone("File to upload avatar image", http.StatusBadRequest,"error",data)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	data := gin.H{
		"is_uploaded" :true,
	}
	respone := helper.ApiRespone("Avatar successfully upload", http.StatusOK,"success",data)
	c.JSON(http.StatusBadRequest, respone)
	return
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	formatter := user.FormatUser(currentUser, "")
	response := helper.ApiRespone("Success fetch user data", http.StatusOK, "success",formatter)
	c.JSON(http.StatusOK, response)
}