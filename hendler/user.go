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
		errorMassage := gin.H{"error":errors}

		respone := helper.ApiRespone("register account failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	newUser , err := h.userService.RegisterUser(input)	
	if err != nil{
		errors := helper.FormatterValidationError(err)
		errorMassage := gin.H{"error":errors}
		respone := helper.ApiRespone("register account failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	formatter := user.FormatUser(newUser, "token")
	respone := helper.ApiRespone("Account has been regitered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK,respone)
}