/*
 * Craftgate API Client
 *
 * Contact: destek@craftgate.io
 */

package main

import (
	"fmt"
	craftgate "github.com/craftgate/craftgate-go-client/v1.0.0/adapter"
)

func main() {
	client, _ := craftgate.New("api-key", "secret-key", "https://api.craftgate.io")
	fmt.Println(client.Info())
}
