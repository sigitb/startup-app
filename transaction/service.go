package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
	"strconv"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService  payment.Service
}

type Service interface {
	GetTransactionsByCampaignId(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserId(userId int)([]Transaction,error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProsesPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{
		repository,
		campaignRepository,
		paymentService,
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

func (s *service) GetTransactionsByUserId(userId int)([]Transaction,error) {
	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return []Transaction{},err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}

	campaign, err := s.campaignRepository.FindById(input.CampaignId)
	if err != nil {
		return transaction, err
	}

	if campaign.Id == 0 {
		return transaction, errors.New("campaign not found")
	}

	transaction.CampaignId = input.CampaignId
	transaction.Amount = input.Amount
	transaction.UserId = input.User.Id
	transaction.Status = "pending"

	
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		Id: newTransaction.Id,
		Amount: newTransaction.Amount,
	}
	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}
	newTransaction.PaymentUrl = paymentUrl

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) ProsesPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction,err := s.repository.GetById(transaction_id)
	if err != nil {
		return err
	}
	
	if(input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept"){
		transaction.Status = "paid";
	}else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindById(transaction.CampaignId)
	if err != nil {
		return err
	}
	if(updatedTransaction.Status == "paid"){
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount
		_,err := s.campaignRepository.UpdateCampaign(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}