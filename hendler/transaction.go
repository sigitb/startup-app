package hendler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
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

func (h *transactionHendler) CreateTransaction(c *gin.Context)  {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		fmt.Println("err")
		errors := helper.FormatterValidationError(err)
		errorMassage := gin.H{"errors":errors}

		respone := helper.ApiRespone("Create Transaction failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		respone := helper.ApiRespone("Failed to create transaction", http.StatusBadRequest,"error", err.Error())
		c.JSON(http.StatusBadRequest, respone)
		return
	}
	respone := helper.ApiRespone("Transaction create successfully", http.StatusOK,"success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusBadRequest, respone)
}

func (h *transactionHendler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ApiRespone("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	err = h.service.ProsesPayment(input)
	if err != nil {
		response := helper.ApiRespone("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	c.JSON(http.StatusOK,input)
}