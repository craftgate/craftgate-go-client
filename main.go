package main

import "fmt"

func main() {
	apiKey := "sandbox-YEhueLgomBjqsnvBlWVVuFsVhlvJlMHE"
	secretKey := "sandbox-tBdcdKVGmGupzfaWcULcwDLMoglZZvTz"

	var client = CraftgateClient(apiKey, secretKey)

	res, _ := client.Search(InstallmentParams{binNumber: "487074", price: 100})
	fmt.Println(res)
}
