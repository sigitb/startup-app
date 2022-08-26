package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{
		repository,
		campaignRepository,
	}
}

func (s *service) GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign,err := s.campaignRepository.FindById(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.Id{
		return []Transaction{}, errors.New("Not an owner of campaign")
	}

	transactions, err := s.repository.GetByCampaignId(input.Id)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
