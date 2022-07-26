package main

import "fmt"

func main() {
	baseURL := "https://sandbox-api.craftgate.io"
	apiKey := "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE"
	secretKey := "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz"
	Craftgate := CraftgateClient(apiKey, secretKey, baseURL)
	res, _ := Craftgate.Installment.SearchInstallments(SearchInstallmentRequest{BinNumber: "487074", Price: 100.00})
	fmt.Println(res)
}
