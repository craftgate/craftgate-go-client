/*
 * Craftgate API Client
 *
 * Contact: destek@craftgate.io
 */

package main

import (
	"fmt"
	craftgate "github.com/craftgate/craftgate-go-client/adapter"
)

func main() {
	client, _ := craftgate.New("api-key", "secret-key", "https://api.craftgate.io", make(map[string]string))
	fmt.Println(client.Info())
}
