package campaign

import "bwastartup/user"

type GetCampaignDetailInput struct {
	Id int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name"  binding:"required" validate:"required"`
	ShortDescription string `json:"short_description"  binding:"required" validate:"required"`
	Description      string `json:"description" binding:"required" validate:"required"`
	GoalAmount       int    `json:"goal_amount"  binding:"required" validate:"required" label:"Jumlah Tujuan"`
	Perks            string `json:"perks"  binding:"required" validate:"required"`
	User             user.User
}