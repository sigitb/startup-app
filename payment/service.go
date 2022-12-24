package payment

import (
	"bwastartup/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
    midclient.ServerKey = "SB-Mid-server-E3oIdm2vBBwKZUKY60OdbtZO"
    midclient.ClientKey = "SB-Mid-client-xVWt5pbNQ3LaImTL"
    midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
        Client: midclient,
    }
	 snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.Id),
			GrossAmt: int64(transaction.Amount),
		},
	 }
	 snapTokenResp, err := snapGateway.GetToken(snapReq)
	 
	 if err != nil {
		return "", err
	 }
	 return snapTokenResp.RedirectURL,nil
}