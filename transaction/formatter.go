package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0{
		return []CampaignTransactionFormatter{}
	}

	var transactionFormatter []CampaignTransactionFormatter

	for _,transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

type UserTransactionFormater struct{
	Id        int    `json:"id"`
	Amount    int    `json:"amount"`
	Status      string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Campaign CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct{
	Name      string `json:"name"`
	ImageUrl  string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormater {
	formatter := UserTransactionFormater{}

	formatter.Id = transaction.Id
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormater := CampaignFormatter{}

	campaignFormater.Name = transaction.Campaign.Name
	campaignFormater.ImageUrl = ""

	if len(transaction.Campaign.CampaignImage) > 0 {
		campaignFormater.ImageUrl = transaction.Campaign.CampaignImage[0].FileName
	}

	formatter.Campaign = campaignFormater
	return formatter;
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormater {
	if len(transactions) == 0{
		return []UserTransactionFormater{}
	}

	var transactionFormatter []UserTransactionFormater

	for _,transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}