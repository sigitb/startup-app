package hendler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHendler struct {
	service campaign.Service
}

func NewCampaignHendler(s campaign.Service) *campaignHendler{
	return &campaignHendler{s}
}

func (h *campaignHendler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id")) 
	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		respone := helper.ApiRespone("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, respone)
	}  

	respone := helper.ApiRespone("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK,respone)
}

func (h *campaignHendler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		respone := helper.ApiRespone("Failed to get detail of campaign", http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	campaigns, err := h.service.GetCampaignById(input)
	if err != nil {
		respone := helper.ApiRespone("Failed to get detail of campaign", http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	if campaigns.Id == 0{
		respone := helper.ApiRespone("Detail of campaign not found", http.StatusNotFound,"error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	respone := helper.ApiRespone("Detail Campaign",http.StatusOK, "success", campaign.FormatDetailCampaign(campaigns))
	c.JSON(http.StatusOK, respone)
}

func (h *campaignHendler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	// err := c.ShouldBindJSON(&input)
	// if err != nil{
	// 	errors := helper.FormatterValidationError(err)
	// 	errorMassage := gin.H{"errors":errors}

	// 	respone := helper.ApiRespone("Create Campaign failed", http.StatusUnprocessableEntity, "error", errorMassage)
	// 	c.JSON(http.StatusBadRequest, respone)
	// 	return
	// }

	// validation validate in input 

	validate := input.CreateCampaign()
	
	if validate != nil {
		 validator := make(map[string]interface{})
			for _, v := range validate {
				validator[v.Key] = v.Message
			}
		respone := helper.ApiRespone("Create Campaign failed", http.StatusBadRequest, "error", validator)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil{
		respone := helper.ApiRespone("Create Campaign failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	respone := helper.ApiRespone("Create Campaign successfuly", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, respone)
}

func (h *campaignHendler) UpdateCampaign(c *gin.Context){
	var inputId campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputId)

	if err != nil {
		respone := helper.ApiRespone("Failed to update campaign", http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	var inputData campaign.CreateCampaignInput

	errs := c.ShouldBindJSON(&inputData)
	if errs != nil{
		errors := helper.FormatterValidationError(errs)
		errorMassage := gin.H{"errors":errors}

		respone := helper.ApiRespone("Update Campaign failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {
		respone := helper.ApiRespone("Failed to update campaign", http.StatusBadRequest,"error", err.Error())
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	respone := helper.ApiRespone("Update Campaign successfuly", http.StatusOK, "success", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, respone)
}

func (h *campaignHendler) CreateCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil{
		errors := helper.FormatterValidationError(err)
		errorMassage := gin.H{"errors":errors}

		respone := helper.ApiRespone("Create Campaign failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id
	input.User = currentUser
	file, err := c.FormFile("file")
	if err != nil{
		data := gin.H{
			"is_uploaded" : false, 
		}
		respone := helper.ApiRespone("Failed to upload campaign image", http.StatusBadRequest,"error",data)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	path := fmt.Sprintf("images/%d-%s",userId,file.Filename) 
	err = c.SaveUploadedFile(file, path)
	if err != nil{
		data := gin.H{
			"is_uploaded" : false, 
		}
		respone := helper.ApiRespone("Failed to upload campaign image", http.StatusBadRequest,"error",data)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	_ ,err = h.service.SaveCampaignImage(input, path) 
	if err != nil{
		data := gin.H{
			"is_uploaded" : false, 
		}
		respone := helper.ApiRespone("Failed to upload campaign image", http.StatusBadRequest,"error",data)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	data := gin.H{
		"is_uploaded" :true,
	}
	respone := helper.ApiRespone("Campaign Image successfully uploaded", http.StatusOK,"success",data)
	c.JSON(http.StatusBadRequest, respone)
	return
}