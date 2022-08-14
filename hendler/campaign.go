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

	respone := helper.ApiRespone("List of campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK,respone)
}