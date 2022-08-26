package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignId(campaignId int)([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignId(campaignId int)([]Transaction, error)  {
	var transaction []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("id desc").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}