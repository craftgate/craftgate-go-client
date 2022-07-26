package main

import "fmt"

func main() {
	apiKey := "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE"
	secretKey := "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz"
	Craftgate := CraftgateClient(apiKey, secretKey)
	res, _ := Craftgate.InstallmentApi.Search(InstallmentParams{binNumber: "487074", price: 100.00})
	fmt.Println(res)
}
