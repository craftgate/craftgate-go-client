package adapter

import (
	"craftgate-go-client/adapter/rest"
	"craftgate-go-client/model"
	"fmt"
	"net/http"
	"time"
)

type SettlementReporting struct {
	Opts model.RequestOptions
}

type SearchPayoutCompletedTransactionsRequest struct {
	SettlementFileId int64                `schema:"settlementFileId,omitempty"`
	SettlementType   model.SettlementType `schema:"settlementType,omitempty"`
	StartDate        time.Time            `schema:"startDate,omitempty"`
	EndDate          time.Time            `schema:"endDate,omitempty"`
}

type SearchPayoutBouncedTransactionsRequest struct {
	StartDate time.Time `schema:"startDate,omitempty"`
	EndDate   time.Time `schema:"endDate,omitempty"`
}

type RetrievePayoutDetailsRequest struct {
	PayoutDetailId int64
}

type SearchPayoutCompletedTransactionsResponse struct {
	PayoutId                      int64          `json:"payoutId"`
	TransactionId                 int64          `json:"transactionId"`
	TransactionType               string         `json:"transactionType"`
	PayoutAmount                  float64        `json:"payoutAmount"`
	Currency                      model.Currency `json:"currency"`
	MerchantId                    int64          `json:"merchantId"`
	MerchantType                  string         `json:"merchantType"`
	SettlementEarningsDestination string         `json:"settlementEarningsDestination"`
	SettlementSource              string         `json:"settlementSource"`
}

type SearchPayoutBouncedTransactionsResponse struct {
	Id                int64        `json:"id"`
	Iban              string       `json:"iban"`
	CreatedDate       TimeResponse `json:"createdDate"`
	UpdatedDate       TimeResponse `json:"updatedDate"`
	PayoutId          int64        `json:"payoutId"`
	PayoutAmount      float64      `json:"payoutAmount"`
	ContactName       string       `json:"contactName"`
	ContactSurname    string       `json:"contactSurname"`
	LegalCompanyTitle string       `json:"legalCompanyTitle"`
	RowDescription    string       `json:"rowDescription"`
}

type PayoutDetailResponse struct {
	RowDescription                string                            `json:"rowDescription"`
	PayoutDate                    TimeResponse                      `json:"payoutDate"`
	Name                          string                            `json:"name"`
	Iban                          string                            `json:"iban"`
	PayoutAmount                  float64                           `json:"payoutAmount"`
	Currency                      string                            `json:"currency"`
	MerchantId                    int64                             `json:"merchantId"`
	MerchantType                  string                            `json:"merchantType"`
	SettlementEarningsDestination string                            `json:"settlementEarningsDestination"`
	SettlementSource              string                            `json:"settlementSource"`
	BounceStatus                  string                            `json:"bounceStatus"`
	PayoutTransactions            []PayoutDetailTransactionResponse `json:"payoutTransactions"`
}

type PayoutDetailTransactionResponse struct {
	TransactionId   int64   `json:"transactionId"`
	TransactionType string  `json:"transactionType"`
	PayoutAmount    float64 `json:"payoutAmount"`
}

func (api *SettlementReporting) SearchPayoutCompletedTransactions(request SearchPayoutCompletedTransactionsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/settlement-reporting/v1/settlement-file/payout-completed-transactions", api.Opts.BaseURL), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[SearchPayoutCompletedTransactionsResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *SettlementReporting) SearchPayoutBouncedTransactions(request SearchPayoutBouncedTransactionsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/settlement-reporting/v1/settlement-file/bounced-sub-merchant-rows", api.Opts.BaseURL), nil)
	req.URL.RawQuery, _ = QueryParams(request)
	res := model.Response[model.DataResponse[SearchPayoutBouncedTransactionsResponse]]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}

func (api *SettlementReporting) RetrievePayoutDetails(request RetrievePayoutDetailsRequest) (interface{}, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/settlement-reporting/v1/settlement-file/payout-details/%d", api.Opts.BaseURL, request.PayoutDetailId), nil)
	res := model.Response[PayoutDetailResponse]{}
	resErr := rest.SendRequest(req, &res, api.Opts)
	return &res, resErr
}
