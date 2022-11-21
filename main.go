/*
 * Craftgate API Client
 *
 * Contact: destek@craftgate.io
 */

package main

import (
	craftgate "craftgate-go-client/adapter"
	"fmt"
)

func main() {
	client, _ := craftgate.New("api-key", "secret-key", "https://api.craftgate.io")
	fmt.Println(client.Info())
}
