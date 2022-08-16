package hendler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
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

	respone := helper.ApiRespone("Detail Campaign",http.StatusOK, "success", campaign.FormatDetailCampaign(campaigns))
	c.JSON(http.StatusOK, respone)
}