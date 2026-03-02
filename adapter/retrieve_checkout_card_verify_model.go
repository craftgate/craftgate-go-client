package adapter

type RetrieveCheckoutCardVerifyResponse struct {
    Token *string             `json:"token"`
    Card  *StoredCardResponse `json:"card"`
}
