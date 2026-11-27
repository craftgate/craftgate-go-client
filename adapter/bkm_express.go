package adapter

import (
	"context"
	"net/http"
)

type BkmExpress struct {
	Client *Client
}

func (api *BkmExpress) GenerateToken(ctx context.Context, request BkmExpressGenerateTokenRequest) (*BkmExpressGenerateTokenResponse, error) {
    newRequest, err := api.Client.NewRequest(ctx, http.MethodPost, "/payment/v2/bkm-express/generate-token", request)
    if err != nil {
        return nil, err
    }

    response := &Response[BkmExpressGenerateTokenResponse]{}
    err = api.Client.Do(ctx, newRequest, response)
    if err != nil {
        return nil, err
    }

    return response.Data, nil
}
