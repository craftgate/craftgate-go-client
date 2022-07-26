package main

type Pos struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	BankID int    `json:"bankId"`
}

type Payout struct {
	PaidPrice                     float64 `json:"paidPrice"`
	Currency                      string  `json:"currency"`
	MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
	SubMerchantMemberPayoutAmount float64 `json:"subMerchantMemberPayoutAmount"`
}

type PaymentTransactions struct {
	ID                            int     `json:"id"`
	ExternalID                    string  `json:"externalId"`
	Name                          string  `json:"name"`
	TransactionStatus             string  `json:"transactionStatus"`
	BlockageResolvedDate          string  `json:"blockageResolvedDate"`
	Price                         float64 `json:"price"`
	PaidPrice                     float64 `json:"paidPrice"`
	WalletPrice                   int     `json:"walletPrice"`
	MerchantCommissionRate        int     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
	SubMerchantMemberID           int     `json:"subMerchantMemberId"`
	SubMerchantMemberPrice        float64 `json:"subMerchantMemberPrice"`
	SubMerchantMemberPayoutRate   int     `json:"subMerchantMemberPayoutRate"`
	SubMerchantMemberPayoutAmount float64 `json:"subMerchantMemberPayoutAmount"`
	Payout                        Payout  `json:"payout"`
}

type Payment struct {
	ID                           int                   `json:"id"`
	PaymentPhase                 string                `json:"paymentPhase"`
	CreatedDate                  string                `json:"createdDate"`
	Price                        int                   `json:"price"`
	PaidPrice                    float64               `json:"paidPrice"`
	WalletPrice                  int                   `json:"walletPrice"`
	Installment                  int                   `json:"installment"`
	Currency                     string                `json:"currency"`
	BuyerMemberID                int                   `json:"buyerMemberId"`
	PaymentType                  string                `json:"paymentType"`
	PaymentGroup                 string                `json:"paymentGroup"`
	ConversationID               string                `json:"conversationId"`
	PaymentStatus                string                `json:"paymentStatus"`
	MerchantCommissionRate       int                   `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64               `json:"merchantCommissionRateAmount"`
	BankCommissionRate           int                   `json:"bankCommissionRate"`
	BankCommissionRateAmount     float64               `json:"bankCommissionRateAmount"`
	BinNumber                    string                `json:"binNumber"`
	LastFourDigits               string                `json:"lastFourDigits"`
	CardHolderName               string                `json:"cardHolderName"`
	BankCardHolderName           string                `json:"bankCardHolderName"`
	CardType                     string                `json:"cardType"`
	CardAssociation              string                `json:"cardAssociation"`
	CardBrand                    string                `json:"cardBrand"`
	Pos                          Pos                   `json:"pos"`
	PaymentTransactions          []PaymentTransactions `json:"paymentTransactions"`
}

/*
func (c *Client) Pay(ctx context.Context, paymentParams PaymentParams) (*PaymentResponse, error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/faces?limit=%d&page=%d", c.BaseURL, limit, page), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := FacesList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
*/
