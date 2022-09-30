package hendler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHendler struct {
	service transaction.Service
}

func NewTransactionHendler(service transaction.Service) *transactionHendler {
	return &transactionHendler{service}
}

func (h *transactionHendler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		respone := helper.ApiRespone("Failed to get campaign transaction", http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignId(input)
	if err != nil {
		respone := helper.ApiRespone("Failed to get campaign transactions", http.StatusBadRequest,"error", err.Error())
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	respone := helper.ApiRespone("Campaign transactions", http.StatusOK,"success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, respone)
}

func (h *transactionHendler) GetUserTransactions(c *gin.Context){
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id

	transactions,err := h.service.GetTransactionsByUserId(userId)
	if err != nil {
		respone := helper.ApiRespone("Failed to get user transactions", http.StatusBadRequest,"error", nil)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	respone := helper.ApiRespone("Campaign transactions", http.StatusOK,"success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, respone)
}