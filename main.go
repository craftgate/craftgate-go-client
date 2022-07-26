package main

import "fmt"

func main() {
	apiKey := "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE"
	secretKey := "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz"

	var client = CraftgateClient(apiKey, secretKey)

	res, err := client.Search(InstallmentParams{binNumber: "487074", price: 100})
	if err == nil {
		fmt.Println(res)
	}

}
