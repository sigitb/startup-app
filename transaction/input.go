package transaction

import "bwastartup/user"

type GetCampaignTransactionsInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}